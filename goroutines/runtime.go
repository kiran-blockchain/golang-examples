package main

import ("fmt"
"time"
)

func main(){
	 fmt.Println("I am main function")
	 go trySomething()
	 go trySomething()
	 go trySomething2()
	 
	 fmt.Println(time.Second)
	 time.Sleep(10*time.Second)
	 fmt.Println("I am Awake");
}
func trySomething(){
 	fmt.Println("I am a TrySome thing")
}

func trySomething2(){
	fmt.Println("I am a TrySome thing2")
}
