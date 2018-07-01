// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func index_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Whoah, Go is neat!")
// }

// func about_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "lol doing Go and other sht during work!")
// }

// func main() {
// 	http.HandleFunc("/", index_handler)
// 	http.HandleFunc("/about/", about_handler)
// 	http.ListenAndServe(":8000", nil)
// }
