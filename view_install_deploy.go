package main

import (
	"net/http"
	"encoding/json"
	"log"
	"time"
	"github.com/julienschmidt/httprouter"
)

type ConfigGeneral struct {
	Debug string `json:"debug"`
}

type ConfigWxmp struct {
	Url string `json:"url"`
	Token string `json:"token"`
	Key string `json:"key"`
	Method string `json:"method"`
}

type ConfigCelery struct {
	Broker string `json:"broker"`
	Backend string `json:"backend"`
}

type ConfigServer struct {
	Host string `json:"host"`
	Port uint `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Roles []string `json:"roles"`
}

type InstallConfig struct {
	General ConfigGeneral `json:"general"`
	Wxmp ConfigWxmp `json:"wxmp"`
	Celery ConfigCelery `json:"celery"`
	Servers []ConfigServer `json:"server"`
}

type InstallResponse struct {
	Result string
	Detail string
}

func (self *InstallModule) ViewInstallDeploy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	select {
		case <-self.ChannelInstallBusy:
			self.startDeploy(w,r,ps)
		case <-time.After(100 * time.Millisecond):
			self.abortDeploy(w,r,ps)
	}
}

func (self *InstallModule) abortDeploy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Deploy aborted")
	self.Response.Result = "error"
	self.Response.Detail = "There is a deploy instance already running!"

	self.ResponseJson(w, self.Response)
}

func (self *InstallModule) startDeploy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Deploy started")

	config := &InstallConfig{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(config); err != nil {
		log.Println("Failed to decode install post data")
	}

	//launch install deploy routine
	channelQuit := self.createDelployRoutine(config)

	//stuck here for waiting for quit signal
	<-channelQuit

	self.Response.Result = "success"
	self.Response.Detail = "/install/success"
	self.ResponseJson(w, self.Response)

	self.ChannelInstallBusy <- false
}



