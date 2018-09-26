package main

import (
	"fmt"
	"net/http"
)

func main() {

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World!")
}
