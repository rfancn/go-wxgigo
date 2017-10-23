package main

import(
	//"net/http"
	//"log"
	//"fmt"
	//"log"
)


type InstallModule struct {
	SysModule
}

func (self *InstallModule) Init(env *Env) {
	self.super("install", env)

	//all install module puts in secure https
	self.env.secureRouter.GET("/install", self.ViewInstallIndex)
	self.env.secureRouter.POST("/install/save", self.ViewInstallSave)
	self.env.secureRouter.GET("/install/progress", self.ViewInstallProgress)
}