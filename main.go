package main

import (
	"flag"
)

func main() {
	port := flag.String("listenAddr", ":3000", "port");
	flag.Parse();

	service := NewLoggingService(NewMetricService(&priceFetcher{}));

	server := NewJSONAPIServer(*port, service);

	server.Run();

}