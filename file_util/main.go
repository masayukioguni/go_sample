package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	contensts,err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(contensts,err)
		return 
	}

	for i:=0; i< len(contensts); i++ {
		fmt.Println(string(contensts[i]))
	}

	fmt.Println(string(contensts))
}
