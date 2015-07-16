package gojw

import (
	"fmt"
	"testing"
)

func TestJaroWinkler(t *testing.T) {
	jw1 := JaroWinkler("Hi I am a string", "Hi")
	fmt.Println("JW of 'Hi I am a string' and 'Hi' is: ", jw1)
	if jw1 != 0 {
		t.Fail()
	}
	jw2 := JaroWinkler("Hi", "Hi I am a string")
	if jw2 != 0 {
		t.Fail()
	}
	fmt.Println("JW of 'Hi I am a string' and 'Hi' is: ", jw2)
	jw3 := JaroWinkler("There is no spoon", "There is a hammer")
	fmt.Println("JW of 'There is no spoon', and 'There is a hammer' is: ", jw3)
	if jw3 != 0.5 {
		t.Fail()
	}
	jw4 := JaroWinkler("Yo yo yo wassup", "Allo mate")
	if jw4 != 1.0 {
		t.Fail()
	}
	fmt.Println("JW of 'Yo yo yo wassup' and 'Allo mate' is: ")
}
