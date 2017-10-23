package main

import (
	"log"
	"net/http"
	//"github.com/rs/cors"
)

func main() {
	env := NewEnv()
	if env == nil {
		log.Fatal("Error prepare running env")
	}

	moduleManager := NewModuleManager(env)
	moduleManager.loadModules()

	//launch secure https server
	go 	http.ListenAndServeTLS(":443", "cert.pem", "key.pem", env.secureRouter)

	//launch http server
	log.Fatal(http.ListenAndServe(":8890", env.httpRouter))

}
