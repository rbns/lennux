package main

import (
	"log"
	"math/rand"
	"net/http"
)

var urls = []string{"http://www.fedoraproject.org",
	"http://www.redhat.com",
	"http://www.suse.com",
	"http://www.opensuse.org",
	"http://www.archlinux.org",
	"http://www.debian.org"}

func rotate(w http.ResponseWriter, r *http.Request) {
	idx := rand.Intn(len(urls))
	http.Redirect(w, r, urls[idx], http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", rotate)
	log.Fatal(http.ListenAndServe("127.0.0.1:7890", nil))
}
