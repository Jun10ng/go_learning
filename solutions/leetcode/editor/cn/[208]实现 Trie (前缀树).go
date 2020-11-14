package main

//实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。
//
// 示例: 
//
// Trie trie = new Trie();
//
//trie.insert("apple");
//trie.search("apple");   // 返回 true
//trie.search("app");     // 返回 false
//trie.startsWith("app"); // 返回 true
//trie.insert("app");   
//trie.search("app");     // 返回 true 
//
// 说明: 
//
// 
// 你可以假设所有的输入都是由小写字母 a-z 构成的。 
// 保证所有输入均为非空字符串。 
// 
// Related Topics 设计 字典树 
// 👍 432 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	name string
	subTrie map[string]*Trie
	//加一个tag，来标记，是否为一个完整的单词
	isWord bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		name:"",
		subTrie:make(map[string]*Trie,0),
		isWord:false,
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	root := this
	for _,v := range word {
		value,ok := root.subTrie[string(v)]
		if ok {
			root = value
		}else {
			newNode := make(map[string]*Trie,0)
			root.subTrie[string(v)] = &Trie{
				name :string(v),
				subTrie:newNode,
				isWord:false,
			}
			root = root.subTrie[string(v)]
		}
	}
	root.isWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	isWord,ok := this.searchbyWord(word)
	if ok && isWord {
		return true
	}
	return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	_,ok :=this.searchbyWord(prefix)
	return ok
}

func (this *Trie) searchbyWord(word string) ( bool, bool) {

	var (
		root = this
		value *Trie
		ok bool
	)
	for _,v := range word {
		value,ok = root.subTrie[string(v)]
		//fmt.Println(value)
		if ok {
			root = value
			continue
		}
		return false ,false
	}
	return value.isWord, true
}

// func (this *Trie) match(re string ) bool{
//     for k,v := range this.subTrie {
//         if v.name == re {
//             return true
//         }
//     }

//     return false
// }

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */


//leetcode submit region end(Prohibit modification and deletion)
