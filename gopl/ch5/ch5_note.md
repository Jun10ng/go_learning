

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

## 闭包包了什么

闭包/Closure 在维基百科中的解释为：

*是引用了自由变量的函数。这个被引用的自由变量将和这个函数一同存在，即使已经离开了创造它的环境也不例外。所以，有另一种说法认为闭包是由函数和与其相关的引用环境组合而成的实体。*

### 语法糖与设计模式

闭包实际上就是一种语法糖机制，而这种语法糖机制可以简化编程，从而可以提高代码的**可读性**。

比如，有时候对外部的局部变量进行访问，没这种语法糖机制将会编写冗余的代码。而这正也是可以把这种闭包机制归结为一种**设计模式**。

这种设计模式的好处就是可以很轻易的访问一个函数的内部的局部变量，再简单点就是通过函数嵌套的方式，闭包可以很轻易的实现函数**内部和外部之间的连接**。

我们知道函数是一段可执行代码，编译后就失效了，而这些函数在编译时就确定了，在执行时，不会发生变化，每个函数在内存中只有一份实例。而闭包并且在执行时候根据不同的引用变量和函数组合起来可以发生变化，也就意味着可以**产生多个实例**。还有一个好处就是函数调用结束时就会自动失效，而闭包的好处就是可以**让这些变量始终保持在内存中**，不会随着函数的调用而消失。

### 函数值

Go语言中不允许函数嵌套定义，但是可以用匿名函数来实现嵌套。在这里就得知道，在Go语言中，函数也是一种类型，这意味着可以把**函数当成一个值**来传递和返回。函数既可以作为一种返回类型又可以作为其他函数的参数。所以，这样**很容易使用函数类型来实现闭包**。

```go
//如果写成
//var visitAll  func(items []string){...} 报错

var visitAll  func(items []string)
	visitAll = func(items []string){//匿名函数
		for _,item := range items{
			if !visited[item] {
				visited[item] = true
				visitAll(m[item])
				//递归入栈
				order = append(order,item)
			}
		}
	}


```

### 闭包包了什么？

包了**函数**和**环境**，

函数指外部函数的**内部函数**。环境指闭包保存/记录了它产生时的**外部函数的所有环境。**

看一个例子：

```
func foo(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}
```

环境就是指内部函数（匿名函数）用到的不属于它的变量`x`。

### 闭包延迟绑定

**在执行的时候去外部环境寻找最新的数值**，

```go
func foo(x int) []func() {
    var fs []func()
    values := []int{1, 2, 3, 5}
    val:=0
    for _, val = range values {
        fs = append(fs, func() {
            fmt.Printf("foo val = %d\n", x+val)
        })
    }
    //val:=11
    return fs
}
// Q4实验：
fs := foo(11)
for _, f := range fs {
    f()
}
```

输出：

```
foo val = 16
foo val = 16
foo val = 16
foo val = 16
```

这就是所谓的 **闭包延迟绑定**，就是指，闭包虽然包了环境，但是包的是最新的值，当真正运行闭包函数的时候，函数再去环境找值，所以采用的上述例子中，闭包函数中的 `val`一直是5。如果我们把函数中注释掉的那句 `val = 11`那么输出的就会是22了。

## defer

你只需要在调用普通函数或方法前加上关键字defer，就完成了defer所需要的语法。

当defer语句被执行时，跟在defer后面的函数会被延迟执行。在该函数执行完毕后，defer语句才会执行（理解为把defer语句**压入栈底**，最后执行，**在return语句更新返回值变量后再执行**）。不论该函数是通过return正常结束，还是由于panic导致的异常结束。你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相反。

defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。通过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。**释放资源的defer应该直接跟在请求资源的语句后**。

连接

```go
 resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
```
文件
```go
 f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()
```
锁
```go
var mu sync.Mutex
var m = make(map[string]int)
func lookup(key string) int {
    mu.Lock()
    defer mu.Unlock()
    return m[key]
}
```

### defer+闭包

对匿名函数采用defer机制，可以使其观察函数的返回值。

```go
func double(x int) (result int) {
    defer func() { fmt.Printf("double(%d) = %d\n", x,result) }()
    return x + x
}
_ = double(4)
```

