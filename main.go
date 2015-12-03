package main

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type Hduration struct {
	From     string
	Duration string
}

func main() {

	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/", index)
	http.HandleFunc("/compare", pform)

	// Get a free port
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Panic(err)
	}

	// Find the hostname
	hostname, _ := os.Hostname()
	if a, ok := ln.Addr().(*net.TCPAddr); ok {
		host := fmt.Sprintf("http://%s:%d", hostname, a.Port)
		fmt.Println("Serving from", host)
	}
	if err := http.Serve(ln, nil); err != nil {
		log.Panic(err)
	}
}

func pform(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	h := Hduration{}
	h.From = r.Form.Get("from")
	then, err := time.Parse("2006-01-02", h.From)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	duration := time.Since(then)
	h.Duration = duration.String()

	tmpl, err := template.New("test").Parse("Hours since {{.From}} are {{.Duration}}")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, h)
	if err != nil {
		panic(err)
	}

	log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, h.From, r.UserAgent())

}

func index(w http.ResponseWriter, r *http.Request) {

	t, err := template.New("foo").Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<style>
body { font-family: sans-serif; }
form { display: flex; flex-direction: column; }
input { font-size: 1.8em; padding: 0.3em; box-sizing: border-box; display: block; flex: 1;}
</style>
</head>
<body>
<form action="/compare" method="post">
<input type="date" name="from" required>
<input type="submit">
</form>
</body>
</html>`)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)

	log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.UserAgent())

}
