package main


import (
	"os"
	"fmt"
)


func main() {
	fmt.Println("Args: %d",len(os.Args))

	for i,e := range os.Args {
		fmt.Printf("%d: %s\n",i,e)
	}
	fmt.Println()

}
