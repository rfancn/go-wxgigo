package main

import (
	"net/http"
	"bytes"
	"github.com/julienschmidt/httprouter"
	y2h "github.com/rfancn/go-y2h"
	"log"
)


var INSTALL_WIZARD_STEPS = []string {
	"general",
	"wxmp",
	"celery",
	"server",
}
var INSTALL_WIZARD_YAMLS = map[string]string {
	"general": "step_general.yaml",
	"wxmp":     "step_wxmp.yaml",
	"celery":   "step_celery.yaml",
	"server":   "step_server.yaml",
}

//yaml definition for edit server
const INSTALL_SERVER_EDIT_MODAL_YAML = "modal_server.yaml"

var gY2H = y2h.New()

func (self *InstallModule) ViewInstallIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("in install index")

	//build context
	context := make(map[string]interface{})
	context["wizard"] = self.getInstallWizard()
	context["modal"] = self.getServerEditModal()

	self.ResponseRender(w,"index.html", context)
}


func (self *InstallModule) getServerEditModal() map[string]string {
	yamlFilePath := self.env.getAssetPath(INSTALL_SERVER_EDIT_MODAL_YAML)
	if ok := gY2H.Read(yamlFilePath); !ok{
		return nil
	}

	serverEditModal := make(map[string]string)
	serverEditModal["html"] = gY2H.GetHtml()
	inlineJs, ok := gY2H.GetJavascript()["inline"]
	if ok {
		serverEditModal["inlineJs"] = inlineJs
	}

	return serverEditModal
}

//build wizard from YAML file
func (self *InstallModule) getInstallWizard() map[string]string {
	var wizard = make(map[string]string)
	var inlineJsBuffer bytes.Buffer
	var externalJsBuffer bytes.Buffer

	//only iterate steps by name then we can make sure inline js
	//are assembled by order
	for _, stepName := range INSTALL_WIZARD_STEPS {
		yamlFilename, _ := INSTALL_WIZARD_YAMLS[stepName]
		yamlFilePath := self.env.getAssetPath(yamlFilename)
		if ok := gY2H.Read(yamlFilePath); !ok{
			continue
		}

		//get HTML output
		wizard[stepName] = gY2H.GetHtml()

		//get Javascript output
		for jsType, jsContent := range gY2H.GetJavascript() {
			switch jsType {
			case "inline":
				inlineJsBuffer.WriteString(jsContent)
			case "external":
				externalJsBuffer.WriteString(jsContent)
			}
		}
	}

	wizard["inlineJs"] = inlineJsBuffer.String()
	wizard["externalJs"] = externalJsBuffer.String()

	return wizard
}
