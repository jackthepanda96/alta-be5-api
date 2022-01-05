package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	ID      int
	Title   string
	Content string
}

var data = []Article{
	{1, "Tahanan lepas", "tahanan lapas, kabur dengan mobdin"},
	{2, "Tahun baru sepi", "tahun baru disurabaya dikabarkan sepi"},
}

func main() {
	http.HandleFunc("/articles", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			var result, err = json.Marshal(data)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
				return
			}
			rw.Write(result)
			return
		}
		http.Error(rw, "", http.StatusBadRequest)
	})
	fmt.Println("Starting at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
