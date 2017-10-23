package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/julienschmidt/httprouter"
	//"fmt"
	"fmt"
	//"time"
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


func (self *InstallModule) ViewInstallSave(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("in install save")

	config := &InstallConfig{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(config); err != nil {
		log.Println("Failed to decode install post data")
	}

	go routineInstallSave(config)

	<-channelInstallQuit

	fmt.Println("go to here")

	self.ResponseJson(w,"test")
}



