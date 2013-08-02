package main

import (
	"index/suffixarray"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	sometext, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		log.Println(err)
	}

	index := suffixarray.New(sometext)

	offsets := index.Lookup([]byte("*"), -1)

  for _, value := range offsets[0:100] {
    log.Println(value, string(sometext[value]))
  }
}
