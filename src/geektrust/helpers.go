package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func isValidDate(dateString string) bool {
	cDate, _ := time.Parse(YYYYMMDD, dateString)
	return today.Unix() < cDate.Unix()
}

func parseDate(dateString string) string {
	dateSlice := strings.Split(dateString, "")
	date := strings.Join(dateSlice[0:2], "")
	month := strings.Join(dateSlice[2:4], "")
	year := strings.Join(dateSlice[4:], "")
	cDate := strings.Join([]string{year, month, date}, "-")
	return cDate
}

func createCourse(data []string) string {
	courseId := "OFFERING-" + data[0] + "-" + data[1]
	minEmps, _ := strconv.Atoi(data[3])
	maxEmps, _ := strconv.Atoi(data[4])
	cDate := parseDate(data[2])

	// create new course
	course := Course{CourseName: data[0], Instructor: data[1], Date: cDate, MinEmployees: minEmps, MaxEmployees: maxEmps, CourseOfferingId: courseId, Allotted: false}
	CoursesList = append(CoursesList, course)
	return course.CourseOfferingId
}

func createRegistration(data []string) (string, string) {
	empName := strings.Split(data[0], "@")
	regId := "REG-COURSE-" + empName[0] + "-" + data[1]

	// create registration
	register := Registration{EmployeeEmail: data[0], CourseOfferingId: data[1], CourseRegistrationId: regId, Status: "ACCEPTED"}
	RegistrationsList = append(RegistrationsList, register)

	return register.CourseRegistrationId, register.Status
}

func setRegistrationStatus(key int, status string) {
	RegistrationsList[key].Status = status
}

func incrementRegistationsCount(key int) {
	CoursesList[key].NoOfRegistrations = CoursesList[key].NoOfRegistrations + 1 // increment by 1
}

func decrementtRegistationsCount(key int) {
	CoursesList[key].NoOfRegistrations = CoursesList[key].NoOfRegistrations - 1 // decrement by 1
}

func getCoursebyId(cId string) (Course, int, error) {
	for i, course := range CoursesList {
		if course.CourseOfferingId == cId {
			return course, i, nil
		}
	}
	res := Course{}
	return res, 0, errors.New("course not found")
}

func getRegistrationbyId(rId string) (Registration, int, error) {
	for i, reg := range RegistrationsList {
		if reg.CourseRegistrationId == rId {
			return reg, i, nil
		}
	}
	res := Registration{}
	return res, 0, errors.New("registration not found")
}

func getRegistrationbyCourseId(cId string) (Registration, int, error) {
	for i, reg := range RegistrationsList {
		if reg.CourseOfferingId == cId {
			return reg, i, nil
		}
	}
	res := Registration{}
	return res, 0, errors.New("registration not found")
}
