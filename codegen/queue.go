package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	tt := template.Must(template.ParseFiles("./templates/queue.tpl"))
	for i := 1; i < len(os.Args); i++ {
		dest := strings.ToLower(os.Args[i]) + "_queue.go"
		file, err := os.Create(dest)
		if err != nil {
			fmt.Printf("Could not create %s: %s (skip)\n", dest, err)
			continue
		}

		vals := map[string]string{
			"MyType":  os.Args[i],
			"Package": "main",
		}
		tt.Execute(file, vals)

		file.Close()
	}
}
