package service


type IGreeterService  interface{
	greet() string
}
type HelloWorldGreeter struct{
}

func (service HelloWorldGreeter) greet() (string){
	return "Hello world"
}