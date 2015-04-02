package main

import (
	"flag"
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type Person struct {
	Username string
	Password []byte
	Todos    string
}

// func main() {

// 	hashPass, err := bcrypt.GenerateFromPassword([]byte("test1"), bcrypt.DefaultCost)

// 	err = c.Insert(&Person{"Mike", hashPass, "yeah"})

// 	if err != nil {
// 		log.Printf("I am the insert error", err)
// 	}

// 	result := Person{}

// 	err = c.Find(bson.M{"username": "Mike"}).One(&result)

// 	if err != nil {
// 		log.Printf("I am the retrieve error", err)
// 	}

// 	err = bcrypt.CompareHashAndPassword(result.Password, []byte("test1"))

// 	if err != nil {
// 		log.Printf("I am the compare error", err)
// 		return
// 	}

// 	fmt.Println("Password:", result.Password)
// 	fmt.Println("Password:", hashPass)
// }

func mainHandler(w http.ResponseWriter, r *http.Request) {

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(404)
	}
	w.Write([]byte(")]}',['sup',  1]"))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// session, err := mgo.Dial("mongodb://localhost/")

	// if err != nil {
	// 	log.Printf("I am the error", err)
	// 	log.Printf("I am the session", session)
	// }

	// defer session.Close()

	// //setting c to the current DB session.
	// c := session.DB("test").C("people")

	// hashPass, err := bcrypt.GenerateFromPassword([]byte("test1"), bcrypt.DefaultCost)

	// err = c.Insert(&Person{"Mike", hashPass, "yeah"})

	// if err != nil {
	// 	log.Printf("I am the insert error", err)
	// }

	// result := Person{}

	// err = c.Find(bson.M{"username": "Mike"}).One(&result)

	// if err != nil {
	// 	log.Printf("I am the retrieve error", err)
	// }

	// err = bcrypt.CompareHashAndPassword(result.Password, []byte("test1"))

	// if err != nil {
	// 	log.Printf("I am the compare error", err)
	// 	return
	// }

	log.Printf("Password:", r.Body)
}

func main() {

	port := flag.Int("port", 3000, "server port")
	dir := flag.String("directory", "client/", "static assets")
	flag.Parse()

	fs := http.Dir(*dir)
	handler := http.FileServer(fs)

	http.Handle("/", handler)
	http.HandleFunc("/main", mainHandler)
	http.Post.HandleFunc("/auth/login", loginHandler)
	http.HandleFunc("/auth/signup", signupHandler)

	log.Printf("Running on port %d\n", *port)
	addr := fmt.Sprintf("127.0.0.1:%d", *port)

	http.ListenAndServe(addr, nil)
}
