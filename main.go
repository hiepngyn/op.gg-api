package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	testurl := "https://op.gg/lol/champions/jinx/build/adc"
	body, err := Fetch(context.Background(), testurl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("got %d bytes\n", len(body))
	//os.WriteFile("jinx-adc.html", body, 0644)
}
