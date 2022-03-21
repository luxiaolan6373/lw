调用乐玩插件 

安装

```go
go get github.com/luxiaolan6373/lw
```

封装了70%的命令，其它的请有志之士帮忙提交吧。。依葫芦画瓢就行了

编译和调试前记得将环境改成 32位的

```go
 go env -w GOARCH=386 
```

这是一个列子非常简单

```go
package main

import (
	"github.com/luxiaolan6373/lw"
)
func main() {
    // 调用例子
	com, err := lw.New()
	if err != nil {
		panic("注册失败")
		return
	} else {
        //打印版本号 所有的函数都可以用com.lw.xx的方式调用
		print(com.lw.Ver())
	}
	
}


```

