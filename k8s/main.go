package main

import (
	"fmt"

	"github.com/prembhaskal/carvel-stuff/k8s/pkglist"
)

func main() {
	fmt.Printf("hello world\n")

	pkglist.ListPackages()
}
