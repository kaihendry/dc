package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Hduration struct {
	From string
	Desc string
}

func main() {

	http.HandleFunc("/favicon.ico", http.NotFound)
	http.HandleFunc("/", index)
	http.HandleFunc("/compare", pform)
	fmt.Println("http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func pform(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	h := Hduration{}
	h.From = r.Form.Get("from")
	fmt.Println(h.From)
	then, err := time.Parse("2006-01-02", h.From)

	if err != nil {
		panic(err)
	}

	duration := time.Since(then)
	h.Desc = duration.String()

	tmpl, err := template.New("test").Parse("Hours since {{.From}} are {{.Desc}}")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, h)
	if err != nil {
		panic(err)
	}

	log.Printf("%s %s %s %s\n", r.RemoteAddr, r.Method, r.URL, r.UserAgent())

}

func index(w http.ResponseWriter, r *http.Request) {

	t, err := template.New("foo").Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
</head>
<body>
<form action="/compare" method="post">
<input type="date" name="from"><br>
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
