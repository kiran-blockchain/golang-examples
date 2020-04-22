// package main

// import (
// 	"fmt"
// 	"runtime/debug"
// 	"time"
// )

// func a() {
// 	defer recoverUser()
// 	// n := []int{1, 5, 6}
// 	// fmt.Println("", n[3])
// 	go b()

// 	fmt.Println("I am cool i esccaped")
// 	time.Sleep(5 * time.Second)
// }
// func main() {
// 	//defer recoverUser()
// 	a()
// 	fmt.Println("I am cool i main")
// }

// // func b() {
// // 	defer recoverUser()
// // 	n := []int{1, 5, 6}
// // 	fmt.Println("", n[3])
// // }

// // func main() {
// // 	defer fmt.Println("Handling in main")
// // 	handler("kira", "kiran")
// // }

// // func handler(fname string, lname string) {
// // 	defer recoverUser()
// // 	if fname == "" {
// // 		panic("runtime error: first name cannot be empty")
// // 	}
// // 	if lname == "" {
// // 		panic("runtime error: last name cannot be empty")
// // 	}
// // 	fmt.Println("", fname)
// // 	fmt.Println("", lname)
// // }

// func recoverUser() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Revocered from panic", r)
// 		debug.PrintStack()
// 	}
// }
