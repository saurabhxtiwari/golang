// CommandLineArguments
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println(os.Args)

	fmt.Println(len(os.Args))

	for index, val := range os.Args {
		fmt.Println(strconv.Itoa(index))
		fmt.Println(val)
	}

}
