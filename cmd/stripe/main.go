package main

import (
	"fmt"
	"os"

	"github.com/bokimilinkovic/simple-accounts-app/pkg/database/gorm"
	"github.com/bokimilinkovic/simple-accounts-app/pkg/handler"
	custommiddleware "github.com/bokimilinkovic/simple-accounts-app/pkg/handler/middleware"

	"github.com/bokimilinkovic/simple-accounts-app/pkg/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	secretKey = []byte("boki-demo")
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

	db.AutoMigrate(&model.User{}, &model.StripePayment{}, &model.Movie{})
	// populate DB
	// check if empty table
	var count int64
	var movies []model.Movie
	if err := db.Find(&movies).Count(&count).Error; err != nil {
		panic(err)
	}
	if count == 0 {
		for _, movie := range movies {
			if err := db.Save(&movie).Error; err != nil {
				logrus.WithField("movie", movie.Title).Error(err.Error())
				continue
			}
		}
	}

	// create server, register endpoints...
	e := echo.New()
	// Middleware
	// e.Use(echomiddleware.Logger())
	// e.Use(echomiddleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	authHandler := handler.NewAuthHandler(db, secretKey)
	userLoaderMiddleware := custommiddleware.NewUserLoader(secretKey)
	movieHandler := handler.NewMovieHandler(db)

	e.POST("/signin", authHandler.Signin)
	e.GET("/movies", movieHandler.GetMovies, userLoaderMiddleware.Do)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", viper.Sub("server").GetInt("port"))))

}
