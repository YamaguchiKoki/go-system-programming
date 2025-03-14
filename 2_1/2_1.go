package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "%d\t%b\t%x\n", 65, 65, 65)
}
