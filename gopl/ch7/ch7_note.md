
# ch7_note
<!-- TOC -->

- [ch7_note](#ch7_note)
    - [fmt.Fprintf](#fmtfprintf)
    - [T 与 *T](#t-与-t)
    - [接口的比较](#接口的比较)
        - [如何获得接口类型信息](#如何获得接口类型信息)
        - [包含nil指针的接口和nil接口](#包含nil指针的接口和nil接口)
    - [重要接口介绍](#重要接口介绍)
        - [排序接口](#排序接口)
    - [http.Handler接口](#httphandler接口)
        - [简化url与hander的分配关系](#简化url与hander的分配关系)
    - [断言](#断言)
        - [使用接口来判别错误类型](#使用接口来判别错误类型)
        - [使用类型断言判断行为](#使用类型断言判断行为)

<!-- /TOC -->
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

**这十分巧妙**通过接口赋值我们还可以限制一个具体类暴露出来的方法，这可以用来完成权限设置。

```go
os.Stdout.Write([]byte("hello")) // OK: *os.Stdout实现了Writer接口
os.Stdout.Close()                // OK: *os.Stdout实现了Closer接口

var w io.Writer
w = os.Stdout            // 用Writer接口的w变量赋值
w.Write([]byte("hello")) // OK: 有Writer接口
w.Close()                // compile error: Closer接口已经被屏蔽掉了,close方法自然被屏蔽了
```

## 接口的比较

接口值可以使用==和!＝来进行比较。两个接口值相等仅当它们都是nil值，或者它们的动态类型相同并且动态值也根据这个动态类型的==操作相等。因为接口值是可比较的，所以它们可以用在map的键或者作为switch语句的操作数。

然而，如果两个接口值的动态类型相同，但是这个动态类型是不可比较的（比如切片），将它们进行比较就会失败并且panic:

```go
var x interface{} = []int{1, 2, 3}
fmt.Println(x == x) // panic: comparing uncomparable type []int

```
考虑到这点，接口类型是非常与众不同的。其它类型要么是安全的可比较类型（如基本类型和指针）要么是完全不可比较的类型（如切片，映射类型，和函数），但是在比较接口值或者包含了接口值的聚合类型时，我们必须要意识到潜在的panic。同样的风险也存在于使用接口作为map的键或者switch的操作数。只能比较你非常确定它们的动态值是可比较类型的接口值。

### 如何获得接口类型信息

- 使用fmt包的 %T 占位符
    ```go
    var w io.Writer
    fmt.Printf("%T\n", w) // "<nil>"
    w = os.Stdout
    fmt.Printf("%T\n", w) // "*os.File"
    w = new(bytes.Buffer)
    fmt.Printf("%T\n", w) // "*bytes.Buffer"
    ```

- 反射
    见十二章

### 包含nil指针的接口和nil接口

**警告！**
一个不包含任何值的nil接口值和一个刚好包含nil指针的接口值是不同的。这个细微区别产生了一个容易绊倒每个Go程序员的陷阱。

```go
const debug = true

func main() {
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
```

按照我们的预计，把debug设置成false，那么buf就会为nil，则f被调用的时候，out被赋值也是nil，但是**事实**，buf并不是nil，而是一个包含nil指针的接口，所以out永远都不等于nil。
以buf的io.Writer接口为例，他由两字段组成，type和value，`bytes.Buffer`实现了Writer接口，当我们声明 `var buf *bytes.Buffer` 时，type就被赋值为 `*byte.Buffer`(一个指针了)，而value才是nil，所以整体不为nil。

那么什么才是空接口呢？`var buf io.Writer`,这样以最基础的接口声明的变量，是一个空接口，所以上述代码应该改成。

```go
var buf io.Writer
if debug {
	buf = new(bytes.Buffer) // enable collection of output
}
f(buf) // OK
```

## 重要接口介绍

### 排序接口

sort包内置的提供了根据一些排序函数来对任何序列排序的功能。它的设计非常独到。
在很多语言中，排序算法都是和序列数据类型关联，同时排序函数和具体类型元素关联。
而Go语言的sort.Sort函数不会对具体的序列和它的元素做任何假设。
它使用了一个接口类型sort.Interface来指定通用的排序算法和可能被排序到的序列类型之间的约定。这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。

```go
package sort

type Interface interface {
	Len() int
	Less(i, j int) bool // i, j are indices of sequence elements
	Swap(i, j int)
}
```

我们需要定义一个排序类型的切片，这个实现了这三个方法，然后再对这个类型的一个实例调用sort.Sort函数。
以下为对一个字符串切片进行排序

```go
//切片
type StringSlice []string
//
func (p StringSlice) Len() int{
       return len(p) 
}
//
func (p StringSlice) Less(i, j int) bool {
     return p[i] < p[j]
}
//
func (p StringSlice) Swap(i, j int){ 
    p[i], p[j] = p[j], p[i] 
}
```

对字符串切片是很常用的需要，所以sort包提供了**StringSlic**类型，也提供了sort.Strings(name)

**反向排序，奇妙的实现方法** 不需要重新定义一个Less方法，只需要使用

```go
sort.Sort(sort.Reverse(StringSlice(name)))
```

Reverse的底层实现是定义了一个不公开的类，成员是第六章里讲的那个匿名成员
reverse的Less方法调用了内嵌的sort.Interface值的Less方法，但是通过交换索引的方式使排序结果变成逆序。
reverse的另外两个方法Len和Swap隐式地由原有内嵌的sort.Interface提供。

```go
type reverse struct{ Interface }  // sort.Interface

func (r reverse) Less(i, j int) bool { return r.Interface.Less(j, i) }

func Reverse(data Interface) Interface { return reverse{data} }
```

难道按每个字段排序，只能对比一个字段吗？而且每次都要实现实现3个方法吗？下面的例子利用结构体来做到了只需要重写Less方法即可，本质上用的是go中的函数为一等公民（函数值）的特点。

```go

type customSort struct {
    t    []*Track
    //只需要传入不同的less函数就可以做到不同的排序
	less func(x, y *Track) bool
}

func (x customSort) Len() int {
	return len(x.t)
}
func (x customSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}
func (x customSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

```


```go
//调用方法，传入了一个匿名函数
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {//先比较Title
			return x.Title < y.Title
		}
		if x.Year != y.Year {//再比较Year
			return x.Year < y.Year
		}
		if x.Length != y.Length {//再比较Length
			return x.Length < y.Length
        }
        //如果上面三个都相等
		return false
	}})
```

sort包为[]int、[]string和[]float64的正常排序提供了特定版本的函数和类型。对于其他类型，例如[]int64或者[]uint，尽管路径也很简单，还是依赖我们自己实现。

## http.Handler接口


```go
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
```

ListenAndServe函数需要一个例如“localhost:8000”的服务器地址，和一个所有请求都可以分派的Handler接口实例。它会一直运行，**直到这个服务因为一个错误而失败**（或者启动失败），它的返回值一定是一个非空的错误。
所以，我们只需要对一个实例实现Handler接口即可，ListenAndServe会自动调用h的ServeHttp方法。

```go
func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
```
高级版,考虑了多个不同的URL

```go
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
        msg := fmt.Sprintf("no such page: %s\n", req.URL)
        http.Error(w, msg, http.StatusNotFound) // 404
	}
}
```
### 简化url与hander的分配关系

net/http包提供了一个请求多路器ServeMux来简化URL和handlers的联系。一个ServeMux将一批http.Handler聚集到一个单一的http.Handler中。再一次，我们可以看到满足同一接口的不同类型是可替换的：web服务器将请求指派给任意的http.Handler而不需要考虑它后面的具体类型。

在下面的程序中，我们创建一个ServeMux并且使用它将URL和相应处理/list和/price操作的handler联系起来，这些操作逻辑都已经被分到不同的方法中。然后我们在调用ListenAndServe函数中使用ServeMux为主要的handler。

```go
func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
```

使用mux的Handle方法时，第二个参数是http.HandlerFunc(db.list)，db.list并不满足http.handle接口，所以不能直接传给mux.Handle。
语句http.HandlerFunc(db.list)是一个转换而非一个函数调用，因为http.Handle是一个类型，而不是一个方法。他体现了Go的语言特性，因此HandlerFunc是一个**让函数值满足一个接口**的适配器。

```go
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

在一个应用程序的多个文件中定义HTTP handler也是非常典型的，如果它们必须全部都显式地注册到这个应用的ServeMux实例上会比较麻烦。

所以为了方便，net/http包提供了一个全局的ServeMux实例DefaultServerMux和包级别的http.Handle和http.HandleFunc函数。现在，为了使用DefaultServeMux作为服务器的主handler，我们不需要将它传给ListenAndServe函数；nil值就可以工作。

主函数可简化为：

```go
func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
```

## 断言


类型断言有两种方式。他们的代码都是`x.(T)`，无论哪种情况下，**x都必须是一个接口！先声明接口，再给x赋值**

第一种，**如果T是一个具体类型**，那么，`f = x.(T)` 就是把一个接口x转化成T具体类型。但是前提是T实现了接口x。如：

```go
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)      // success: f == os.Stdout
	c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
```

第二种与第一种相反，**T是一个接口类型**，那么, x.(T)则会返回x的T接口的动态类型,f中不在具备x内除了T接口外的其他方法，

```go
var w io.Writer
w = os.Stdout
f := w.(*os.File)      // success: f == os.Stdout
c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
```

这两种都是吧x转换成其子集，（子类型或者子接口）
断言失败一般会导致panic的，所以在断言前，我们需要判断断言**一个变量是否为特定类型**，这也可以用断言来实现

```go
if w, ok := w.(*os.File); ok {
	// ...use w...
}
```

### 使用接口来判别错误类型

在os包内，有三个方法，他们被用来判断错误的类型，其底层实现是判断错误中是否包含特定的字符串。


```go
package os

func IsExist(err error) bool
func IsNotExist(err error) bool
func IsPermission(err error) bool
```

使用案例，打开一个文件，期间发生了错误，我们要判断这个错误是什么？然后交给程序尝试自主解决。

```go
_, err := os.Open("/no/such/file")
fmt.Println(os.IsNotExist(err)) // "true"
```

### 使用类型断言判断行为

类型断言还可以用来判断某个类有没有特定的方法。
比如我们需要判断一个类型是否有writeString方法，我们可以像下面这样写：在函数内创建一个临时的接口，然后用类型断言，如果失败的话再说。**这是一种很巧妙的做法，灵活的使用到了接口和断言。**

```go
// writeString writes s to w.
// If w has a WriteString method, it is invoked instead of w.Write.
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) // avoid a copy
	}
	return w.Write([]byte(s)) // allocate temporary copy
}
```
