package service
import ("testing"
"github.com/stretchr/testify/assert"
)

func TestReturnsHelloWorld(t *testing.T){
	var expectedResult string = "Hello world"
	var  greeterService IGreeterService = HelloWorldGreeter{}

	var  actualResult string = greeterService.Greet()
	assert.Equal(t, expectedResult, actualResult)
}

func TestDoesNotReturnsHelloMe(t *testing.T){
	var expectedResult string = "Hello me"
	var  greeterService IGreeterService = HelloWorldGreeter{}

	var  actualResult string = greeterService.Greet()
	assert.NotEqual(t, expectedResult, actualResult)
}


func TestReturnsHelloDevOps(t *testing.T){
	var expectedResult string = "Hello DevOps"
	var  greeterService IGreeterService = HelloDevOpsGreeter{}

	var  actualResult string = greeterService.Greet()
	assert.Equal(t, expectedResult, actualResult)
}

func TestDoesNotReturnsHelloDevOps(t *testing.T){
	var expectedResult string = "Hello me"
	var  greeterService IGreeterService = HelloDevOpsGreeter{}

	var  actualResult string = greeterService.Greet()
	assert.NotEqual(t, expectedResult, actualResult)
}