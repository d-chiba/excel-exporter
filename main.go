package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

var (
	sheetNames = flag.String("s", "", "カンマ区切りのシート名")
)

func usage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("\texcel-exporter [Flags] FILE.xlsx\n")
	fmt.Printf("Output:\n")
	fmt.Printf("\tSheetName.txt[, SheetName2.txt]\n")
	fmt.Printf("Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	excelFileName := flag.Arg(0)
	if excelFileName == "" {
		usage()
		os.Exit(2)
	}
	run(excelFileName, strings.Split(*sheetNames, ","))
}
func run(excelFileName string, sheets []string) {
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}
	for _, sheet := range xlFile.Sheets {
		if contains(sheets, sheet.Name) {
			b := new(bytes.Buffer)
			for _, row := range sheet.Rows {
				list := []string{}
				for _, cell := range row.Cells {
					text, _ := cell.String()
					list = append(list, text)
				}
				_, err := b.WriteString(strings.Join(list, ",") + "\r\n")
				if err != nil {
					panic(err)
				}
			}
			err := ioutil.WriteFile(sheet.Name+".txt", b.Bytes(), 0644)
			if err != nil {
				panic(err)
			}
		}
	}
}

func contains(list []string, n string) bool {
	if *sheetNames == "" {
		return true
	}
	for _, s := range list {
		if s == n {
			return true
		}
	}
	return false
}
