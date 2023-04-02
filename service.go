//lint:file-ignore ST1006 heeh...

package main

import (
	"context"
	"fmt"
	"time"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{
	//
}

func (self *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error){
	return MockPriceFetcher(ctx, ticker);

}

var priceMocks = map[string]float64{
	"BTC": 28_000.0,
	"ETH": 3000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error){

	time.Sleep(time.Millisecond * 100);
	price, ok := priceMocks[ticker];
	
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker);
	}

	return price, nil;
}