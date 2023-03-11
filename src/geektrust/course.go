package main

import (
	"strings"
)

func addCourse(argList []string) string {
	if len(argList) != CommandArgCount["add_course"] { // validate no.of args
		return Response["invalidArgs"]
	}

	// validate unique instructor, course name
	for _, item := range CoursesList {
		if item.CourseName == argList[0] && item.Instructor == argList[1] {
			return Response["duplicateErr"]
		}
	}

	res := createCourse(argList)
	return res
}

func registerForCourse(argList []string) string {
	if len(argList) != CommandArgCount["add_registration"] { // validate no.of args
		return Response["invalidArgs"]
	}

	course, courseIndex, err := getCoursebyId(argList[1])
	if err != nil {
		return Response["courseErr"]
	}

	reg, _, err := getRegistrationbyCourseId(argList[1])
	if err == nil && reg.EmployeeEmail == argList[0] {
		return Response["duplicateErr"] // if registration for course exists for the same email id
	}

	if course.NoOfRegistrations == course.MaxEmployees {
		return Response["courseFullErr"]
	} else if !isValidDate(course.Date) {
		return Response["courseCancelledErr"]
	} else if course.NoOfRegistrations < course.MaxEmployees && !course.Allotted {
		regId, regStatus := createRegistration(argList)
		incrementRegistationsCount(courseIndex)

		return regId + " " + regStatus
	}
	return ""
}

func allotCourse(argList []string) string {
	if len(argList) != CommandArgCount["allot_course"] { // validate no.of args
		return Response["invalidArgs"]
	}
	course, courseIndex, err := getCoursebyId(argList[0])
	if err != nil {
		return Response["courseErr"]
	}

	if course.NoOfRegistrations == 0 { // zero registrations check
		return Response["zeroRegs"]
	}

	// result list
	var (
		resultList []string
		isAllotted bool
	)

	for i, ele := range RegistrationsList {
		if ele.CourseOfferingId == argList[0] {
			status := ele.Status
			if !isValidDate(course.Date) || course.NoOfRegistrations < course.MinEmployees {
				setRegistrationStatus(i, RegStatus["course_canceled"])
				status = RegStatus["course_canceled"]
			} else {
				setRegistrationStatus(i, RegStatus["confirmed"])
				isAllotted, status = true, RegStatus["confirmed"]
			}
			res := ele.CourseRegistrationId + " " + ele.EmployeeEmail + " " + ele.CourseOfferingId + " " + status
			resultList = append(resultList, res)
		}
	}
	CoursesList[courseIndex].Allotted = isAllotted

	return "\n" + strings.Join(resultList, "\n")
}

func cancelCourseRegistration(argList []string) string {
	if len(argList) != CommandArgCount["cancel_registration"] { // validate no.of args
		return Response["invalidArgs"]
	}

	reg, regIndex, err := getRegistrationbyId(argList[0])
	if err != nil {
		return Response["registrationErr"]
	}

	course, courseIndex, err := getCoursebyId(reg.CourseOfferingId)
	if err != nil {
		return Response["courseErr"]
	}

	if !course.Allotted && isValidDate(course.Date) {
		setRegistrationStatus(regIndex, RegStatus["cancel_accepted"])
		decrementtRegistationsCount(courseIndex)
	}
	if course.Allotted && course.CourseOfferingId == reg.CourseOfferingId {
		setRegistrationStatus(regIndex, RegStatus["cancel_rejected"])
	}
	return RegistrationsList[regIndex].CourseRegistrationId + " " + RegistrationsList[regIndex].Status
}
