package utils

import (
	"fmt"
	"testing"
)

// 单元测试
func TestGet(t *testing.T) {
	body, err := Get("www.baidu.com", 20)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("body = ", body)
	}
}

// 性能测试
func BenchmarkGet(b *testing.B) {
	// b.N会根据函数的运行时间取一个合适的值
	for i := 0; i < b.N; i++ {
		combination(i+1, rand.Intn(i+1))
	}
}

// 并发性能测试
func BenchmarkGetParallel(b *testing.B) {
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m := rand.Intn(100) + 1
			n := rand.Intn(m)
			combination(m, n)
		}
	})
}

/*
go test http_test.go http.go           # 单元测试
go test --cover http_test.go http.go   # 单元测试覆盖率
go test -bench=. http_test.go http.go  # 性能测试

goconvey
 */
