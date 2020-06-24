package main

import (
	"net/http"
	"os"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}


func checkAuth(r *http.Request) bool {
	h := r.Header.Get("X-AuthKey")
	if authkey == "" || h != authkey {
		return false
	}
	return true
}
