package server

import "net/http"

type CORSHandler struct {
	Next http.Handler
}

func (handler CORSHandler) ServeHTTP(out http.ResponseWriter, in *http.Request) {
	origin := in.Header.Get("Origin")
	if origin == "" {
		handler.Next.ServeHTTP(out, in)
		return
	}

	header := out.Header()
	header.Set("Access-Control-Allow-Origin", origin)
	header.Set("Access-Control-Allow-Credentials", "true")

	if in.Method == http.MethodOptions {
		header.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,HEAD")
		header.Set("Access-Control-Allow-Headers", "authorization,content-type,content-length")
		header.Set("Access-Control-Max-Age", "86400")
		return
	}

	handler.Next.ServeHTTP(out, in)
}
