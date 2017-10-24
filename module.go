package main

import (
	"fmt"
	"github.com/flosch/pongo2"
	"net/http"
	"encoding/json"
)

type Moduler interface {
	Init(env *Env)
}

type SysModule struct {
	name string
	env *Env
}

func (self *SysModule) Init(env *Env) {
	self.env = env
}

func (self *SysModule) super(moduleName string, env *Env) {
	self.name = moduleName
	env.moduleName = moduleName
	self.Init(env)
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

func (self *SysModule) ResponseRender(w http.ResponseWriter, templateFile string, context map[string]interface{}) {
	output := self.Render(templateFile, context)
	fmt.Fprint(w, output)
}

func (self *SysModule) ResponseText(w http.ResponseWriter, response string) {
	fmt.Fprint(w, response)
}

func (self *SysModule) ResponseJson(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(js))
}