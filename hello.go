package main
import (
        "fmt"
        "net/http"
        "os"
)
func handler(w http.ResponseWriter, r *http.Request) {
        h, _ := os.Hostname()
        fmt.Fprintf(w, "Hello FASHION CLOUD , I'm served from %s!", h)
        mongodbUrl := os.Getenv("MONGODB_URL")
}
func main() {
        http.HandleFunc("/", handler)
        http.ListenAndServe(":80", nil)
}