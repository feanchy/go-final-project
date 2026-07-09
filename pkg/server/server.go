package server

import "net/http"

func Run() {
	http.NewServeMux()
	http.FileServer(.http.Dir(./))
}
