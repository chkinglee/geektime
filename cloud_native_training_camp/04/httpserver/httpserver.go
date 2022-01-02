package main

import (
	"cloud_native_training_camp/04/httpserver/src/handler"
	"log"
	"net/http"
)

func main() {

	handler.RegistryHttpRoutine()
	err := http.ListenAndServe(":8077", nil)
	if err != nil {
		log.Fatal(err)
	}
}
