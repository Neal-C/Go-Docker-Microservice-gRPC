//lint:file-ignore ST1006 hmm...

package client

import (
	"context"
	"net/http"
	// "net/url"
	"encoding/json"
	"fmt"
	"github.com/Neal-C/Go-Docker-Microservice-gRPC/types"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (self *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error){

	endpoint := fmt.Sprintf("%s?ticker=%s", self.endpoint, ticker);
	// endpoint2 := url.Parse(self.endpoint).Query().Add("ticker", ticker).String()

	request, err := http.NewRequest("GET", endpoint, nil);

	if err != nil {
		return nil, err;
	};

	response, err := http.DefaultClient.Do(request);

	if err != nil {
		return nil, err;
	}

	if response.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(response.Body).Decode(&httpErr); err != nil{
			return nil, err;
		}
		return nil, fmt.Errorf("service responded with a non ok status code : %s", httpErr["error"]);
	}

	var priceResponse types.PriceResponse;

	if err := json.NewDecoder(response.Body).Decode(&priceResponse); err != nil{
		return nil, err;
	}

	return &priceResponse, nil;

}