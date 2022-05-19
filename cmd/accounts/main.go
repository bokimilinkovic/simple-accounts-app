package main

import (
	"fmt"
	"os"

	"github.com/adjust/rmq"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"github.com/bokimilinkovic/simple-accounts-app/pkg/database/gorm"
	"github.com/bokimilinkovic/simple-accounts-app/pkg/database/redis"
	"github.com/bokimilinkovic/simple-accounts-app/pkg/handler"
)

func init() {
	viper.SetConfigName("config") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // config file path
	viper.AutomaticEnv()     // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

}

func main() {

	var dbConfig gorm.Config
	if err := viper.Sub("database").Unmarshal(&dbConfig); err != nil {
		panic(err)
	}

	db, err := gorm.CreateConnection(dbConfig)
	if err != nil {
		panic(err)
	}

	var redisConfig redis.RedisConfig
	if err := viper.Sub("redis").Unmarshal(&redisConfig); err != nil {
		panic(err)
	}
	connection := rmq.OpenConnection("producer", "tcp", fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port), 2)
	queue := connection.OpenQueue("transactions")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	accountHandler := handler.NewAccountHandler(db)
	transactionHandler := handler.NewTransactionsHandler(db, queue)

	e.GET("/accounts", accountHandler.GetAccounts)
	e.POST("/accounts", accountHandler.CreateAccount)
	e.GET("/account/:id", accountHandler.GetAccount)
	e.DELETE("/accounts/:id", accountHandler.DeleteAccount)
	e.PUT("/accounts/:id", accountHandler.UpdateAccount)

	e.POST("/transactions", transactionHandler.CreateTransaction)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", viper.Sub("server").GetInt("port"))))
}
