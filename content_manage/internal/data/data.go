package data

import (
	"content_manage/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
// 如果有多個要注入直接(NewContentRepo1, NewContentRepo2, ...)往後增加即可
var ProviderSet = wire.NewSet(NewData, NewContentRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// 初始化所需連線 .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	// 初始化DB，資料寫在../configs/config.yaml
	mysqlDB, err := gorm.Open(mysql.Open(c.GetDatabase().GetSource()))

	if err != nil {
		panic(err)
	}

	db, err := mysqlDB.DB()

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	return &Data{db: mysqlDB}, cleanup, nil
}
