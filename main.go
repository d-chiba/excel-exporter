package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/naoina/toml"
	"github.com/tealeg/xlsx"
)

type config struct {
	InputDir  string
	OutputDir string
}

var (
	configName = flag.String("c", "", "設定ファイル名(default: ~/.excel-exporter.toml)")
	sheetNames = flag.String("s", "", "カンマ区切りのシート名")
)

func usage() {
	fmt.Printf("Usage:\n")
	fmt.Printf("\texcel-exporter [Flags] FILE.xlsx\n")
	fmt.Printf("Output:\n")
	fmt.Printf("\tSheetName.txt[, SheetName2.txt]\n")
	fmt.Printf("Config file Example:\n")
	fmt.Printf("\tInputDir  = \"/Path/To/Excel/Dir/\"\n")
	fmt.Printf("\tOutputDir = \"/Path/To/Csv/Dir/\"\n")
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
	c := parseConfig(*configName)
	run(c, excelFileName, strings.Split(*sheetNames, ","))
}

func parseConfig(name string) config {
	if name == "" {
		home, e := homedir.Dir()
		if e != nil {
			panic(e)
		}
		name = home + "/.excel-exporter.toml"
	}

	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var c config
	if err := toml.Unmarshal(buf, &c); err != nil {
		panic(err)
	}
	return c
}

func run(c config, excelFileName string, sheets []string) {
	xlFile, err := xlsx.OpenFile(c.InputDir + "/" + excelFileName)
	if err != nil {
		panic(err)
	}
	for _, sheet := range xlFile.Sheets {
		if contains(sheets, sheet.Name) {
			b := new(bytes.Buffer)
			_, err := b.WriteString("\ufeff")
			if err != nil {
				panic(err)
			}
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
			err = ioutil.WriteFile(c.OutputDir+"/"+sheet.Name+".txt", b.Bytes(), 0644)
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
