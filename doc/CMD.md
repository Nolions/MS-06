# GO Lang 常用指令

## 查看環境變數

    go env

## 編譯&執行

    go run

## 編輯成執行檔

    go build <file.go>

### 跨平台編譯

***Build to Windows from MAC***

    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build <file.go>

***Build to Linux from MAC***

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build <file.go>

***Build to MAC from Linux***

    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build <file.go>

***Build to Windows from Linux***

    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build <file.go>

## 執行測試

    go test
