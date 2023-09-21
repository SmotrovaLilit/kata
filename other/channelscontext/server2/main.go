package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := r.URL.Query().Get("number")
		switch n {
		case "0":
			time.Sleep(1 * time.Second)
		case "1":
			time.Sleep(5 * time.Second)
		case "2":
			time.Sleep(10 * time.Second)
		case "3":
			time.Sleep(15 * time.Second)
		case "4":
			time.Sleep(20 * time.Second)
		}
		_, _ = w.Write([]byte("Number " + n))
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
