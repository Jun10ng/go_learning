## 题目列表

### ch4

#### slice

**练习 4.3：** 重写reverse函数，使用数组指针代替slice。

**练习 4.4：** 编写一个rotate函数，通过一次循环完成旋转。

**练习 4.5：** 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。

**练习 4.6：** 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回

**练习 4.7：** 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？

#### map

**练习 4.8：** 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等Unicode中不同的字符类别。

**练习 4.9：** 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。

#### JSON
**练习 4.10：** 修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。

**练习 4.13：** 使用开放电影数据库的JSON服务接口，允许你检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像。编写一个poster工具，通过命令行输入的电影名字，下载对应的海报。

#### 文本和HTML

**练习 4.14：** 创建一个web服务器，查询一次GitHub，然后生成BUG报告、里程碑和对应的用户信息。

#### 函数
**练习 5.1： **修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
**练习 5.5： **实现countWordsAndImages。（参考练习4.9如何分词）
**练习5.19： ** 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。

#### 方法
**练习6.1:** 为bit数组实现下面这些方法

```go
func (*IntSet) Len() int      // return the number of elements
func (*IntSet) Remove(x int)  // remove x from the set
func (*IntSet) Clear()        // remove all elements from the set
func (*IntSet) Copy() *IntSet // return a copy of the set
```

**练习 6.2：** 定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)。
**练习 6.3：** (*IntSet).UnionWith会用`|`操作符计算两个集合的并集，我们再为IntSet实现另外的几个函数IntersectWith（交集：元素在A集合B集合均出现），DifferenceWith（差集：元素出现在A集合，未出现在B集合），SymmetricDifference（并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A）。
**练习6.4: ** 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作。

**练习 6.5：** 我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。当然了，这里我们可以不用简单粗暴地除64，可以定义一个常量来决定是用32还是64，这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)
//这题就是增加一个常量