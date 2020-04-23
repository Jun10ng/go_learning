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

1. 一个slice的元素是间接引用的，一个slice甚至可以包含自身。虽然有很多办法处理这种情形，但是没有一个是简单有效的。
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

我一直不清楚什么叫 **切片有可包含它自己** ，于是我就谷歌了，以下是我的整理。

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

```
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

