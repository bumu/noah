package loader

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"

	companyv1 "noah/gen/go/company/v1"

	"github.com/go-resty/resty/v2"
)

func createCompanyProfile(data []*companyv1.CreateCompanyProfileRequest) {
	for _, line := range data {
		CallHttpCreate(line, "/apis/v1/company/create")
	}
}

func CallHttpCreate(data interface{}, api string) {
	fmt.Print(data)
	client := resty.New()

	url := "http://127.0.0.1:8080" + api
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		// SetResult(&AuthSuccess{}). // or SetResult(AuthSuccess{}).
		Post(url)

	if err != nil {
		fmt.Println("call http error", err)
		return
	}

	fmt.Println("call http done", resp, err)
}

func readExcelFile(filename string) ([]*companyv1.CreateCompanyProfileRequest, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return nil, err
	}

	// Read origin data
	// headers := []string{}
	list := make([]*companyv1.CreateCompanyProfileRequest, 0, len(rows))

	for idx, row := range rows {
		if idx == 0 {
			// headers = row
			continue
		}

		item := &companyv1.CreateCompanyProfileRequest{}

		item.Uen = ""
		item.Name = row[0]
		item.Address = row[3]
		item.Status = 0

		if item.Website = row[1]; !strings.Contains(item.Website, "http") {
			item.Website = ""
		}

		if item.Phone = row[4]; item.Phone == "" {
			if strings.Contains(row[1], "+65") {
				item.Phone = row[1]
			}

			if strings.Contains(row[2], "+65") {
				item.Phone = row[2]
			}
		}

		item.Phone = strings.ReplaceAll(item.Phone, "tel:", "")

		list = append(list, item)

	}

	return list, nil
}

func LoadCompany() {
	filename := "company.xlsx"
	// read data from file
	data, _ := readExcelFile(filename)

	createCompanyProfile(data)
}
