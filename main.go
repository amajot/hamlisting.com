// Package main defines functions to initialize database access via Beego's ORM,
// enables persistent sessions and starts the app's built-in server
package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// Preload values used by Beego on startup
	_ "hamlistings/models"
	_ "hamlistings/routers"
	_ "github.com/go-sql-driver/mysql"
)

// Init resgisters the hamlistings database with Beego's ORM
// IMPORTANT: Edit hamlistings/conf/app.conf changing 'mysqluser', 'mysqlpass', 'mysqlhost',
// and 'mysqldb' to the values you use to access your MySQL database

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	mysqlReg := beego.AppConfig.String("mysqluser") + ":" +
		beego.AppConfig.String("mysqlpass") + "@tcp(" + beego.AppConfig.String("mysqlhost") + ":3306)/" +
		beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", mysqlReg)
}

// Main is the program's entry point after all imports and
// related init functions have been loaded/executed
// Explicit Beego options can set here before calling Beego.Run()
func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
