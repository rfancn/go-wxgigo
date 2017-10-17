package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)


func main() {
	env := NewEnv()
	if env == nil {
		log.Fatal("Error prepare running env")
	}

	router := httprouter.New()
	//server static assets file
	router.ServeFiles("/assets/*filepath", http.Dir(env.assetsRootDir))

	moduleManager := NewModuleManager(env, router)
	moduleManager.loadModules()

	log.Fatal(http.ListenAndServe(":443", router))
}
