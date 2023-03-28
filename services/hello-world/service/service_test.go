package service
import ("testing"
"github.com/stretchr/testify/assert"
)

func TestReturnsHelloWorld(t *testing.T){
	var expectedResult string = "Hello world"
	var  greeterService IGreeterService = HelloWorldGreeter{}

	var  actualResult string = greeterService.greet()
	assert.Equal(t, expectedResult, actualResult)
}

func TestDoesNotReturnsHelloMe(t *testing.T){
	var expectedResult string = "Hello me"
	var  greeterService IGreeterService = HelloWorldGreeter{}

	var  actualResult string = greeterService.greet()
	assert.NotEqual(t, expectedResult, actualResult)
}