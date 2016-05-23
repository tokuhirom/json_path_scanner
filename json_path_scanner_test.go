package json_path_scanner

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestInt(t *testing.T) {
	ch := make(chan *PathValue)
	go func() {
		Scan(3, ch)
	}()

	p := <-ch
	log.Print(p)
	if p.Path != "$" {
		t.Fatalf("Path should be $ but %s", p.Path)
	}
	if p.Value != 3 {
		t.Fatalf("Value should be 3 but %v", p.Value)
	}
}

func TestFloat(t *testing.T) {
	ch := make(chan *PathValue)
	go func() {
		Scan(3.14, ch)
	}()

	p := <-ch
	log.Print(p)
	if p.Path != "$" {
		t.Fatalf("Path should be $ but %s", p.Path)
	}
	if p.Value != 3.14 {
		t.Fatalf("Value should be 3.14 but %v", p.Value)
	}
}

func TestString(t *testing.T) {
	ch := make(chan *PathValue)
	go func() {
		Scan("Foo", ch)
	}()

	p := <-ch
	log.Print(p)
	if p.Path != "$" {
		t.Fatalf("Path should be $ but %s", p.Path)
	}
	if p.Value != "Foo" {
		t.Fatalf("Value should be 'Foo' but %v", p.Value)
	}
}

func TestArray(t *testing.T) {
	ch := make(chan *PathValue)
	go func() {
		Scan([]interface{}{5963, 4649}, ch)
	}()

	p := <-ch
	log.Print(p)
	if p.Path != "$[0]" {
		t.Fatalf("Path should be $[0] but %s", p.Path)
	}
	if p.Value != 5963 {
		t.Fatalf("Value should be 5963 but %v", p.Value)
	}

	p = <-ch
	if p.Path != "$[1]" {
		t.Fatalf("Path should be $[1] but %s", p.Path)
	}
	if p.Value != 4649 {
		t.Fatalf("Value should be 4649 but %v", p.Value)
	}
}

func TestMap(t *testing.T) {
	ch := make(chan *PathValue)
	go func() {
		m := make(map[string]interface{})
		m["hoge"] = "fuga"
		Scan(m, ch)
	}()

	p := <-ch
	if p.Path != "$.hoge" {
		t.Fatalf("Path should be $.hoge but %s", p.Path)
	}
	if p.Value != "fuga" {
		t.Fatalf("Value should be 'fuga' but %v", p.Value)
	}
}

func ExampleSynopsys() {
	ch := make(chan *PathValue)
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
		Scan(m, ch)
	}()

	for p := range ch {
		fmt.Printf("%s => %v\n", p.Path, p.Value)
	}

	// Unordered  output:
	// $.hoge => fuga
	// $.x[0].y => 3
	// $.x[0].z[0] => 1
	// $.x[0].z[1] => 2
	// $.x[0].z[2] => 3
}
