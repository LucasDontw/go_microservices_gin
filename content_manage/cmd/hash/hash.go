package main

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"hash/fnv"
	"math/big"
	"github.com/google/uuid"
)

const contentNumTables = 4

func main() {
	// cms_content.t_content_details_0
	// cms_content.t_content_details_1
	// cms_content.t_content_details_2
	// cms_content.t_content_details_3

	u := uuid.New().String()
	getContentDetailsTable(u)
}

func getContentDetailsTable(contentID string) string {
	tableIndex := getContentTableIndex(contentID)
	table := fmt.Sprintf("cms_content.t_content_details_%d", tableIndex)
	log.Infof("content_id = %s, table = %s", contentID, table)

	return table
}

func getContentTableIndex(uuid string) int {
	hash := fnv.New64()
	_, _ = hash.Write([]byte(uuid)) // 將 uuid 字串轉換為位元組陣列，並寫入 hash。這一步將 uuid 的值傳入哈希函數進行處理。
	hashValue := hash.Sum64() // 取得哈希值 hashValue，該值是一個 64 位的整數。這個哈希值是 uuid 經過 FNV-1a 哈希函數處理後的結果，用於分配到不同的分表。
	fmt.Println("hashValue = ", hashValue)

	bigNum := big.NewInt(int64(hashValue))
	mod := big.NewInt(contentNumTables) // 代表分表的總數，即多少個分表。
	tableIndex := bigNum.Mod(bigNum, mod).Int64() // tableIndex 的範圍在 [0, contentNumTables-1] 之間

	return int(tableIndex)
}