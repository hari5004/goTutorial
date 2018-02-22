package main
import (
"log"
"net/http"
"encoding/json"
)

func handler(w http.ResponseWriter, r*http.Request) {
	//fmt.Fprintf(w,"Hi there, %s!", r.URL.Path[1:])
json.NewEncoder(w).Encode("Hello ")
}
func main() {
http.HandleFunc("/",handler)
log.Fatal(http.ListenAndServe(":8080",nil))
}
