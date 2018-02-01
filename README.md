# TwiML

## Installation

You can use the SDK using the `go` command.

    $ go get github.com/Abhishek-Nagarkoti/TwiML

You can also install by cloning this repository into your `GOPATH`.

### Generate TwiML XML

```go
package main

import "github.com/Abhishek-Nagarkoti/TwiML"

func main()  {
  println(TwiML.ResponseElement{
    Contents: []interface{}{
      new(TwiML.SayElement).SetContents("Hello, world!"),
    },
    }.String())
}
```

This generates the following XML:

```xml
<Response>
  <Say>Hello, world!</Say>
</Response>
```

