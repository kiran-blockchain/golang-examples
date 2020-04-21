/* Go string is a sequence of variable-width characters
UTF-8 is the standard, Go doesn't need to encode and decode strings
Go Strings are value types and immutable
The initial value of a string is empty "" by default
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	//String length
	str := "I love my country"
	fmt.Println("\nString length:", len(str))

	//Print ASCII
	fmt.Println("\nAscii value of A is ", "A"[0])

	//String ToUpper
	fmt.Println("\nString ToUpper:", strings.ToUpper(str))

	//String ToLower
	fmt.Println("\nString ToLower:", strings.ToLower(str))

	str1 := "INDIA"
	//String HasPrefix()
	fmt.Println("\nString HasPrefix:", strings.HasPrefix(str1, "IN"))

	//String HasSuffix()
	fmt.Println("\nString HasPrefix:", strings.HasSuffix(str1, "IA"))

	//String Join()
	var strArry = []string{"a", "b", "c", "d"}
	fmt.Println("\nString Join:", strings.Join(strArry, "*"))

	//String Repeat()
	var newstr = "New "
	fmt.Println("\nString Repeat:", strings.Repeat(newstr, 4))

	//String Contains()
	string1 := "Hi...there"
	fmt.Println("\nString Contains:", strings.Contains(string1, "th"))

	//String Index()
	fmt.Println("\nString Index:", strings.Index(str, "th"))

	//String Count()
	fmt.Println("\nString Count:", strings.Count(str, "e"))

	//String Replace()
	fmt.Println("\nString Replace:", strings.Replace(str, "e", "Z", 2))

	//String Split()
	fmt.Println("\n\nString Split:")
	str2 := "I,love,my,country"
	var arr []string = strings.Split(str2, ",")
	fmt.Println(len(arr))
	for i, v := range arr {
		fmt.Println("Index : ", i, "value : ", v)
	}
	fmt.Println("\n\n")
	fmt.Printf("%q\n", strings.Split("x,y,z", ","))
	fmt.Printf("%q\n", strings.Split(" John and Jack and Johnny and Jinn ", "and"))
	fmt.Printf("%q\n", strings.Split(" abc ", ""))
	fmt.Printf("%q\n", strings.Split("", "Hello"))

	//String Compare()
	fmt.Println("\nString Compare:")
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))

	//String Trim()
	fmt.Println("\nString Trim:")
	fmt.Println(strings.TrimSpace(" \t\n I love my country  \n\t\r\n"))

	//String ContainsAny()
	fmt.Println("\nString ContainsAny:")
	fmt.Println(strings.ContainsAny("Hello", "A"))
	fmt.Println(strings.ContainsAny("Hello", "o & e"))
	fmt.Println(strings.ContainsAny("Hello", ""))
	fmt.Println(strings.ContainsAny("", ""))

}
