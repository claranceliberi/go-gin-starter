package main

import (
	"github.com/claranceliberi/ampersand/pkg/books"
	"github.com/claranceliberi/ampersand/pkg/common/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbHost := viper.Get("DB_HOST").(string)
	dbPort := viper.Get("DB_PORT").(string)
	dbUser := viper.Get("DB_USER").(string)
	dbPass := viper.Get("DB_PASSWORD").(string)
	dbName := viper.Get("DB_NAME").(string)
	dbUrl := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=disable"

	r := gin.Default()
	h := db.Init(dbUrl)

	books.RegisterRoutes(r, h)
	// register more routes here

	r.Run(port)
}
