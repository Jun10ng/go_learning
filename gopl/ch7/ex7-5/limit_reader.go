package ch7

import "io"

/* 严重怀疑这题不是实现func LimitReader，而是实现Type LimitReader，但是做法差不多吧，就是有点别扭 */

// limitReader 类是用来做LimitReader方法的内部类
type limitReader struct {
	r        io.Reader
	n, limit int
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	lrRead, lrErr := lr.r.Read(p[:lr.limit])
	lr.n += lrRead
	// 如果读的字节超过n个，则返回EOF
	if lr.n > lr.limit {
		lrErr = io.EOF
	}
	return lr.n, lrErr
}

// LimitReader 返回一个只读前n个字节的Reader
func LimitReader(r io.Reader, n int64) io.Reader {
	lr := &limitReader{
		r:     r,
		n:     0,
		limit: int(n),
	}
	return lr

}
