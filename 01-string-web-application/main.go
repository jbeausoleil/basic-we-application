package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const portNumber = 8080

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This is my home page")
	if err != nil {
		return
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	sum, _ := addValues(2, 2)
	_, err := fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
	if err != nil {
		return
	}
}

func Divide(w http.ResponseWriter, r *http.Request) {
	var x float32 = 100.0
	var y float32 = 10.0

	f, err := divideValues(x, y)
	if err != nil {
		fmt.Println(err)
	}
	_, err = fmt.Fprintf(w, fmt.Sprintf("This is the divide page, and the result of "+
		"%f divided by %f is %f", x, y, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	return x / y, nil
}

// Lowercase names scope the names to the current package (invisible to other packages)
func addValues(x, y int) (int, error) {
	sum := x + y
	return sum, nil
}

func main() {
	// Listen for a request sent by a web function
	http.HandleFunc("/", Home)
	http.HandleFunc("/test", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("staring application on port %d", portNumber))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", portNumber), nil))
}
