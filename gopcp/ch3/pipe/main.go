package main

import (
	"bufio"
	"log"
	"os"
)

/*
使用管道实现在命令输入数据，实时存储到文件内
管道只有两端都有值的情况下才会执行，否则会阻塞
所以必须并发调用。
*/

func main() {
	rp, wp, _ := os.Pipe()

	go func() {
		file, err := os.OpenFile("/home/aiden/golang/go_learning/gopcp/ch3/pipe/out.txt", os.O_RDWR, 0666)
		if err != nil {
			log.Fatal("can not open file")
		}
		data := make([]byte, 100)
		for {
			n, _ := rp.Read(data)
			file.Write(data[:n])
			file.WriteString("\n")
			data = data[n:]
		}
	}()
	cin := bufio.NewScanner(os.Stdin)
	for cin.Scan() {
		data := []byte(cin.Text())
		wp.Write(data)
	}

}
