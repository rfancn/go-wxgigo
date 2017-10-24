package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"time"
)

type ProgressInfo struct {
	Percentage uint
	Message  string
}

var currentProgress *ProgressInfo

func (self *InstallModule) ViewInstallProgress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("in install progess")

	for {
		select {
			case currentProgress = <-channelInstallProgress:
				fmt.Println("recevied progress info")
				self.responseProgress(w)
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Timed out!")
				self.responseProgress(w)
				return
		}
	}
}

func (self *InstallModule) responseProgress(w http.ResponseWriter) {
	js, err := json.Marshal(currentProgress)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(js))
	self.ResponseJson(w, string(js))
}