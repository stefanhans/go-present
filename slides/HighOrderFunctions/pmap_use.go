package main
import (. "bitbucket.org/stefanhans/golang-ctx/presentations/HighOrderFunctions/hof"
	"fmt"
	"runtime"
	"time")

func main() {
	tenTimes := func(x int) int { time.Sleep(time.Duration(1 * time.Millisecond))
		return x * 10 }
	var list ListOfInt
	for i := 0; i < 10; i++ { list = append(list, i) }

	start := time.Now()
	fmt.Printf("list%v.Map(tenTimes) yields %v\n", list, list.Map(tenTimes))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Printf("list%v: PMap(tenTimes, %v) yields %v\n", list, runtime.NumCPU(),
		list.PMap(tenTimes, runtime.NumCPU()))
	fmt.Println(time.Since(start))
	fmt.Printf("and list%v is immutable\n", list)

}
