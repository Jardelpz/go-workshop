package src

import (
	"fmt"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	dbName = "go_workshop"
	dbUser = "admin"
	dbPort = "5432"
	// dbHost = "127.0.0.1"
	// dbPass = "admin"
)

var (
	dbHost string = os.Getenv("DB_HOST")
	dbPass string = os.Getenv("DB_PASSWD")
)

func dbConnect() *gorm.DB {
	urlConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPass)

	db, err := gorm.Open("postgres", urlConn)

	db.DB().SetMaxIdleConns(10)  //maximo de conexoes
	db.DB().SetMaxOpenConns(100) // maximo de conexoes abertas

	if err != nil {
		fmt.Println("Error openning conection with DataBase")
		panic(err)
	}

	return db

}

func selectAll(table interface{}, c *gin.Context) interface{} {
	db := dbConnect()
	if err := db.Find(reflect.ValueOf(table).Interface()).Error; err != nil {
		checkErr(err, 500, c)
	}
	return table
}

func selecDebtId(ID string, c *gin.Context) (debt, *gorm.DB) {
	db := dbConnect()
	var debtFind debt

	if err := db.Where("id = ? ", ID).First(&debtFind).Error; err != nil {
		checkErr(err, 500, c)
	}

	return debtFind, db
}

func selectUserID(ID int, c *gin.Context) (user, *gorm.DB) {
	db := dbConnect()
	var userFind user

	if err := db.Where("id = ?", ID).First(&userFind).Error; err != nil {
		checkErr(err, 500, c)
	}

	return userFind, db
}

func checkErr(err error, statusCode int, c *gin.Context) {
	if err != nil {
		c.AbortWithError(statusCode, err)
		return
	}
}
