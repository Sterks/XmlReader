package router

import (
	"log"
	"net/http"
	"time"

	"github.com/Sterks/XmlReader/cmd/readerxml/controllers"

	"github.com/gorilla/mux"
)

//StartServer ...
func StartServer() {
	var dir string
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HomeHandler)
	http.Handle("/", r)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
