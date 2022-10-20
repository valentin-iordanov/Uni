package routes

import (
	"net/http"
	"valio/routes/bgtoll"
)

func Routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/bgtoll",bgtoll.BgTollHanler)

	// router.HandleFunc("/bgtoll",)

	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	return router
}


