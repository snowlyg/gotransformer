## map 数据转换格式
```go
package main

import (
"fmt"
"time"

"github.com/snowlyg/gotransform"
)

// 格式化数据
type Response struct {
	Name         string
    Value        int
    Rmk          string
}


func main()  {

    response := Response{}
     m := map[string]interface{}{
       "Name": "name",
       "Value":1,
       "Rmk":"Rmk",
    }
    g := gotransform.NewTransform(&response, m, time.RFC3339)
    err := g.Transformer()
    if err != nil {
        _ = fmt.Sprintf("err:%v",err)
    }

    _ = fmt.Sprintf("response:%v",response)

}

```