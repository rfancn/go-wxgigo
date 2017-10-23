package main

import (
	"github.com/julienschmidt/httprouter"
)

var SERVER_MODULES =  map[string]Moduler{
	"install": &InstallModule{},
}

type ModuleManager struct {
	env *Env
	httpRouter *httprouter.Router
	httpsRouter *httprouter.Router
	//stores initilialized modules
	modules	  map[string]Moduler
}

func NewModuleManager(env *Env) *ModuleManager {
	manager := &ModuleManager{}
	manager.env = env
	manager.modules = make(map[string]Moduler)
	return manager
}

func (self *ModuleManager) loadModules() {
	for moduleName, module := range SERVER_MODULES {
		module.Init(self.env)
		self.modules[moduleName] = module
	}
}

