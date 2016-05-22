# json\_path\_scanner
[![CircleCI](https://circleci.com/gh/tokuhirom/json_path_scanner.svg?style=svg)](https://circleci.com/gh/tokuhirom/json_path_scanner)

Scan JSON and returns list of JSON path and value.

## SYNOPSIS

Here's a code:

```
package main

import (
	"encoding/json"
	"fmt"
	"github.com/tokuhirom/json_path_scanner"
	"log"
)

func main() {
	ch := make(chan json_path_scanner.PathValue)
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
		fmt.Printf("%s => %s\n", p.Path, p.Value)
	}
}
```

Output:

    $.hoge => fuga
    $.x[0].z[0] => %!s(float64=1)
    $.x[0].z[1] => %!s(float64=2)
    $.x[0].z[2] => %!s(float64=3)
    $.x[0].y => %!s(float64=3)

## LICENSE

    The MIT License (MIT)
    Copyright © 2016 Tokuhiro Matsuno, http://64p.org/ <tokuhirom@gmail.com>

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the “Software”), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
    THE SOFTWARE.

