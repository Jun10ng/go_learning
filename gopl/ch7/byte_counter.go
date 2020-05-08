//byte_counter

package ch7

// ByteCounter 字节数
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	// 虽然都是int型的别名，但是golang并不支持隐式转换
	*c += ByteCounter(len(p))
	return len(p), nil
}
