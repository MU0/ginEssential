package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"xjtu.teach/ginEssential/common"
)

func main() {
	InitConfig()

	db := common.InitDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	r := gin.Default()

	r = CollectRoute(r)

	//监听端口
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
