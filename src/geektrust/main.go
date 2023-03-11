package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mapCommands(argList []string) {
	res := ""
	switch argList[0] {
	case Commands["add_course"]:
		res = addCourse(argList[1:])
	case Commands["add_registration"]:
		res = registerForCourse(argList[1:])
	case Commands["allot_course"]:
		res = allotCourse(argList[1:])
	case Commands["cancel_registration"]:
		res = cancelCourseRegistration(argList[1:])
	default:
		res = Response["dataErr"]
	}
	fmt.Println(argList[0], " - ", strings.Join(argList[1:], ", "), " : ", res)
}

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 { // zero arguments check
		fmt.Println("Please provide the input file path")
		return
	}

	// filePath := cliArgs[0]
	file, _ := os.Open(cliArgs[0])

	// if err != nil {
	// 	fmt.Println("Error opening the input file")
	// 	return
	// }

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		argList := strings.Fields(scanner.Text())
		mapCommands(argList) // map arguments to func
	}
}
