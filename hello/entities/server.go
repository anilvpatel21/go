package entities

import (
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

var (
	InstanceTemplate *template.Template
)

func Start() {
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/", RecoverWrap(http.HandlerFunc(LoginPage))).Methods("GET")
	router.Handle("/login_callback", RecoverWrap(http.HandlerFunc(LoginPage))).Methods("GET")
	router.Handle("/home", RecoverWrap(http.HandlerFunc(HomePage))).Methods("GET")
	router.Handle("/auth", RecoverWrap(http.HandlerFunc(ResetAuth))).Methods("GET")

	log.Println("Listening on 80")
	log.Fatal(http.ListenAndServe(":80", router))
}

func RecoverWrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				http.Redirect(w, req, "/", 307)
			}
		}()
		h.ServeHTTP(w, req)
	})
}

func LoadTemplates() {
	log.Println("Loading gohtml templates")
	var err error

	InstanceTemplate, err = template.ParseFiles(Config.TemplateDir + "/vm_page.gohtml")
	if err != nil {
		log.Fatalln("error parsing vm_page.gohtml: ", err.Error())
	}
}
