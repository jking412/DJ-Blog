package test

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSomethingWithPlaceholder(t *testing.T) {

	testObj := new(MyMockObject)

	testObj.On("DoSomething", mock.Anything).Return(true, nil)

	// call the code we are testing
	targetFuncThatDoesSomethingWithObj(testObj)

	// assert that the expectations were met
	testObj.AssertExpectations(t)

}

func targetFuncThatDoesSomethingWithObj(obj *MyMockObject) {
	// do something with obj
	fmt.Println(obj.DoSomething(1))
}
