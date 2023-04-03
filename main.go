package main

import (
	"flag"
)

func main() {

	// fmt.Println("start")
	// client := client.New("http://127.0.0.1:3000/");
	// fmt.Println("cient created")


	port := flag.String("listenAddr", ":3000", "port");
	flag.Parse();
	
	service := NewLoggingService(NewMetricService(&priceFetcher{}));

	server := NewJSONAPIServer(*port, service);
	server.Run();
	// go server.Run();

	// price, err := client.FetchPrice(context.Background(), "ETH");

	// if err != nil {
	// 	log.Fatal(err);
	// }

	// fmt.Printf("%+v \n", price);
	// fmt.Println("done")

}