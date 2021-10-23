package main

import (
	"fmt"
	"sort"
	"strings"
)

var (
	words = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
)

func main() {
	list := make(map[string][]string)

	for _, word := range words {
		key := sortStr(word)
		list[key] = append(list[key], word)
	}

	for _, words := range list {
		fmt.Print(words, " ")
		fmt.Println()
	}
}

func sortStr(k string) string {
	s := strings.Split(k, "")
	sort.Strings(s)

	return strings.Join(s, "")
}
