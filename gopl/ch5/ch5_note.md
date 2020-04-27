# ch5_note

## 函数声明

如果一组形参或返回值有相同的类型，我们不必为每个形参都写出参数类型。下面2个声明是等价的：

```go
func f(i, j, k int, s, t string)                 { /* ... */ }
func f(i int, j int, k int,  s string, t string) { /* ... */ }
```

#### 函数返回值

当函数不需要返回值是，可以为空。

当需要返回值时，有两种写法，效果一样。

如果一个函数将所有的返回值都显示的变量名，那么该函数的return语句可以省略操作数。这称之为bare return。

```go
func add(x ,y int) int {
	return x+y
}

func add(x,y int)(z int){
	z = x+y
	return 
}
```

在函数调用时，**Go语言没有默认参数值**，也没有任何方法可以通过参数名指定形参，因此形参和返回值的变量名对于函数调用者而言没有意义。



（重要）**实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。**

你可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以Go实现的。这样的声明定义了函数标识符。

```go
package math

func Sin(x float64) float //implemented in assembly language
```

## 错误

#### 错误处理策略

1. 传播错误

   这意味着函数中某个子程序的失败，会变成该函数的失败。把流程中某个子函数的错误“传播”给主流程函数，并中断。

   ```go
   res, err := subFunc(arg)
   if err != nil{
   	return nil, err
   }
   ```

   这样的错误返回也可以包装下

   ```go
   if err != nil {
   	return nil, fmt.Errorf(" %s : %v", arg,err)
   }
   ```

   可以额外了解下 go 1.14的 unwarp error 和 %w 占位符。这样使用传播错误可以使错误信息模拟调用过程，呈链式。当错误最终由main函数处理时，错误信息应提供清晰的从原因到后果的因果链。由于错误信息经常是以链式组合在一起的，所以错误信息中应**避免大写和换行符**。最终的错误信息可能很长，我们可以通过类似grep的工具处理错误信息（译者注：grep是一种文本搜索工具）。

2. 重试

   如果错误的发生时偶然性的（类似TCP的部分错误处理），那么我们可以采用重试的策略，但是要注意重试的时间和次数，防止无限制的重试。**在所有重试之后如果还是失败的话，再返回错误**。

   

   ```go
   // WaitForServer 尝试连接url参数对应的服务器
   // 它持续一分钟的重连，并采用指数级的等待时间增加
   // 如果所有重试都失败了，将返回错误
   func WaitForServer(url string) error {
   	const timeout = 1 * time.Minute
   	deadline := time.Now().Add(timeout)
   	for tries := 0; time.Now().Before(deadline); tries++ {
   		_, err := http.Head(url)
   		if err == nil {
   			return nil // success
   		}
   		log.Printf("server not responding (%s);retrying…", err)
   		time.Sleep(time.Second << uint(tries)) // 指数递增
   	}
   	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
   }
   ```

   

3. 第一时间结束

   这个策略一般用在 main 文件中，当主流程遇到错误时，直接退出结束程序。

   这个策略一般与 错误传播 策略合用，将子函数传播至主流程中，然后依照错误的严重性判断是否结束程序。

   ```go
   // 在主程序中
   if err := WaitForServer(url); err != nil {
       log.Fatalf("wrong %v\n",err)
   	os.Exit(1)
   }
   ```

4. 只输出错误信息，不中断

   就是调用日志系统，常见于一些小问题，如图片丢失等。

5. 忽略错误

   

```go
dir, err := ioutil.TempDir("", "scratch")
if err != nil {
	return fmt.Errorf("failed to create temp dir: %v",err)
}
// ...use temp dir…
os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically
```

尽管os.RemoveAll会失败，但上面的例子并没有做错误处理。这是因为操作系统会定期的清理临时目录。正因如此，虽然程序没有处理错误，但程序的逻辑不会因此受到影响。**我们应该在每次函数调用后，都养成考虑错误处理的习惯，当你决定忽略某个错误时，你应该清晰地写下你的意图。**

### 函数值/闭包

函数可以被赋值给一个变量（函数值），但是初次赋值后，这个函数值的函数类型就固定了。

```go
函数类型 = 参数类型+返回值类型
func square(n int) int { return n * n } //func(int) int
```

```go
	func square(n int) int { return n * n }
	func negative(n int) int { return -n }
	func product(m, n int) int { return m * n }

	f := square
	fmt.Println(f(3)) // "9"

	f = negative
	fmt.Println(f(3))     // "-3"
	fmt.Printf("%T\n", f) // "func(int) int"

	f = product // compile error: can't assign func(int, int) int to func(int) int
```

**函数类型的零值是nil。**,函数可以与nil比较，但是不能在函数之间比较，所以也不能作为map的key。

此外函数值也可作为函数的参数传入，类型为函数类型。

```go
func mainFunc(n int, subFunc func(n int)){
    
}
```

### 匿名函数

```go
func mainFunc(func(n int)){
	
}
```

更为重要的是，通过这种方式定义的函数可以访问完整的词法环境（lexical environment），这意味着在函数中定义的内部函数可以引用该函数的变量，如下例所示：

```go
// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
```

函数squares返回另一个类型为 func() int 的函数。对squares的一次调用会生成一个局部变量x并返回一个匿名函数。每次调用匿名函数时，该函数都会先使x的值加1，再返回x的平方。第二次调用squares时，会生成第二个x变量，并返回一个新的匿名函数。新匿名函数操作的是第二个x变量。

squares的例子证明，函数值不仅仅是一串代码，**还记录了状态**。在squares中定义的匿名内部函数可以访问和更新squares中的局部变量，这意味着匿名函数和squares中，存在变量引用。这就是函数值属于引用类型和函数值不可比较的原因**。Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包。**

### 使用闭包实现拓扑排序

，考虑这样一个问题：给定一些计算机课程，每个课程都有前置课程，只有完成了前置课程才可以开始当前课程的学习；我们的目标是选择出一组课程，这组课程必须确保按顺序学习时，能全部被完成。每个课程的前置课程如下：

```go
// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}
```

使用dfs，但是golang的dfs实现有点不一样，如果使用闭包的话。

```go

import (
	"fmt"
	"sort"
)

func f(p map[string][]string) {
	for i,c :=range topoSort(p){
		fmt.Printf("%d:\t%s\n", i+1, c)
	}
}

//拓扑排序
func topoSort(m map[string][]string) []string  {
	var order []string
	visited := make(map[string]bool)

	//其实并没有运行，只是定义声明了visitAll函数而已
	var visitAll  func(items []string)
	visitAll = func(items []string){
		for _,item := range items{
			if !visited[item] {
				visited[item] = true
				visitAll(m[item])
				//递归入栈
				order = append(order,item)
			}
		}
	}

	//topo排序真正的流程
	var keys []string
	for key:=range m{
		keys = append(keys,key)
	}
	//保证输出顺序固定
	sort.Strings(keys)
	visitAll(keys)
	return order
}
```

