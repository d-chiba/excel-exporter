package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

func Usage() {
	fmt.Println(`
Usage: excel-exporter FILE.xlsx > OUTPUT.csv
`)
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	excelFileName := flag.Arg(0)
	if excelFileName == "" {
		Usage()
		os.Exit(2)
	}
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			list := []string{}
			for _, cell := range row.Cells {
				text, _ := cell.String()
				list = append(list, text)
			}
			fmt.Printf("%s\r\n", strings.Join(list, ","))
		}
	}
}
