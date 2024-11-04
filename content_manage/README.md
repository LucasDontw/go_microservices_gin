
#需安裝項目#
##kratos:##
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest


##protoc:##
至GIHBUT下載protoc-28.3-win64.zip，放到想要的地方，然後添加環境變數，{{自訂路徑}}\protoc-28.3-win64\bin
https://github.com/protocolbuffers/protobuf/releases/

##make:##
於官方網站下載執行檔，且添加環境變數C:\Program Files (x86)\GnuWin32\bin
https://gnuwin32.sourceforge.net/downlinks/make.php


##初始化##
1. make init
2. make all
3. make build
4. kratos run


##proto檔產出Go檔##
make api


##微服務化邏輯順序##
1. API 層：接收並解析請求，傳遞參數給 Service 層。

2. Service 層：呼叫 Biz 層的 Usecase 方法處理業務邏輯。

3. Biz 層：調用 Data 層的 Repository 進行數據存取操作。

4. Data 層：執行數據庫操作（如新增、查詢），返回結果給 Biz 層。

5. Biz 層：處理 Data 層返回的結果（例如驗證數據、進行業務邏輯），並返回給 Service 層。

6. Service 層：將 Biz 層的結果傳回給 API 層。

7. API 層：將結果格式化為 HTTP 回應，回傳給客戶端。
完整流程圖請參考./all_lay_flow.drawio

##wire說明##
只要各部分的wire.NewSet()有將特定變數註冊進去，wire就會將所有相對應到的類別，
直接注入到相對應的函數變數中，即可直接使用


##補充##
Biz跟Data層有兩個相近的type struct，兩個雖然可能一模一樣，但不可以合成一個使用，因為使用意義不同，
分開兩個也是避免過度耦合，也能促進層級間的單一職責原則

biz => 通常是用來定義`業務邏輯`需要的數據結構，它關注的是`業務邏輯`本身，而不是資料存取細節
data => 結構體是針對`資料庫`或其他`持久化層`設計的，可能包含一些儲存相關的字段（如 ID、資料庫索引或時間戳等），以及適合資料庫存取的結構