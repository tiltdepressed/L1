package main

import (
	"fmt"
	"reflect"
)

func getType(variable interface{}) string {
	switch variable.(type) {
	case int:
		return "INT"
	case string:
		return "STRING"
	case bool:
		return "BOOL"
	default:
		t := reflect.TypeOf(variable)
		if t.Kind() == reflect.Chan {
			return "CHAN"
		}
	}
	return "undefined"
}

func main() {
	var testInt int
	var testString string
	var testBool bool
	var testChanInt chan int
	var testChanStr chan string
	var testChanBool chan bool
	var testChanFloat chan float64
	var testNil interface{} = nil

	fmt.Println(getType(testInt))
	fmt.Println(getType(testString))
	fmt.Println(getType(testBool))
	fmt.Println(getType(testChanInt))
	fmt.Println(getType(testChanStr))
	fmt.Println(getType(testChanBool))
	fmt.Println(getType(testChanFloat))
	fmt.Println(getType(testNil))
}
