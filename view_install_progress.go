package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (self *InstallModule) updateProgress(percentage uint, message string) {
	self.mutex.Lock()
	self.Progress.Percentage = percentage
	self.Progress.Message = message
	self.mutex.Unlock()
}

func (self *InstallModule) ViewInstallProgress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	self.ResponseJson(w, self.Progress)
}