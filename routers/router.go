/* Copyright 2016 Andy Majot. All rights reserved.
 */

// Package routers associates defined routes with controller functions
// that can handle those routes. Note that the Beego framework can handle
// complex routes, but so far this app uses only simple routes.
package routers

import (
	"github.com/astaxie/beego"
	"hamlistings/filters"
	"hamlistings/controllers"
)

func init() {
	//auth filtering
	beego.InsertFilter("/logout", beego.BeforeRouter, filters.Auth)
	beego.InsertFilter("/profile", beego.BeforeRouter, filters.Auth)
	beego.InsertFilter("/profile/delete", beego.BeforeRouter, filters.Auth)
	
	//routers
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.UserController{}, "get,post:Register")
	beego.Router("/resendValidation/:username", &controllers.UserController{}, "get,post:ResendValidation")
	beego.Router("/login", &controllers.UserController{}, "get,post:Login")
	beego.Router("/profile", &controllers.UserController{}, "get,post:Profile")
	beego.Router("/profile/delete", &controllers.UserController{}, "get:DeleteProfile")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/listings", &controllers.ListingsController{}, "get:Listing")
	beego.Router("/construction", &controllers.ConstructionController{}, "get:Construction")
	beego.Router("/verify/:uuid", &controllers.UserController{}, "get:Verify")
	beego.Router("/forgot", &controllers.UserController{}, "get,post:Forgot")
	beego.Router("/reset/:uuid", &controllers.UserController{}, "get,post:Reset")
}
