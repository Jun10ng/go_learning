# go 语言圣经第四章读书笔记

## Slice

和数组不同的是，**slice之间不能比较**，因此我们不能使用==操作符来判断两个slice是否含有全部相等元素。

不过标准库提供了高度优化的 `bytes.Equal`  函数来判断两个字节型 slice 是否相等`（[]byte）`，但是对于其他类型的 slice ，我们必须自己展开每个元素进行比较：

```go
func equal(x, y []string) bool {
    if len(x) != len(y) {
        return false
    }
    for i := range x {
        if x[i] != y[i] {
            return false
        }
    }
    return true
}
```

而 `bytes.Equal` 函数源码是其实就是把两个字节数组转成 `string` 相比较而已。

```go
func Equal(a, b []byte) bool {
	// Neither cmd/compile nor gccgo allocates for these string conversions.
	return string(a) == string(b)
}
```

### 为什么不支持两个 slice 相比呢？

1. 一个slice的元素是间接引用的，一个slice甚至可以包含自身（当 slice类型是 `[]interface{}`)。虽然有很多办法处理这种情形，但是没有一个是简单有效的。
2. 第二个原因，因为slice的元素是间接引用的，一个固定的slice值(译注：指slice本身的值，不是元素的值)在不同的时刻可能包含不同的元素，因为底层数组的元素可能会被修改。而例如Go语言中map的key只做简单的浅拷贝，它要求key在整个生命周期内保持不变性(译注：例如slice扩容，就会导致其本身的值/地址变化)。而用深度相等判断的话，显然在map的key这种场合不合适。对于像指针或chan之类的引用类型，==相等测试可以判断两个是否是引用相同的对象。一个针对slice的浅相等测试的==操作符可能是有一定用处的，也能临时解决map类型的key问题，但是slice和数组不同的相等测试行为会让人困惑。因此，安全的做法是直接禁止slice之间的比较操作。

slice唯一合法的比较操作是和nil比较，例如：

```go
if summer == nil { /* ... */ }
```

*注：slice的零值为 nil*

#### slice 包含它自身是什么意思

在 gopl 中原文是

*First, unlike array elements, the elements
of a slice are indirect, making it possible for **a slice to contain itself**. Although there are ways to deal with such cases, none is simple, efficient, and most importantly, obvious.*

我一直不清楚什么叫 **切片有可能包含它自己** ，于是我就谷歌了，以下是我的整理。

* 证明一：

  ```go
  //证明一
  	//这段代码运行会报 stack overflow
  	//因为是s[2]是个自身，不断取值，不断循环
  	//这也证明了slice是可以指向自身的，
  	//但是不能被打印出来
  	s := []interface{}{"one",2,nil}
  	s[2] = s
  	fmt.Printf("s is %v\n",s)
  	fmt.Printf("s[2] is %v",s[2])
  ```

* 证明二

```go
//证明二
	//这个例子说明了 s2为s的第三个元素，
	//s2不但是个切片而且，前两个值和s相同
	//避开了会导致溢出的第三个元素
	//用另一种方式证明了切片能包含它自身
	s:=[]interface{}{"one","two",nil}
	s[2] = s
	s2 := s[2].([]interface{})
	fmt.Printf("s2[0] is %v, s2[1] is %v \n",s2[0],s2[1])
	s3 := s2[2].([]interface{})
	fmt.Printf("s3[0] is %v, s3[1] is %v \n",s3[0],s2[1])
        // output:
        //"one","two"
        //"one","two"
```

#### 为什么数组就不会包含自身

```go
	//当s被放入[]inteface{}时，会创建一个副本
	//所以已经不是自身了
	s := [2]interface{}{"one", nil}
	s[1] = s
	fmt.Printf("s is %v \n",s) //output:[one [one <nil>]]
	fmt.Printf("s address is %p ,while s[1] address is %p.\n",&s,&s[1])
	// output: s address is 0xc000004560 ,while s[1] address is 0xc000004570.
	//明显看得出两个地址并不一样
	fmt.Println(s[0])
	s2 := s[1].([2]interface{})
	fmt.Println(s2[0])
	s3 := s2[1].([2]interface{})
	fmt.Println(s3)
```

因此用interface{} 来承接一个其他数据类型的数组也是一个**深拷贝数组**的奇门巧技。

### slice如何比较

- `reflect`比较的方法

  ```go
  func StringSliceReflectEqual(a, b []string) bool {
      return reflect.DeepEqual(a, b)
  }
  ```

- 循环遍历比较的方法



### 测试 slice 是否为空

一个零值的 `slice `等于 `nil` 。一个`nil`值的`slice`并没有底层数组。一个 `nil` 值的`slice`的长度和容量都是 0 ，但是也有非`nil`值的`slice`的长度和容量也是0的，例如 `[]int{}` 或 `make([]int, 3) [3:]`。与任意类型的nil值一样，我们可以用 `[]int(nil)` 类型转换表达式来生成一个对应类型`slice`的  `nil`值。

