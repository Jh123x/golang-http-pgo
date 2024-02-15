package main

import (
	"log"
	"net/http"
	// _ "net/http/pprof" // Used for PGO
)

var safeMap = NewSafeMap[string, User]()

func handleRegister(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Handling register")
	if r.Method != "POST" {
		w.Write([]byte("Invalid request method"))
		return
	}
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	user := User{Username: username, Password: password, Email: email}
	safeMap.Store(username, user)
	log.Default().Println("register successful with username: ", username, " and email: ", email, " and password: ", password)
	w.Write([]byte("User registered successfully"))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Invalid request method"))
		return
	}
	log.Default().Println("Handling login")
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	user := safeMap.Load(username)
	if user.Password == password {
		log.Default().Println("Login successful with username: ", username, " and password: ", password)
		w.Write([]byte("Login successful"))
		return
	}
	log.Default().Println("Login failed")
	w.Write([]byte("Invalid credentials"))
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Invalid request method"))
		return
	}
	log.Default().Println("Handling update")
	r.ParseForm()
	username := r.FormValue("username")
	email := r.FormValue("email")
	user := safeMap.Load(username)
	user.Email = email
	safeMap.Store(username, user)
	log.Default().Println("Email updated successfully for username: ", username, " and email: ", email)
	w.Write([]byte("Email updated successfully"))
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("Invalid request method"))
		return
	}
	log.Default().Println("Handling delete")
	r.ParseForm()
	username := r.FormValue("username")
	safeMap.Delete(username)
	log.Default().Println("User deleted successfully with username: ", username)
	w.Write([]byte("User deleted successfully"))
}

func handleDefault(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte("Invalid request method"))
		return
	}
	log.Default().Println("Handling default")
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", handleDefault)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/update", handleUpdate)
	http.HandleFunc("/delete", handleDelete)

	log.Default().Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
