package webserver

import (
	"net/http"
)

func ServidorWeb() {
	http.HandleFunc("/", home)        //el endpoint y su funcion
	http.ListenAndServe(":8080", nil) //puerto y un segundo parametro
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html") //servir archivo index
}
