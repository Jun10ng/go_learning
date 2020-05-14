package search

// defaultMathcher 实现了默认匹配器
type defaultMathcher struct {
}

// init 函数将默认匹配器注册到程序里
func init() {
	var matcher defaultMathcher
	Register("default", matcher)
}

// Search 实现了默认匹配器的行为
func (m defaultMathcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
