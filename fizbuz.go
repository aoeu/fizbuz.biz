package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	pathError = "Path parameter must be a number."
	rangeError = "Number must be from 1 to 100."
)


func handler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[1:]
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprint(w, pathError)
		return
	}
	if n < 1 || n > 100 {
		fmt.Fprint(w, rangeError)
		return
	}
	f, b := n%3 == 0, n%5 == 0
	if f || b {
		s = ""
	}
	if f {
		s = "Fizz"
	}
	if b {
		s += "Buzz"
	}
	fmt.Fprintf(w, s)
}

func main() {
	args := struct{ port string }{}
	flag.StringVar(&args.port, "port", "80", "The port to serve on.")
	flag.Parse()
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(args.port, nil); err != nil {
		log.Fatal(err)
	}
}
