package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (self *SysModule) ViewTemplate(templateFile string) httprouter.Handle{
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		self.ResponseRender(w, templateFile, nil)
	}
}