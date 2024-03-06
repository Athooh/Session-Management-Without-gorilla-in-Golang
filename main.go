package main

import (
	"net/http"
	"text/template"

	"github.com/icza/session"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("template/*.html"))

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "Athooh" && password == "Tiotea" {
		sess := session.NewSessionOptions(&session.SessOptions{
			CAttrs: map[string]interface{}{"username": username},
		})
		session.Add(sess, w)
		http.Redirect(w, r, "/welcome", http.StatusSeeOther)
	} else {
		data := map[string]interface{}{
			"error": "Invalid Username or Password.",
		}
		temp.ExecuteTemplate(w, "index.html", data)
	}
}
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	username := sess.CAttr("username")
	data := map[string]interface{}{
		"username": username,
	}
	temp.ExecuteTemplate(w, "welcome.html", data)
}
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	sess := session.Get(r)
	if sess != nil {
		session.Remove(sess, w)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.ListenAndServe(":8080", nil)
}
