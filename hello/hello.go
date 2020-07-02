package main

import "fmt"

const englishHelloPrefix = "hello, "

func Hello(name string) string{
	if name == ""{
		name = "world"
	}
	return englishHelloPrefix + name
}

func main(){
	fmt.Println(Hello("world"))
}