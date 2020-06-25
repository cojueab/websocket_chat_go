package routes

import "net/http"

func activeStatic(){
	http.Handle("/", http.FileServer(http.Dir("./public")))
}
