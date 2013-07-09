package main

import (
	"github.com/ChuckHa/blog/blog"
	"github.com/ChuckHa/blog/entry"
	"fmt"
)

func main() {
	blogs := blog.List()
	for _, blog := range blogs {
		fmt.Printf("========= %s ==========\n", blog.Title)
		for _, entry := range blog.GetEntries() {
			fmt.Printf("++++++++ %s ++++++++\n", entry.Title)
			fmt.Printf("%s wrote this on %s\n\n%s\n",
				entry.Author, entry.Created, entry.Content)
		}
	}

	fmt.Println("======= List of Entries ========")

	entries := entry.List()
	for _, entries := range entries {
		fmt.Println(entries.Title)
	}
}
