//lint:file-ignore ST1006 heeh...

package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceFetcher
}

func NewLoggingService(service PriceFetcher) PriceFetcher {
	return &loggingService{
		next: service,
	}
}

func (self *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error){
	defer func(begin time.Time){

		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"took": time.Since(begin),
			"err": err,
			"price": price,
		}).Info("fetchPrice")
	}(time.Now());

	return self.next.FetchPrice(ctx, ticker);
}