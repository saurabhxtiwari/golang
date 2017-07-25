// EnvironmentVariables
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("MYGOENV", "MYGOENVVALUE")
	fmt.Println(os.Getenv("MYGOENV"))

	for _, val := range os.Environ() {
		fmt.Println(val)
		pair := strings.Split(val, "=")
		fmt.Println(pair[0])
		fmt.Println(pair[1])
	}
}
