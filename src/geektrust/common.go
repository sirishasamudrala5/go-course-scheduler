package main

import "time"

// persist data
var CoursesList = []Course{}
var RegistrationsList = []Registration{}

const (
	YYYYMMDD = "2006-01-02"
)

var today = time.Now()

var RegStatus = map[string]string{
	"accepted":        "ACCEPTED",
	"cancel_accepted": "CANCEL_ACCEPTED",
	"cancel_rejected": "CANCEL_REJECTED",
	"course_canceled": "COURSE_CANCELED",
	"confirmed":       "CONFIRMED",
}

var Response = map[string]string{
	"dataErr":            "INPUT_DATA_ERROR",
	"courseErr":          "INPUT_DATA_ERROR - course not found",
	"registrationErr":    "INPUT_DATA_ERROR - registration not found",
	"duplicateErr":       "INPUT_DATA_ERROR - duplicate record",
	"invalidArgs":        "INPUT_DATA_ERROR - insufficient/invalid inputs",
	"courseFullErr":      "COURSE_FULL_ERROR",
	"courseCancelledErr": "COURSE_CANCELED",
	"zeroRegs":           "COURSE_CANCELED - zero registrations",
}

var Commands = map[string]string{
	"add_course":          "ADD-COURSE-OFFERING",
	"add_registration":    "REGISTER",
	"cancel_registration": "CANCEL",
	"allot_course":        "ALLOT-COURSE",
}

var CommandArgCount = map[string]int{
	"add_course":          5,
	"add_registration":    2,
	"cancel_registration": 1,
	"allot_course":        1,
}
