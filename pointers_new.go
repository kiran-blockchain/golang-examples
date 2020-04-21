package main 

import "fmt"

// func  main()  {
// 	a :=100
// 	fmt.Println(a);

// 	addressOfa :=&a

// 	*addressOfa = 20;
// 	fmt.Printf("pointer %T with value %v \n", addressOfa,addressOfa)
// 	fmt.Println(a);
// }

//example 2
func changeValue(p [3]int){
 p[0] =100
 p[1] =200
 p[2] = 400
}

func main (){
	a:=[3]int{1,4,5}
	changeValue(a)
	fmt.Printf("a = %v \n",a)
}