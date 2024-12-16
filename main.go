package main

import (
	"fmt"
	"runtime"
	"time"
)

func GetWeatherapi() {
	// 创建一个大对象
	largeData := make([]byte, 10*1024*1024) // 10 MB
	fmt.Println(largeData[:100])
	return
}

func main() {
	fmt.Println("hello,world!")
	return
	buffer := make([]byte, 1024*5)
	buffer[0] = 48
	buffer[1] = 49
	buffer[2] = 50
	allData := string(buffer[:100])
	fmt.Println(len(allData))
	fmt.Println(string(buffer[:100]))
	return
	// 调用 GetWeatherapi 函数
	GetWeatherapi()

	// 手动触发垃圾回收
	//runtime.GC()
	// 观察内存使用情况
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys = %v MiB\n", m.Sys/1024/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)
	time.Sleep(time.Minute)
}
