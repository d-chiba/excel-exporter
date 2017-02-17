# excel-exporter

* エクセルファイルを CSV に変換する

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

