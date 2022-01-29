# go-counter

## 概要

Pythonの`collections.Counter`みたいな機能をGolangで実装しました。

## 使い方

```
go get github.com/maeda6uiui/go-counter
```

```go
package main

import (
	"fmt"

	"github.com/maeda6uiui/go-counter"
)

func main() {
	ss := []string{"a", "a", "a", "b", "c", "d", "a", "a", "d", "c"}
	c := counter.NewCounter(ss)

	keys, freqs := c.MostCommon()
	for i := 0; i < len(keys); i++ {
		fmt.Printf("%v %v\n", keys[i], freqs[i])
	}
}
```

```
a 5
d 2
c 2
b 1
```
