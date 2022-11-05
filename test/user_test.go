package test

import (
	"fmt"
	"testing"
)

func TestSomethingWithPlaceholder(t *testing.T) {

	testObj := new(MyMockObject)

	testObj.On("DoSomething", 1).Return(false, nil)

	// call the code we are testing
	targetFuncThatDoesSomethingWithObj(testObj)

	// assert that the expectations were met
	testObj.AssertExpectations(t)

}

func targetFuncThatDoesSomethingWithObj(obj *MyMockObject) {
	// do something with obj
	fmt.Println(obj.DoSomething(1))
}
