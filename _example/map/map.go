package main

import (
	"fmt"
	"time"

	gtf "github.com/snowlyg/gotransformer"
)

type Response struct {
	Name  string
	Value int
	Rmk   string
}

func main() {

	// map to struct
	response := Response{}
	m := map[string]interface{}{
		"Name":  "name",
		"Value": 1,
		"Rmk":   "Rmk",
	}
	g := gtf.NewTransform(&response, m, time.RFC3339)
	err := g.Transformer()
	if err != nil {
		_ = fmt.Sprintf("err:%v", err)
	}

	_ = fmt.Sprintf("response:%v", response)

}
