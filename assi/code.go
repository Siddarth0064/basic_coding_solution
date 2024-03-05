package main

import (
	"fmt"
	"hash/fnv"
	"net/http"
)

var url map[uint32]string

func main() {
	url := make(map[uint32]string)
	fmt.Println(url)
	http.HandleFunc("/shorten", shortenUrl)
	http.HandleFunc("/redirect", redirectUrl)
	http.ListenAndServe(":8080", nil)

}
func shortenUrl(w http.ResponseWriter, r *http.Request) {
	longUrl := r.FormValue("longurl")
	if longUrl == "" {
		http.Error(w, "please provide a long URl", http.StatusBadRequest)
		return
	}
	hash := hashURL(longUrl)
	shortUrl := fmt.Sprintf("/redirect?hash=%d", hash)
	url[hash] = longUrl
	fmt.Fprintf(w, "short URL : http://localhost:8080%s", shortUrl)

}
func redirectUrl(w http.ResponseWriter, r *http.Request) {
	hashstr := r.FormValue("hash")
	if hashstr == "" {
		http.Error(w, "short URL not provided", http.StatusBadRequest)
		return
	}
	var hash uint32
	fmt.Sscanf(hashstr, "%d", &hash)
	longURL, ok := url[hash]
	if !ok {
		http.Error(w, "short url not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
func hashURL(url string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(url))
	return h.Sum32()
}
