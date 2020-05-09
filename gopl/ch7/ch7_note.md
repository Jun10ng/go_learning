
# ch7_note

## fmt.Fprintf

实际上，`fmt.Printf`和`fmt.Sprintf`都是对`fmt.Fprintf`的封装。

```go
package fmt

func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)
func Printf(format string, args ...interface{}) (int, error) {
    return Fprintf(os.Stdout, format, args...)
}
func Sprintf(format string, args ...interface{}) string {
    var buf bytes.Buffer
    Fprintf(&buf, format, args...)
    return buf.String()
}
```

可以看出Fprintf的第一个参数是 `io.Writer`类型的。
Fprintf的前缀F表示文件（File）也表明格式化输出结果应该被写入第一个参数提供的文件中。在Printf函数中的第一个参数os.Stdout是`*os.File`类型；在Sprintf函数中的第一个参数&buf是一个指向可以写入字节的内存缓冲区，然而它并不是一个文件类型尽管它在某种意义上和文件类型相似。
以下为`io.Writer`类型的接口定义：

```go
package io

// Writer是用来包装Write方法的接口。
type Writer interface {
    // Write 从p中写入长度为len(p)的字节，到下方的数据流
    //它返回写入的字节数，这个数小于等于len(p),大于等于0
    // 任何错触发都会导致write方法提前停止 
    // write方法必须返回 非空错误 如果它返回的n小于len(p) 
    // Write必须不会修改切片信息，哪怕是临时的
    // Write必须不会修改切片信息，哪怕是临时的
    // 实现必须不保留p
    Write(p []byte) (n int, err error)
}

```

Printf和Sprintf之所以能在内部调用Fprintf是因为`*os.File`和`*bytes.Buffer`，这些类型他们都有一个特定签名和行为的Write方法。也就是他们都实现了Writer接口。

`io.Writer`类型是用得最广泛的接口之一了，因为他提供了**所有类型的写入byte的抽象，包括文件类型，内存缓存区，网络连接，HTTP客户端，压缩工具，哈希等**，io包中还有许多有用的接口类型(一般只有一个方法的接口都以er结尾，如只有Write方法的Writer接口)

- Reader可以代表任意可以**读作bytes的类型**

- Closer代表可以**被关闭的值**

```go
package io
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Closer interface {
    Close() error
}
```

## T 与 *T

对于每一个命名过的具体类型T；它的一些方法的接收者是类型T本身，而另一些接受者则是`*T`的指针。
在T类型的变量上调用一个`*T`的方法是合法的，只要这个参数是一个变量；编译器隐式的获取了它的地址。但这仅仅是一个语法糖：T类型的值不拥有所有`*T`指针的方法，这样它就可能只实现了更少的接口。

看下面的例子：

```go
type IntSet struct { /* ... */ }
func (*IntSet) String() string
var _ = IntSet{}.String() // compile error: String requires *IntSet receiver
//while
var s IntSet
var _ = s.String() // OK: s is a variable and &s has a String method
```

**（这十分巧妙）**通过接口赋值我们还可以限制一个具体类暴露出来的方法，这可以用来完成权限设置。

```go
os.Stdout.Write([]byte("hello")) // OK: *os.Stdout实现了Writer接口
os.Stdout.Close()                // OK: *os.Stdout实现了Closer接口

var w io.Writer
w = os.Stdout            // 用Writer接口的w变量赋值
w.Write([]byte("hello")) // OK: 有Writer接口
w.Close()                // compile error: Closer接口已经被屏蔽掉了,close方法自然被屏蔽了
```
