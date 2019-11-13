package errors

import (
	"fmt"
	"testing"
)

func step1() error {
	return step2()
}
func step2() error {
	return step3()
}
func step3() error {
	return New("from 3")
}

func step1w() error {
	return step2w()
}
func step2w() error {
	return step3w()
}
func step3w() error {
	return WrapSkip("from 3", 1)
}

func TestFrame(t *testing.T) {
	err := step1()
	fmt.Printf("%+v", err)
	fmt.Printf("%+v", step1w())
}
