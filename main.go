package main

import (
	"fmt"

	//	"github.com/gansidui/ahocorasick"
	"github.com/cloudflare/ahocorasick"
)

func main() {
	//	ac := ahocorasick.NewMatcher()

	dictionary := []string{"hello", "world", "hello google", "世界", "google", "golang", "c++", "love"}

	//	ac.Build(dictionary)
	ac := ahocorasick.NewStringMatcher(dictionary)

	ret := ac.Match([]byte("whee google hello世界, hello google, i love golang!!!"))

	for _, i := range ret {
		fmt.Println(dictionary[i])
	}
}
