package main

import (
	"golang.org/x/exp/maps"
)

func main() {
	maps.Keys(map[string]string{"hello": "world"})
}
