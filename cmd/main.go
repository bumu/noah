package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	ipinfov1 "apigw/gen/go/ipinfo/v1"

	"gopkg.in/resty.v1"
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
					rec.Domain = strings.TrimPrefix(field, "https://")
					rec.Domain = strings.TrimPrefix(rec.Domain, "http://")
					rec.Domain = strings.TrimPrefix(rec.Domain, "/")
				case 2:
					rec.IpStart = field
				case 3:
					rec.IpEnd = field
				case 4:
					n, err := strconv.ParseInt(field, 10, 64)
					if err == nil {
						rec.IpStartInt = uint32(n)
					}
				case 5:
					n, err := strconv.ParseInt(field, 10, 64)
					if err == nil {
						rec.IpEndInt = uint32(n)
					}
				}
			}

			// rec.IpStartInt = ipkit.Ip2Int(rec.IpStart)
			// rec.IpEndInt = ipkit.Ip2Int(rec.IpEnd)

			CallHttpCreate(rec, "/v1/ipdb/datacenter/create")

			// dataList = append(dataList, rec)
			// break
		}
	}
	return dataList
}

func createIpdbClientIpList(data [][]string) []ipinfov1.IpdbClientIpRequest {
	var dataList []ipinfov1.IpdbClientIpRequest
	for i, line := range data {
		if i > 0 { // omit header line
			fmt.Println(line)
			var rec ipinfov1.IpdbClientIpRequest
			for j, field := range line {
				switch j {
				case 0:
					rec.Ip = field
				case 1:
					rec.IspDomain = field
				case 2:
					rec.IspName = field
				case 3:
					rec.IpCity = field
					/*
						case 4:
							n, err := strconv.ParseInt(field, 10, 64)
							if err == nil {
								rec.Cnt = n
							}
					*/
				}
			}
			// fmt.Println("===", rec.Ip, ipkit.Ip2Int(rec.Ip))

			CallHttpCreate(rec, "/v1/ipdb/clientip/create")
		}
	}
	return dataList
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

	fmt.Println("call http done", resp, err)
}

func readFile(filename string) [][]string {
	// open file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	csvReader.LazyQuotes = true // enable lazy quotes to handle escaped double quotes
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	filename := "datacenter.csv"
	filename = "clientip.csv"

	data := readFile(filename)

	// convert records to array of structs
	// createIpdbDatacenterList(data)

	createIpdbClientIpList(data)
}
