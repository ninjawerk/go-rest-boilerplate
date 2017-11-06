/**
 * Created by VoidArtanis on 10/22/2017
 */

package shared

import "fmt"
import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"
//import _ "github.com/jinzhu/gorm/dialects/mysql"
// import _ "github.com/jinzhu/gorm/dialects/sqlite"
// import _ "github.com/jinzhu/gorm/dialects/mssql"

var db *gorm.DB
var err error

/*
dbType can be 'MySql', 'Postrges', ''
 */
func Init() {
	////MySQL
	//db, err = gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

	//PostgreSQL
	db, err = gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")

	////SQLite3
	//db, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	//
	////SQL Server
	//db, err = gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")

	if err != nil {
		fmt.Println(err)
	}

	//db.AutoMigrate(&models.Person{})
}


func GetDb() *gorm.DB {
	return db
}

func CloseDb() {
	db.Close()
}
