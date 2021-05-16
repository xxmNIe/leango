package main

import "net/http"
type str struct {

}
type DB map[int]string
func  abc(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("abv"))
}
func  bcv(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("bcv"))
}
func main() {
	//1
	http.HandleFunc("/",abc)
	http.ListenAndServe(":8080",nil)

	//2
	//mux := http.NewServeMux()
	//mux.Handle("/a",http.HandlerFunc(abc))
	//mux.Handle("/b",http.HandlerFunc(bcv))
	////not found handler  but no err
	////http.Handle("/abc",mux)
	//http.ListenAndServe(":8080",mux)


}
