## 包

每个Go程序都是由包组成的。
程序运行的入口是包main。
按照惯例，包名与导入路径的最后一个目录一致。例如，"math/rand"包由package rand语句开始

导入包时可以一个一个导入,如下:

```
import "fmt"
import "math"
```

也可以用圆括号组合导入，谠是"打包"导入语句，打包导入语句是更好的形式。

```
import (
    "fmt"
    "math"
)
```

在Go中，首字母大写的名称是被导出的。在导入包之后，你只能访问所导出的名字，任何未导出的名字是不能被包外的代码访问的。

Foo和FOO都是被导出的名称。名称foo是不会被导出的。

## 
