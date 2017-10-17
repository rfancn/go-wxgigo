package main

import (
	"fmt"
	"github.com/flosch/pongo2"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Moduler interface {
	Init(env *Env, router *httprouter.Router)
}

type SysModule struct {
	name string
	env *Env
	router *httprouter.Router
}

func (self *SysModule) Init(env *Env, router *httprouter.Router) {
	self.env = env
	self.router = router
}

func (self *SysModule) super(moduleName string, env *Env, router *httprouter.Router) {
	self.name = moduleName
	env.moduleName = moduleName
	self.Init(env, router)
}

func (self *SysModule) Render(templateFile string, context map[string]interface{}) string {
	templatePath := self.env.getAssetPath(templateFile)
	t := pongo2.Must(pongo2.FromFile(templatePath))
	output, err := t.Execute(context)
	if err != nil {
		fmt.Println(err)
	}

	return output
}

func (self *SysModule) RenderResponse(w http.ResponseWriter, templateFile string, context map[string]interface{}) {
	output := self.Render(templateFile, context)
	fmt.Fprint(w, output)
}