可能doulbe函数过于简单，看不出这个小技巧的作用，但对于有许多return语句的函数而言，这个技巧**很有用**。

被延迟执行的匿名函数甚至可以**修改**返回值：

```go
func triple(x int) (result int) {
    defer func() { result += x }()
    return double(x)
}
fmt.Println(triple(4)) // "12"
```

### 注意点

#### 循环体的defer

在**循环体中的defer语句**需要特别注意，因为只有在函数执行完毕后，这些被延迟的函数才会执行（而不是循环体结束）。下面的代码会导致系统的文件描述符耗尽，因为在所有文件都被处理之前，没有文件会被关闭。

```go
for _, filename := range filenames {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close() // NOTE: risky; could run out of file
    descriptors
    // ...process f…
}
```

**解决办法**：

将循环体中的defer语句移至另外一个函数。在每次循环时，调用这个函数。

```go
for _, filename := range filenames {
    if err := doFile(filename); err != nil {
        return err
    }
}
func doFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    // ...process f…
}
```

#### 文件系统的延迟错误反馈

许多文件系统会把写入操作的错误延迟到关闭文件时再反馈，这样会导致我们我以为写入操作成功。

```Go
// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", 0, err
    }
    defer resp.Body.Close()
    local := path.Base(resp.Request.URL.Path)
    if local == "/" {
        local = "index.html"
    }
    f, err := os.Create(local)
    if err != nil {
        return "", 0, err
    }
    n, err = io.Copy(f, resp.Body)
    // Close file, but prefer error from Copy, if any.
    if closeErr := f.Close(); err == nil {
        err = closeErr
    }
    return local, n, err
}
```

上例中，通过os.Create打开文件进行写入，在关闭文件时，我们没有对f.close采用defer机制，因为这会产生一些微妙的错误。**许多文件系统，尤其是NFS，写入文件时发生的错误会被延迟到文件关闭时反馈。如果没有检查文件关闭时的反馈信息，可能会导致数据丢失，而我们还误以为写入操作成功。**如果io.Copy和f.close都失败了，我们倾向于将**io.Copy的错误信息反馈给调用者，因为它先于f.close发生，更有可能接近问题的本质**。

## panic

不是所有的panic异常都来自运行时，直接调用内置的panic函数也会引发panic异常；panic函数接受任何值作为参数。当某些不应该发生的场景发生时，我们就应该调用panic。比如，当程序**到达了某条逻辑上不可能到达的路径**：

```
switch s := suit(drawCard()); s {
    case "Spades":                                // ...
    case "Hearts":                                // ...
    case "Diamonds":                              // ...
    case "Clubs":                                 // ...
    default:
        panic(fmt.Sprintf("invalid suit %q", s)) // Joker?
}
```

### Recover

当web服务器遇到不可预料的严重问题时，在崩溃前应该将所有的连接关闭；如果不做任何处理，会使得客户端一直处于等待状态。如果web服务器还在开发阶段，服务器甚至可以将异常信息反馈到客户端，帮助调试。

如果在deferred函数中调用了内置函数recover，并且定义该defer语句的函数发生了panic异常，recover会使程序从panic中恢复，并返回panic value。导致panic异常的函数不会继续运行，但能正常返回。在未发生panic时调用recover，recover会返回nil。

**有时我们很难完全遵循规范**，举个例子，net/http包中提供了一个web服务器，将收到的请求分发给用户提供的处理函数。很显然，我们不能因为某个处理函数引发的panic异常，杀掉整个进程；web服务器遇到处理函数导致的panic时会调用recover，输出堆栈信息，继续运行。这样的做法在实践中很便捷，但也会引起资源泄漏，或是因为recover操作，导致其他问题。

**基于以上原因，安全的做法是有选择性的recover。换句话说，只恢复应该被恢复的panic异常，此外，这些异常所占的比例应该尽可能的低。**为了标识某个panic是否应该被恢复，我们可以将panic value设置成特殊类型。在recover时对panic value进行检查，如果发现panic value是特殊类型，就将这个panic作为errror处理，如果不是，则按照正常的panic进行处理（在下面的例子中，我们会看到这种方式）。