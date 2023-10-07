package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	ipinfov1 "apigw/gen/go/ipinfo/v1"
)

func createIpdbDatacenterList(data [][]string) []ipinfov1.IPDataCenterRequest {
	var dataList []ipinfov1.IPDataCenterRequest
	for i, line := range data {
		if i > 0 { // omit header line
			var rec ipinfov1.IPDataCenterRequest
			for j, field := range line {
				switch j {
				case 0:
					rec.Name = field
				case 1:
					rec.Domain = field
				case 2:
					rec.IpStart = field
				case 3:
					rec.IpEnd = field
				case 4:
					n, err := strconv.ParseInt(field, 10, 64)
					if err == nil {
						rec.CreatedTs = n
					}
				case 5:
					n, err := strconv.ParseInt(field, 10, 64)
					if err == nil {
						rec.UpdatedTs = n
					}
				}
			}

			CallHttp(rec)

			// dataList = append(dataList, rec)
			// break
		}
	}
	return dataList
}

func CallHttp(data ipinfov1.IPDataCenterRequest) {
	fmt.Print(data)
	fmt.Println("call http done")
}

func main() {
	// open file
	f, err := os.Open("datacenter.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	createIpdbDatacenterList(data)
}