```go
var s []int    // len(s) == 0, s == nil
s = nil        // len(s) == 0, s == nil
s = []int(nil) // len(s) == 0, s == nil
s = []int{}    // len(s) == 0, s != nil
```

**如果你需要测试一个slice是否是空的，使用`len(s) == 0`来判断，而不应该用`s == nil`来判断。**

除了和nil相等比较外，一个nil值的slice的行为和其它任意0长度的slice一样；例如reverse(nil)也是安全的。除了文档已经明确说明的地方，所有的Go语言函数应该以相同的方式对待nil值的slice和0长度的slice。

内置的make函数创建一个指定元素类型、长度和容量的slice。**容量部分可以省略，在这种情况下，容量将等于长度。**

```go
make([]T, len)
make([]T, len, cap) // same as make([]T, cap)[:len]
```

### slice 深拷贝

使用 `copy` 可以深拷贝切片

```
copy(z, x) //把 x 的内容 拷贝至 z , 将返回成功复制的元素的个数
```



### slice 的 append 详解

https://github.com/gopl-zh/gopl-zh.github.com/blob/master/ch4/ch4-02-1.md



### 用 slice 实现栈



**push**

```go
stack = append(stack, v) // push v
```

**top**

```go
top := stack[len(stack)-1] // top of stack
```

**pop**

收缩切片

```go
stack = stack[:len(stack)-1] // pop
```

**remove**

要删除slice中间的某个元素并保存原有的元素顺序，可以通过内置的copy函数将后面的子slice向前依次移动一位完成：

```go
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
}
```



## Map

