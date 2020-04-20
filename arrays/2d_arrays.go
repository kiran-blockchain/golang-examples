package main
/*
arra.forEach((x,i)=>{

});
*/

import "fmt"

func main() {
	/* an array with 3 rows and 3 columns*/
	var a = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for index, value := range a {
		fmt.Println(index, value)
		for index2, value2 := range value {
			fmt.Println(index2, value2)
		}
	}
	// var i, j int
	// /* output each array element's value */
	// for i = 0; i < 3; i++ {
	// 	for j = 0; j < 3; j++ {
	// 		fmt.Print(a[i][j])
	// 	}
	// 	fmt.Println()
	// }
}
