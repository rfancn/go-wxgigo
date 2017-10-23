package main

import (
	"runtime"
	"path/filepath"
	"os"
	"log"
	"github.com/kabukky/httpscerts"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Env struct{
	moduleName string
	rootDir string
	assetsRootDir string
	httpRouter *httprouter.Router
	secureRouter *httprouter.Router
}

func NewEnv() *Env {
	// Check if the cert files are available.
	if err := httpscerts.Check("cert.pem", "key.pem"); err != nil {
		// If they are not available, generate new ones.
		err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:443")
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
			return nil
		}
	}

	rootDir, ok := getRootDir()
	if !ok {
		return nil
	}

	assetsRootDir, ok := getSubDir(rootDir, "assets")
	if !ok {
		return nil
	}

	env := &Env{}
	env.rootDir = rootDir
	env.assetsRootDir = assetsRootDir
	//for normal http
	env.httpRouter = env.getDefaultHttpRouter()
	//for secure https
	env.secureRouter = env.getDefaultHttpRouter()
	return env
}

func (env *Env) getDefaultHttpRouter() *httprouter.Router {
	router := httprouter.New()

	//disable automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = false

	//server static assets file
	router.ServeFiles("/assets/*filepath", http.Dir(env.assetsRootDir))

	return router
}

func getRootDir() (string, bool) {
	//curFilePath is: /path/to/go-wxgigo/env.go
	_, curFilePath, _, ok := runtime.Caller(0)
	if !ok {
		return "", false
	}

	return filepath.Dir(curFilePath), true
}

func getSubDir(rootDir string, dirname string) (string, bool) {
	subdir := filepath.Join(rootDir, dirname)
	stat, err := os.Stat(subdir)
	if err != nil {
		return "", false
	}
	if !stat.IsDir() {
		return "", false
	}

	return subdir, true
}


func (self *Env) getAssetPath(assetName string) string {
	extension := filepath.Ext(assetName)
	switch extension {
	case ".html",".htm":
		return filepath.Join(self.assetsRootDir, self.moduleName, "html", assetName)
	case ".js":
		return filepath.Join(self.assetsRootDir, self.moduleName, "js", assetName)
	case ".css":
		return filepath.Join(self.assetsRootDir, self.moduleName, "css", assetName)
	case ".yaml",".yml":
		return filepath.Join(self.assetsRootDir, self.moduleName, "yaml", assetName)
	default:
		return ""
	}
}

