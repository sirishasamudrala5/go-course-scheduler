package main

type Course struct {
	CourseName        string
	Instructor        string
	Date              string
	MinEmployees      int
	MaxEmployees      int
	CourseOfferingId  string
	NoOfRegistrations int
	Allotted          bool
}

type Registration struct {
	EmployeeEmail        string
	CourseOfferingId     string
	CourseRegistrationId string
	Status               string
}
