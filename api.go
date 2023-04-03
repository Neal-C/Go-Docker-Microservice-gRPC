//lint:file-ignore ST1006 heeh...

package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"github.com/Neal-C/Go-Docker-Microservice-gRPC/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error



type JSONAPIServer struct {
	listenAddr string
	service PriceFetcher
}

func NewJSONAPIServer(listenAddr string, service PriceFetcher) *JSONAPIServer {
	
	return &JSONAPIServer{
		listenAddr: listenAddr,
		service: service,
	}
}

func (self *JSONAPIServer) Run(){
	http.HandleFunc("/", makeHTTPHandlerFunc(self.handleFetchPrice));
	http.ListenAndServe(self.listenAddr, nil);
}

func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {

	ctx := context.Background();
	ctx = context.WithValue(ctx, "requestID", rand.Intn(1000))


	return func(w http.ResponseWriter, r *http.Request){
		if err := apiFn(ctx,w,r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			});
		}
	}
}

func (self *JSONAPIServer) handleFetchPrice(ctx context.Context, responseWriter http.ResponseWriter, request *http.Request) error {
	ticker := request.URL.Query().Get("ticker");

	price, err := self.service.FetchPrice(ctx, ticker);

	if err != nil {
		return err;
	}

	priceResponse := types.PriceResponse{
		Price: price,
		Ticker: ticker,
	};

	return writeJSON(responseWriter, http.StatusOK, priceResponse);
}

func writeJSON(responseWriter http.ResponseWriter, status int, v any) error {
	responseWriter.WriteHeader(status);
	return json.NewEncoder(responseWriter).Encode(v);
}