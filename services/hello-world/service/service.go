package service

type IGreeterService  interface{
	Greet() string
}
type HelloWorldGreeter struct{
}
type HelloDevOpsGreeter struct{
}

func (greeter HelloWorldGreeter) Greet() (string){
	return "Hello world"
}

func (greeter HelloDevOpsGreeter) Greet() (string){
	return "Hello DevOps"
}