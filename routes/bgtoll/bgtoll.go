package bgtoll

import (
	"net/http"
	"valio/templates"
)

type BgToll struct {
	Scripts []string
}

func BgTollHanler(w http.ResponseWriter, r *http.Request){
	page := templates.NewTemplate("bgtoll/home.html", templates.WithData(map[string]interface{}{
		"scripts": addJS(),
	}), templates.WithMinified(true)).ToString();
	w.Write([]byte(page))
}


func BgTollReturnInfo(w http.ResponseWriter, r *http.Request){
	r.PostFormValue("value")
}

// func bgtollScripts(scriptsArr []string) string{
// 	scripts := ""

// 	for _,script := range scriptsArr {
// 		scripts += "<script>" + script + "</script> \n"
// 	}

// 	return scripts
// }

func addJS() string {
	return "\n<script>\n" +  templates.NewTemplate("bgtoll/script.js").Minify(true).ToString() + "\n</script>"
}