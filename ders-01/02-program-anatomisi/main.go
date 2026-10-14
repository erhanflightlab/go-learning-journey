package main

import (
	"fmt"
	"os"
	"strings"
)

const appName = "anatomy"

var startedBy string

func init() {
	startedBy = os.Getenv("USERNAME")
}

func main() {
	fmt.Printf("%s started by %q\n", appName, startedBy)
	fmt.Println(shout("go is fun"))
}

func shout(s string) string {
	return strings.ToUpper(s) + "!"
}
