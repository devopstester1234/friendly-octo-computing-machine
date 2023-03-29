package main

import (
	"github.com/devopstester1234/friendly-octo-computing-machine/services/hello-world/service"
	"fmt"
)

func main(){
	var  greeterService service.IGreeterService = service.HelloDevOpsGreeter{}
	
	var greeting = greeterService.Greet()
	fmt.Printf("%s\n", greeting)	

	greeterService = service.HelloWorldGreeter{}
	greeting = greeterService.Greet()
	fmt.Printf("%s\n", greeting)
	// greeterService = service.HelloWorldGreeter{}
	
}