package main

import (
	"github.com/julienschmidt/httprouter"
)

var SERVER_MODULES =  map[string]Moduler{
	"install": &InstallModule{},
}

type ModuleManager struct {
	env *Env
	router *httprouter.Router
	//stores initilialized modules
	modules	  map[string]Moduler
}

func NewModuleManager(env *Env, router *httprouter.Router) *ModuleManager {
	manager := &ModuleManager{}
	manager.env = env
	manager.router = router
	manager.modules = make(map[string]Moduler)
	return manager
}

func (self *ModuleManager) loadModules() {
	for moduleName, module := range SERVER_MODULES {
		module.Init(self.env, self.router)
		self.modules[moduleName] = module
	}
}

