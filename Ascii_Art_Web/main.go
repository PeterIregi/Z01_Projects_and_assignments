package main

import (
	"net/http"
	"log"
	"ascii-art-web/handlers"
)

func main()  {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/ascii-art", handlers.AsciiArt)
	http.HandleFunc("/clear", handlers.ClearOutput)
	
	log.Fatal(http.ListenAndServe(":8080", nil))


	
}