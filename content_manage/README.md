
需安裝項目
kratos:
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

protoc:
至GIHBUT下載protoc-28.3-win64.zip，放到想要的地方，然後添加環境變數，{{自訂路徑}}\protoc-28.3-win64\bin
https://github.com/protocolbuffers/protobuf/releases/

make:
於官方網站下載執行檔，且添加環境變數C:\Program Files (x86)\GnuWin32\bin
https://gnuwin32.sourceforge.net/downlinks/make.php


初始化
1. make init
2. make all
3. make build
4. kratos run

proto檔產出Go檔
make api
