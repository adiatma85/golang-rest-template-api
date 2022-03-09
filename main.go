package main

import (
	"github.com/adiatma85/go-tutorial-gorm/src"
	"github.com/adiatma85/go-tutorial-gorm/src/config"
)

func main() {
	config.Initialize()
	src.Run()
	// databaseHost := viper.GetString("Database.Host")
	// databaseUsername := viper.GetString("Database.Username")
	// databasePassword := viper.GetString("Database.Password")
	// databaseDb := viper.GetString("Database.Database")
	// dsn := mysqlSqlString(databaseHost, databaseUsername, databasePassword, databaseDb)
	// fmt.Println(viper.GetString("Database.Host"))
	// fmt.Println(dsn)
}

// func mysqlSqlString(host, username, passwrod, db string) string {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, passwrod, host, db)
// 	return dsn
// }
