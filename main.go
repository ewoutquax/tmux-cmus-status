package main

import (
	"fmt"
	"os"
)

func main() {
	ptrFile, _ := os.OpenFile("/tmp/cmus.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	ptrFile.Write([]byte(fmt.Sprintf("os.Args: %v\n", os.Args)))
	ptrFile.Close()
}
