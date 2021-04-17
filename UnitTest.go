package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	curl := exec.Command("curl", "www.baidu.com")
	out, err := curl.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(out))
	}
}
