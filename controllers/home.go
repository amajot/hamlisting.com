package controllers

import (
	"html/template"
	"github.com/astaxie/beego"
)

var (
	// Name used for website title
	SiteTitle string = beego.AppConfig.String("apptitle")
)

// Define category code constants
const (
	Unknown = iota
	Adventurer
	Advocate
	Provider
	Supporter
	Home
	Faq
	Blog
	About
)

// TODO identify and add other constant strings
const (
	DenyAutoLogin = "Another user has already enabled auto-login on this device"
)

// type Home Controller declares a receiver for methods that define
// action on the home page. This includes all user auth actions
type HomeController struct {
	beego.Controller
}

/*  Local method activeContent defines an active page layout.
 *  For example the home page as described by landing-layout.tpl
 * 	{{.Header}}
 * 	{{.LayoutContent}}
 * 	{{.Home}}
 * 	{{.Footer}}
 * Note {{.LayoutContent}} is replaced by template content specified by view
 */
func (hc *HomeController) activeContent(view string) {
	hc.Layout = "landing-layout.tpl"
	hc.LayoutSections = make(map[string]string)
	hc.LayoutSections["Header"] = "header.tpl"
	hc.LayoutSections["Footer"] = "footer.tpl"
	hc.TplName = view + ".tpl"

	hc.Data["Website"] = SiteTitle
	hc.Data["xsrftoken"] = template.HTML(hc.XSRFFormHTML())
	hc.Data["IsHome"] = true
}

/* Get selects the active content handler for "/" (the landing page)
 * it presents and processes the selected template
 */
func (hc *HomeController) Get() {

	sess := hc.GetSession("hamlistings")
	if sess != nil {
		hc.Data["InSession"] = 1            // indicate that user has registered
		sm := sess.(map[string]interface{}) // and init controller Data values
		hc.Data["Username"] = sm["username"]
		
	} else {
		hc.Data["InSession"] = ""   // indicate that user has not registered
	}

	hc.activeContent("welcome") // landing page for logged in user



}

