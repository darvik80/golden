package ui

import (
	"net/http"
//	"github.com/gorilla/mux"
	"path/filepath"
	"html/template"
)

type WelcomeAction struct {
	Action
}

func (self *WelcomeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Prepare cache */
	lp := filepath.Join("views", "layout.tmpl")
	fp := filepath.Join("views", "welcome.tmpl")
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		panic(err)
	}

	webSite := self.Site

	/* Get area manager */
	areaManager := webSite.GetAreaManager()

	/* Render */
	outParams := make(map[string]interface{})
	outParams["Areas"] = areaManager.GetAreas()
	tmpl.ExecuteTemplate(w, "layout", outParams)

}
