package main

import "fmt"


func main()  {
	s:="dlkmfijrejdo"
	for i:=range s[1:len(s)-1]{
		fmt.Println(string(s[i]))
	}
}