package assMakeServer

import "net/http"

type HttpHandler struct{}

func (h HttpHandler) ServeHttp(rw http.ResponseWriter, r *http.Request) {
	data := []byte("Hello I am Mukund")

	rw.Write(data)
}
