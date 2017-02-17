# excel-exporter

* エクセルファイルを CSV に変換する

# How to use

```
$ git clone https://github.com/d-chiba/excel-exporter.git
$ cp excel-exporter/bin/excel-exporter-mac /PATH/TO/EXECUTABLE/excel-exporter
$ excel-exporter EXCEL_FILE > CSV_FILE
```

# Install

```
go get -u github.com/kardianos/govendor
go get -u github.com/d-chiba/excel-exporter
govendor fetch +out
go install github.com/d-chiba/excel-exporter
```

# Build

```
GOOS=darwin GOARCH=amd64 go build -o bin/excel-exporter-mac
GOOS=windows GOARCH=amd64 go build -o bin/excel-exporter-win
```

