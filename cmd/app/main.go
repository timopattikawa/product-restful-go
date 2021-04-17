package main

import (
	"log"

	"github.com/spf13/viper"
	mysqlcon "github.com/timopattikawa/mtp-restful-product/common/mysql"
	"github.com/timopattikawa/mtp-restful-product/product/delivery/api"
	"github.com/timopattikawa/mtp-restful-product/product/repository/mysql"
	"github.com/timopattikawa/mtp-restful-product/product/usecase"
)

func main() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config/")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	mysqlConn := viper.GetString("mysqlDB.connection")

	sqlDB, err := mysqlcon.NewConnetions(mysqlConn)

	if err != nil {
		log.Fatal(err)
	}

	repositoryProduct := mysql.New(sqlDB)
	ucProduct := usecase.New(repositoryProduct)

	deliveryProduct := api.New(ucProduct)

	deliveryProduct.Serve()
}
