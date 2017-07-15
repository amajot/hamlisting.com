package filters

import (
	"net/http"
	"strings"
	"github.com/astaxie/beego/context"
)

func Auth(ctx *context.Context) {

	if strings.HasPrefix(ctx.Input.URL(), "/login") {
		return
	}

	sess := ctx.Input.Session("hamlistings")
    
    if sess == nil {
        if !ctx.Input.AcceptsHTML() {
			http.Error(ctx.ResponseWriter, "Not Authorized", 401)
		}
		ctx.Redirect(302, "/login")
    }

}
