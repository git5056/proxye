package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

func run_server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		v := r.Form["img"]
		if len(v) > 0 {		
			rawurl := strings.TrimPrefix(r.URL.RawQuery, "img=")
			//rawurl := strings.Trim(v[0], "")
			fmt.Println(rawurl)
			requestImg <- rawurl
			for {
				//<-noticeChan
				if b, ok := errMap[rawurl]; ok && b{
					errMap[rawurl] = false
					w.WriteHeader(500)
					return
				}

				if item, ok := resultMap[rawurl]; ok {
					defer delete(resultMap, rawurl)
					b := make([]byte, item.Len())
					item.Read(b)
					w.Header().Set("Content-Type", "image/"+strings.TrimPrefix(path.Ext(rawurl), "."))
					w.WriteHeader(200)
					w.Write(b)
					return
				}
				time.Sleep(time.Millisecond * 50)
			}
		}
		w.WriteHeader(404)
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
