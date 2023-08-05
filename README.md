# go-counter

## Overview

Returns the frequency of each element in a string slice.

## Usage

### Installation

```
go get github.com/maeda6uiui/go-counter
```

### Code sample

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

