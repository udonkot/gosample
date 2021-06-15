package main

import (
	"fmt"
	"net/http"

	"github.com/common-nighthawk/go-figure"
)

func handler(w http.ResponseWriter, r *http.Request) {
	myFigure := figure.NewFigure("Hello World", "", true)
	//	myFigure.Print()
	result := myFigure.Slicify()
	myFigure.Print()
	fmt.Fprint(w, result)
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
