package main

import (
	"fmt"
	"time"
)

// 메인의 역할은 Go 루틴을 돌리는 것인데, go 실행시키면 할일이 끝남,
// 그러면 메인 종료, 그러면 메인해서 호출했던 go 루틴도 종료, 그러면 그냥 끝나는겨
// 혹은 다른 코드로 넘어가는데 go 루틴에서 필요한 값이 있으면 기다려야지
// <- chan, 채널에서 값이 가져올 때 까지 기다림
func main() {
	c := make(chan string) // 데이터를 받아오는 채널

	people := [5]string{"A", "B", "C", "D", "E"}

	for _, person := range people {
		go isSexy(person, c) // 들어온 메세지 까지만 쓸 수 있음, 약간 큐 -> 반복문으로 꺼내오기
		
	}
	
	fmt.Println("Waiting messages ...")
	for i:= 0; i < len(people); i++ {
		fmt.Println("Waiting for", i)
		fmt.Println(<-c)
	}

}


func pCount(person string) {
	for i:=0 ; i< 10; i++ {
		fmt.Println(person, "is", i)
		time.Sleep(time.Second)
	}
}

func isSexy (person string, c chan string) {
	time.Sleep(time.Second)
	c <- person + " is sexy"
}