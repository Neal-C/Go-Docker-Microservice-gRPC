//lint:file-ignore ST1006 heeh...

package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{
		next: next,
	}
}

func (self *metricService) FetchPrice(ctx context.Context, ticker string) (float64, error){
	//Metrics go here. Push to Prometheus/TerraForm, whatever
	fmt.Println("metrics were pushed");
	return self.next.FetchPrice(ctx, ticker);
}
