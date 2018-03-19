# Ref : https://jdanger.com/build-a-web-crawler-in-go.html

package main

import {
	"fmt"
	"flag"
	"os"
}

func main(){
	flag.Parse()
	args := flag.Args()

	if len(args) > 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}

	retrieve(args[0])
}

func retrieve(uri string){
	resp, err := http.Get(uri)

	if err != nil {
		fmt.Println("Not able to read uri")
		return nil
	}

}
