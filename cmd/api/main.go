package main

import (
	"log"
	"patriciabonaldy/lana/cmd/api/bootstrap"
)

func main(){
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}