package api

import (
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httpTransport "github.com/go-kit/kit/transport/http"
)

func newHTTPHandler(svc ApiService, endpointFunc func(ApiService) endpoint.Endpoint, decodeFunc httpTransport.DecodeRequestFunc, encodeFunc httpTransport.EncodeResponseFunc) http.Handler {
	return httpTransport.NewServer(
		endpointFunc(svc), decodeFunc, encodeFunc)
}

func MakeHTTPHandler(svc ApiService, allowedOrigins []string) {
	mux := http.NewServeMux()

	corsHandler := CORSMiddleware(allowedOrigins)

	mux.Handle("/user", corsHandler(newHTTPHandler(svc, makeGetUserByIdEndpoint, decodeGetUserByIdRequest, encodeResponse)))
	mux.Handle("/user/id", corsHandler(newHTTPHandler(svc, makeDeleteUserEndpoint, decodeDeleleUserByIdRequest, encodeResponse)))

	log.Fatal(http.ListenAndServe(":5000", mux))
}
