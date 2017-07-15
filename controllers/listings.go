/* Copyright 2016 Andy Majot. All rights reserved.
 */

// Package controllers defines controllers for each router path
package controllers

import (
	//"fmt"
	"html/template"
	//"net/url"
	//"strconv"

	"github.com/astaxie/beego"

	//"hamlistings/models"
)

// HTML value indicating that a checkbox or radio button is selected
const IsChecked = "checked"

// type ListingsController declares a receiver for methods that define
// action common to all "Listing pages", i.e. pages that alcept user Listings
type ListingsController struct {
	beego.Controller
}

// Local method activeContent defines an active page layout. for example
// "Listing pages" are described by listing-layout.tpl
// 	{{.Header}}
// 	{{.LayoutContent}}
// 	{{.Footer}}
// Note {{.LayoutContent}} is replaced by template content as specified by view
func (lc *ListingsController) activeContent(view string) {
	lc.Layout = "listing-layout.tpl"
	lc.LayoutSections = make(map[string]string)
	lc.LayoutSections["Header"] = "header.tpl"
	lc.LayoutSections["Footer"] = "footer.tpl"
	lc.TplName = view + ".tpl"

	lc.Data["Website"] = SiteTitle
	lc.Data["xsrftoken"] = template.HTML(lc.XSRFFormHTML())
}

// Listing controls the presentation and processing of all "listing pages"
func (lc *ListingsController) Listing() {
	// A session will exist if the user has Signed in or Registered
	// If found, session values are used to initialize controller Data,
	// otherwise user may browse as a guest without Listing privileges
	sess := lc.GetSession("hamlistings")
	lc.activeContent("listing")
	if sess != nil {
		lc.Data["InSession"] = 1            // indicate that user has registered
		sm := sess.(map[string]interface{}) // and init controller Data values
		lc.Data["Username"] = sm["username"]

	} else {
		lc.Data["InSession"] = "" // Don't allow Listings (e.g. see adventurer.tpl)
	}


	// Refresh flash content displayed by active content tpl after redirect
	flash := beego.ReadFromRequest(&lc.Controller)
	if fn, ok := flash.Data["notice"]; ok {
		lc.Data["notice"] = fn
	}

}
