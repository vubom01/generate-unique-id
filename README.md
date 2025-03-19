Generate Unique Id
====
Follow package: https://github.com/bwmarrin/snowflake

### ID Format
* The ID as a whole is a 63 bit integer stored in an int64
* 39 bits are used to store a timestamp with millisecond precision, using a custom epoch.
* 1 bit is used to distinguish the type of ID.
* 10 bits are used to store a node id - a range from 0 through 1023.
* 13 bits are used to store a sequence number - a range from 0 through 8191.
```
+------------------------------------------------------------------------------+
| 1 Bit Unused | 39 Bit Timestamp | Bit 1 | 10 Bit NodeID | 13 Bit Sequence ID |
+------------------------------------------------------------------------------+
```

## Getting Started

### Installing

```sh
go get github.com/vubom01/generate-unique-id
```

### Usage
**Example Program:**
```go
package main

import (
	"fmt"

	generator "github.com/vubom01/generate-unique-id"
)

func main() {
	node, err := generator.NewUniqueIDGenerator(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := node.GenerateID()
	fmt.Println(id.Int64())
	fmt.Println(id.Base2())
	fmt.Println(id.String())
}
```