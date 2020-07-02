package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "hello, "
const SpanishHelloPrefix = "Hola, "
const FrenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string{
	if name == ""{
		name = "world"
	}

	if language == spanish {
		return SpanishHelloPrefix + name
	}
	
	if language == french{
		return FrenchHelloPrefix + name
	}
	
	return englishHelloPrefix + name
}

func main(){
	fmt.Println(Hello("world", ""))
}