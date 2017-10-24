package main

import(
	"sync"
)

type ProgressInfo struct {
	Percentage uint
	Message  string
}

type InstallModule struct {
	SysModule
	mutex sync.Mutex
	Progress *ProgressInfo
	//channel indicate already has one session is running deploy process
	ChannelInstallBusy chan bool
	Response *InstallResponse
}

func (self *InstallModule) Init(env *Env) {
	self.super("install", env)

	self.Progress = &ProgressInfo{}
	self.Response = &InstallResponse{}

	//initilize to allow only one session can access ViewInstallSave at one time
	self.ChannelInstallBusy = make(chan bool, 1)
	self.ChannelInstallBusy <- false

	//all install module puts in secure https
	self.env.secureRouter.GET("/install", self.ViewInstallIndex)
	self.env.secureRouter.POST("/install/deploy", self.ViewInstallDeploy)
	self.env.secureRouter.GET("/install/progress", self.ViewInstallProgress)
}

