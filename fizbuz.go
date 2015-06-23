package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/fcgi"
	"strconv"
)

const (
	errorMsg = "Path parameter must be a number from 1 to 100."
)

func handler(w http.ResponseWriter, r *http.Request) {
	if s, ok := cache[r.URL.Path[1:]]; ok {
		fmt.Fprint(w, s)
		return
	}
	http.Error(w, errorMsg, http.StatusNotFound)
}

var cache map[string]string

func main() {
	args := struct{ port string }{}
	flag.StringVar(&args.port, "port", "", "The port to serve on.")
	flag.Parse()
	cache = make(map[string]string, 100)
	for i := 0; i < 101; i++ {
		s := strconv.Itoa(i)
		t := s
		f, b := i%3 == 0, i%5 == 0
		if f || b {
			t = ""
		}
		if f {
			t = "Fizz"
		}
		if b {
			t += "Buzz"
		}
		cache[s] = t
	}
	http.HandleFunc("/", handler)
	switch {
	case args.port != "":
		if err := http.ListenAndServe(args.port, nil); err != nil {
			log.Fatal(err)
		}
	default:
		if err := fcgi.Serve(nil, nil); err != nil {
			log.Fatal(err)
		}
	}
}
