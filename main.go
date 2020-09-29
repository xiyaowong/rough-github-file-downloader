package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	rawLinks := []string{}
	for _, arg := range os.Args[1:] {
		if arg != "" {
			rawLinks = append(rawLinks, arg)
		}
	}

	dlLinks := []string{}
	for _, rawLink := range rawLinks {
		new1 := strings.ReplaceAll(rawLink, "github.com", "cdn.jsdelivr.net/gh")
		new2 := strings.ReplaceAll(new1, "blob/master/", "")
		dlLinks = append(dlLinks, new2)
	}

	for _, link := range dlLinks {
		fmt.Println(link)
		cmd := exec.Command("wget", link)
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}
}
