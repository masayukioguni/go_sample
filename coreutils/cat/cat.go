package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		wrote_size, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		for _, fname := range os.Args[1:] {
			fd, err := os.Open(fname)
			if err != nil {
				log.Fatal(err)
			}

			wrote_size, err := io.Copy(os.Stdout, fd)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
