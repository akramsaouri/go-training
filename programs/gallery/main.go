package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

const (
	LOGIN        = "root"
	PASSWORD     = "root"
	SESSION_NAME = "admin"
	SESSION_KEY  = "loggedIn"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/", home)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func admin(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, SESSION_NAME)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, ok := session.Values[SESSION_KEY]; ok == false {
		http.Redirect(w, r, "/login", 301)
	}
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("admin.html"))
		err := t.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("imagefield")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	files, err := ioutil.ReadDir("images")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var imageUrls []string
	for _, f := range files {
		imageUrls = append(imageUrls, "images/"+f.Name())
	}
	err = t.Execute(w, imageUrls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, SESSION_NAME)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if session.Values[SESSION_KEY] == true {
		http.Redirect(w, r, "/admin", 301)
	}
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("login.html"))
		if flashes := session.Flashes(); len(flashes) > 0 {
			fmt.Println(flashes)
		}
		err := t.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		r.ParseForm()
		if r.Form["login"][0] == LOGIN && r.Form["password"][0] == PASSWORD {
			session.Values[SESSION_KEY] = true
			session.Save(r, w)
			http.Redirect(w, r, "/admin", 301)
		} else {
			session.AddFlash("Incorrect credentials!")
			http.Redirect(w, r, "/login", 301)
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	// session, err := store.Get(r, SESSION_NAME)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// delete(session.Values, SESSION_KEY)
	// session.Save(r, w)
	// http.Redirect(w, r, "/login", 301)
}