* 在Go语言中，一个map就是一个哈希表的**引用**，注意，go 中并不存在引用变量。可以看下 [这篇翻译文章](https://www.cnblogs.com/Jun10ng/p/12725669.html)

* 其中K对应的**key必须是支持==比较运算符的数据类型**，所以map可以通过测试key是否相等来判断是否已经存在。

* map中的元素并不是一个变量，因此我们不能对map的元素进行**取址操作**。禁止对map元素取址的原因是map可能随着元素数量的增长而**重新分配**更大的内存空间，从而可能**导致之前的地址无效**。

### 为什么 map 的遍历是随机的（两次遍历结果不相同）

 [我写的这篇博客](https://www.cnblogs.com/Jun10ng/p/12771077.html) 有解释



### 判断两个 map 是否相等

和slice一样，map之间也**不能进行相等比较**；唯一的**例外是和nil进行比较**。要判断两个map是否包含相同的key和value，我们必须通过一个循环实现

`x`,`y` 都是 map，要判断两个是否相等，如何呢？

不能简单的使用 `x[k] == y[k]`判断，因为 如果 `y`中并不存在 `k`键，那么 `y[k]`返回的是零值（不是nil），此时如果 `x[k]`的值刚好就是零值的话，就会返回真，但其实 `y`甚至连 `k`都没有。

正确做法应当先判断 `y[k]`是否存在

```go
if yv, ok := y[k]; ok{
    //再处理与 x[k]的比较
}
```

### 用map实现set

```
map[string]bool
```

## Struct



### 初始化

以下三个语句是相等的

```go
pp := &Point{1, 2}
pp := new(Point)
*pp = Point{1, 2}
```



### 结构体不能包含自身类型的成员，但是可以包含类型指针

一个命名为S的结构体类型将不能再包含S类型的成员：因为一个**聚合体**不能包含它自身。（该限制同样适用于数组。）但是S类型的结构体可以包含`*S`指针类型的成员，这可以让我们创建递归的数据结构，比如**链表和树结构**等。

### 使用go mod导入自定义包

文件结构如下：

```
test_code
	│  go.mod
	│
	├─p
	│      p_struct.go
	│
	└─q
            q_test.go
            q_use.go

```

各文件内容

```go
// p_struct.go
package p

type Point struct{
	X int
	Y int
}

//q_use.go
package q

import (
	"fmt"
	"test_code/p"
)

func q_use()  {
	p := &p.Point{1,2}
	fmt.Printf("x:%v\ty:%v",p.X,p.Y)
}


//go.mod
//最重要的文件、
module test_code

go 1.14

```

要在 q_use.go 文件中使用 p_struct.go 中的结构体。

在 go.mod 文件中，将 test_code声明为一个模块，然后在 q_use中就可以导入 test_code 下的 p_struct文件了。

注意了，一个项目中，不同文件夹都可以有go.mod ，就行一个大工程是很多小模块组成的。

### 结构体的函数传值与返回

在go中，尽量使用指针来传递结构体。有以下两个好处。

1. 指针比结构体本身小，使得效率高
2. go中，都是值传递，也就是拷贝体，所以如果希望目标结构体被处理修改，请使用指针传值。

### 结构体的比较

如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，直接使用 `==``和！=`。所以可比较的结构体可以用于 `map`的 `key`类型。**不可比较的数据类型包括，Slice, Map, 和函数**

### 匿名成员

为了避免，结构体套结构体，而导致的**过多间接访问**。如：

```go
w.Circle.Center.X = 8
w.Circle.Center.Y = 8
```

Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫**匿名成员**。匿名成员的数据类型必须是**命名的类型或指向一个命名的类型的指针**。下面的代码中，Circle 和Wheel 各自都有一个匿名成员。我们可以说Point类型被嵌入到了Circle结构体，同时Circle类型被嵌入到了Wheel结构体。

得益于匿名嵌入的特性，我们可以**直接访问叶子属性而不需要给出完整的路径**：

```
type Circle struct {
	Point   //匿名成员
	Radius int
}

type Wheel struct {
	Circle  ////匿名成员
	Spokes int
}

var w Wheel
w.X = 8            // equivalent to w.Circle.Point.X = 8
w.Y = 8            // equivalent to w.Circle.Point.Y = 8
w.Radius = 5       // equivalent to w.Circle.Radius = 5
w.Spokes = 20
```

匿名成员的作用，**简单来说就是方便，加快了效率，让代码更直观了**。专业的说法是： **对访问嵌套成员的点运算符提供了简短的语法糖**

## JSON

JavaScript对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。在类似的协议中，JSON并不是唯一的一个标准协议。 XML（§7.14）、ASN.1和Google的Protocol Buffers都是类似的协议，并且有各自的特色，但是由于简洁性、可读性和流行程度等原因，JSON是应用最广泛的一个。

### 结构体转 JSON

这样的数据结构特别适合JSON格式，并且在两者之间相互转换也很容易。将一个Go语言中类似movies的结构体slice转为JSON的过程叫**编组（marshaling）**。编组通过调用json.Marshal函数完成：

```go
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}
```

编组代码

```go
data, err := json.Marshal(movies)
if err != nil {
	log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
```

Marshal函数返回一个编码后的字节slice，包含很长的字符串，并且没有空白缩进；我们将它折行以便于显示：

```
[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingr
id Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Ac
tors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"
Actors":["Steve McQueen","Jacqueline Bisset"]}]
```

为了生成**便于阅读**的格式，另一个**json.MarshalIndent**函数将产生整齐缩进的输出。该函数有两个额外的字符串参数用于表示每一行输出的前缀和每一个层级的缩进：

```go
data, err := json.MarshalIndent(movies, "", "	")
```



```json
[
	{
		"Title": "Casablanca",
		"released": 1942,
		"Actors": [
			"Humphrey Bogart",
			"Ingrid Bergman"
		]
	},
	{
		"Title": "Cool Hand Luke",
		"released": 1967,
		"color": true,
		"Actors": [
			"Paul Newman"
		]
	},
	{
		"Title": "Bullitt",
		"released": 1968,
		"color": true,
		"Actors": [
			"Steve McQueen",
			"Jacqueline Bisset"
		]
	}
]
```

### 结构体成员 Tag

其中Year名字的成员在编码后变成了released，还有Color成员编码后变成了小写字母开头的color。这是因为结构体成员Tag所导致的。一个结构体成员Tag是和在编译阶段关联到该成员的元信息字符串：

```
Year  int  `json:"released"`
Color bool `json:"color,omitempty"`
```

Color成员的Tag还带了一个额外的**omitempty**选项，表示当Go语言结构体成员为空或零值时不生成该JSON对象（这里false为零值）。果然，Casablanca是一个黑白电影，并没有输出Color成员。

### 解码

Go语言中一般叫**unmarshaling**，通过**json.Unmarshal**函数完成。下面的代码将JSON格式的电影数据解码为一个结构体slice，结构体中只有Title成员。通过定义合适的Go语言数据结构，我们可以**选择性**地解码JSON中感兴趣的成员。当Unmarshal函数调用返回，slice将被只含有Title信息的值填充，其它JSON成员将被忽略。

以下代码将 把 movies 切片中所有元素的 Title 成员解码，注意，Unmarshal函数只返回一个error，**真正的数据要用指针数值的形式传值，见第二个参数**

```
var titles []struct{ Title string }
if err := json.Unmarshal(data, &titles); err != nil {
	log.Fatalf("JSON unmarshaling failed: %s", err)
}
fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
```

## 文本和HTML模板

```go
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`
```

如上代码中，"."代表当前值，当前值“.”最初被初始化为调用模板时的参数，在当前例子中对应`github.IssuesSearchResult`类型的变量。

在一个action中，`|`操作符表示将前一个表达式的结果作为后一个函数的输入，类似于UNIX中管道的概念。在Title这一行的action中，第二个操作是一个printf函数，是一个基于fmt.Sprintf实现的内置函数，所有模板都可以直接使用。

### 一个简单的Web服务器

监听端口

```go

func server() {
	//启动一个web服务器
	http.HandleFunc("/", handle)
	http.ListenAndServe("0.0.0.0:8000", nil)
}
```

绑定服务

```go
func handle(w http.ResponseWriter, r *http.Request) {
    
    // 核心业务逻辑
	var result *github.IssuesSearchResult
	var keywords = []string{"repo:golang/go","is:open","json","decoder"}
	result, _ = github.SearchIssue(keywords)
    
    
    // 渲染模板
	var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))
    // 把result 写进 issueList后，放进 w 后做为response 返回
	issueList.Execute(w, result)
}
```

