package entities

import (
	"log"
	"net/http"
)

func LoginPage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, Config.TemplateDir+"/login.html")
}

func HomePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, Config.TemplateDir+"/index.html")
}

func ResetAuth(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, Config.TemplateDir+"/updateAuth.html")
}

func InstancePage(res http.ResponseWriter, req *http.Request) {
	log.Println("Instance Page: ")
}

func HandleError(res http.ResponseWriter, req *http.Request, err error) {
	log.Println("Request: ", req.URL, " Error(InstancePage): ", err)
	//InstanceResponse(res, req, "", "", req.FormValue("zone"))
}
