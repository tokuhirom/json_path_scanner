package main

import (
	"encoding/json"
	"fmt"
	"github.com/tokuhirom/json_path_scanner"
	"log"
)

func main() {
	ch := make(chan *json_path_scanner.PathValue)
	go func() {
		var m interface{}
		err := json.Unmarshal([]byte(`{
                "hoge":"fuga",
                "x":[
                    {
                        "y": 3,
                        "z": [1,2,3]
                    }
                ]
            }`), &m)
		if err != nil {
			log.Fatal(err)
		}
		json_path_scanner.Scan(m, ch)
	}()

	for p := range ch {
		fmt.Printf("%s => %v\n", p.Path, p.Value)
	}
}
