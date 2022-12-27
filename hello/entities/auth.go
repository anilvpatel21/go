package entities

import (
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

const (
	USERNAME = "AdminDashboardUser"
	PASSWORD = "AdminDashboardPass"
)

func BasicAuth(pass handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := r.Cookie(USERNAME)
		if err != nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		password, err := r.Cookie(PASSWORD)
		if err != nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		if username == nil || password == nil {
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		// authorized := validate(username.Value, password.Value)
		// if authorized == false {
		// 	http.Redirect(w, r, "/auth", http.StatusSeeOther)
		// 	return
		// }
		pass(w, r)
	}
}

// func validate(username, password string) bool {
// 	if val, ok := Config.AuthKeyValueMap[username]; ok {
// 		return val == password
// 	}
// 	return false
// }
