package main

import (
	"github.com/Masterminds/sprig/v3"
	"golang.org/x/exp/maps"
)

func main() {
	maps.Keys(map[string]string{"hello": "world"})
	sprig.GenericFuncMap()
}
