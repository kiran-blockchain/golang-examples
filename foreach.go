package main
import "fmt"
func main(){
	a := []string{"My", "Name","Is","Go"}
	a := [][]string

	for index,value := range a {
		fmt.Println(index, value)
	}
}

