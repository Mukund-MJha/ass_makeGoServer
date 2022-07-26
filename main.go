package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type User struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}

func main() {
	fmt.Println("Starting Server...")
	http.HandleFunc("/", ServeHttp)
	http.ListenAndServe("localhost:8080", nil)
}

func ServeHttp(rw http.ResponseWriter, r *http.Request) {
	data := []byte("Hello I am Mukund, ur data is being processed.")
	rw.Write(data)

	fmt.Fprint(rw, "\n Getting users \n")

	time.Sleep(100 * time.Millisecond)
	switch r.Method {
	case "GET":
		u := User{}
		file, _ := os.ReadFile("data.txt")
		data := strings.Split(string(file), "\n")
		u.User = data[0]
		u.Password = data[1]
		fmt.Fprintf(rw, "Incoming Get Request!\n")
		json.NewEncoder(rw).Encode(u)
	case "POST":
		u := User{}
		fmt.Fprintf(rw, "Incoming Post Request!\n")
		json.NewDecoder(r.Body).Decode(&u)
		fmt.Fprintf(rw, "User = %s\n", u.User)
		fmt.Fprintf(rw, "Password = %s\n", u.Password)

		f, _ := os.Create("data.txt")
		f.WriteString(u.User + "\n" + u.Password)
		defer f.Close()
	default:
		fmt.Fprintf(rw, "Only GET and POST methods are supported.")
	}

}
