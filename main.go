package main

import (
	"fmt"
	"os"
	"reflect"
)

type aoc struct{}

func main() {
	var dayNo string
	var runTest bool
	switch len(os.Args) {
	case 2:
		dayNo = os.Args[1]
		runTest = false
	case 3:
		dayNo = os.Args[1]
		if os.Args[2] == "test" {
			dayNo = os.Args[1]
			runTest = true
		} else {
			fmt.Printf("arg2 is either 'test' or empty, not: '%v'\n", os.Args[2])
			return
		}

	}

	dp := "/input/" + dayNo + ".txt"
	dtp := "/input/" + dayNo + ".test.txt"
	wd, _ := os.Getwd()
	inputPath := wd + dp
	testPath := wd + dtp

	a := aoc{}
	ao := reflect.TypeOf(&a)
	methodMap := map[string]string{}
	for i := 0; i < ao.NumMethod(); i++ {
		m := ao.Method(i)
		methodMap[m.Name] = m.Name
	}

	var data string
	if runTest {
		data = testPath
	} else {
		data = inputPath
	}

	input := []reflect.Value{reflect.ValueOf(data)}

	funcName := "Day" + dayNo
	funcName, exists := methodMap[funcName]

	if !exists {
		fmt.Println("That day is not created yet")
		return
	}

	dayfunc := reflect.ValueOf(a).MethodByName(funcName).Call(input)
	fmt.Printf("result %v: %v \n", funcName, dayfunc)
}
