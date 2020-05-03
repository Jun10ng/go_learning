# ch11_note

Go语言的测试技术是相对低级的。它依赖一个**go test测试命令和一组按照约定方式**编写的测试函数，测试命令可以运行这些测试函数。编写相对轻量级的纯测试代码是有效的，而且它很容易延伸到基准测试和示例文档。

## go test

在`*_test.go`文件中，有三种类型的函数：测试函数、基准测试（benchmark）函数、示例函数。go test命令会遍历所有的`*_test.go`文件中符合上述命名规则的函数，生成一个**临时的main包**用于调用相应的测试函数，接着构建并运行、报告测试结果，最后清理测试中生成的**临时文件**。

### 测试函数

一个测试函数是以**Test为函数名前缀**的函数，用于测试程序的一些**逻辑行为**是否正确，go test命令会调用这些测试函数并报告测试结果是PASS或FAIL。

```go
func TestName(t *testing.T) {
	// ...
}
```

其中`t`参数用于报告测试失败和附加的日志信息。

```go
// 判断一个字符串是否前后相等
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
```

在测试文件中：

```go
// 测试前后相等字符串
func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error("IsPalindrome(\"detartrated\") = false")
	}
	if !IsPalindrome("kayak"){
		t.Error("IsPalindrome(\"kayak\") = false")
	}

}

// 测试前后不相等字符串
func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}
```

使用 `t.Error`或 `t.Errorf `报告失败信息，如果执行到这一个语句，就会报错。提醒用户。

**先编写测试用例并观察到测试用例触发了和用户报告的错误相同的描述是一个好的测试习惯。只有这样，我们才能定位我们要真正解决的问题。**

参数`-v`可用于打印每个测试函数的**名字和运行时间**：

```
$ go test -v
=== RUN TestPalindrome
--- PASS: TestPalindrome (0.00s)
=== RUN TestNonPalindrome
--- PASS: TestNonPalindrome (0.00s)
=== RUN TestFrenchPalindrome
--- FAIL: TestFrenchPalindrome (0.00s)
    word_test.go:28: IsPalindrome("été") = false
=== RUN TestCanalPalindrome
--- FAIL: TestCanalPalindrome (0.00s)
    word_test.go:35: IsPalindrome("A man, a plan, a canal: Panama") = false
FAIL
exit status 1
FAIL    gopl.io/ch11/word1  0.017s
```

参数`-run`对应一个正则表达式，只有测试函数名被它正确匹配的测试函数才会被`go test`测试命令运行，可以用这个参数来跑特定参数，比如 `go test -v -run="French|Canal"`

**当然，一旦我们已经修复了失败的测试用例，在我们提交代码更新之前，我们应该以不带参数的`go test`命令运行全部的测试用例，以确保修复失败测试的同时没有引入新的问题。**

#### 表格测试

我们把之前的测试例子合在一个“表格”中：

```go
func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
```

这种表格驱动的测试在Go语言中很常见，我们可以很容易地向表格添加新的测试数据，且后面的测试逻辑也没有冗余，这样我们可以有更多的精力去完善错误信息。

#### 随机测试

表格驱动的测试便于构造基于精心挑选的测试数据的测试用例。另一种测试思路是随机测试，也就是通过构造更广泛的随机输入来测试探索函数的行为。

那么对于一个随机的输入，我们如何能知道希望的输出结果呢？这里有两种处理策略。

1. 是编写另一个对照函数，使用简单和清晰的算法，虽然效率较低但是行为和要测试的函数是一致的，然后针对相同的随机输入检查两者的输出结果。
2. 生成的随机输入的数据遵循特定的模式，这样我们就可以知道期望的输出的模式。

下面的例子使用的是第二种方法：`randomPalindrome` 函数用于随机生成回文字符串。

#### 编写有效的测试

Go语言的极简测试风格与其他语言的测试框架形成鲜明对比。它期望测试者自己完成大部分的工作，定义函数避免重复，就像普通编程那样。编写测试并不是一个机械的填空过程；一个测试也有自己的接口，尽管它的维护者也是测试仅有的一个用户。一个好的测试不应该引发其他无关的错误信息，它只要清晰简洁地描述问题的症状即可，有时候可能还需要一些上下文信息。在理想情况下，维护者可以在不看代码的情况下就能根据错误信息定位错误产生的原因。一个好的测试不应该在遇到一点小错误时就立刻退出测试，它应该尝试报告更多的相关的错误信息，因为我们可能从多个失败测试的模式中发现错误产生的规律。

### 基准测试

**基准测试**函数是以**Benchmark**为函数名前缀的函数，它们用于衡量一些函数的**性能**，go test命令会多次运行基准测试函数以计算一个平均的执行时间。

### 示例函数

示例函数是以**Example**为函数名前缀的函数，提供一个**由编译器保证正确性**的示例文档。