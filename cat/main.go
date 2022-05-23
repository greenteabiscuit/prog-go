package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fp, err := os.Open("./a.xml")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
