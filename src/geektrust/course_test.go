package main

import "testing"

func TestAddCourse(t *testing.T) {
	want := "OFFERING-GoLang-John"
	if got := addCourse([]string{"GoLang", "John", "22022023", "2", "2"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAddDuplicateCourse(t *testing.T) {
	want := "INPUT_DATA_ERROR - duplicate record"
	if got := addCourse([]string{"GoLang", "John", "22022023", "2", "2"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAddCourseWithInvalidData(t *testing.T) {
	want := "INPUT_DATA_ERROR - insufficient/invalid inputs"
	if got := addCourse([]string{"GoLang", "John", "22022023"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAddRegistrationWithoutCourse(t *testing.T) {
	want := "INPUT_DATA_ERROR - course not found"
	if got := registerForCourse([]string{"jkl@mail.com", "OFFERING-Python-Lara"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAddRegistrationWithCourse(t *testing.T) {
	addCourse([]string{"Python", "Lara", "24022023", "1", "5"})
	want := "REG-COURSE-jkl-OFFERING-Python-Lara ACCEPTED"
	if got := registerForCourse([]string{"jkl@mail.com", "OFFERING-Python-Lara"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAddDuplicateRegistration(t *testing.T) {
	want := "INPUT_DATA_ERROR - duplicate record"
	if got := registerForCourse([]string{"jkl@mail.com", "OFFERING-Python-Lara"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAddRegistrationWithCourseWithInvalidData(t *testing.T) {
	addCourse([]string{"Python", "Lara", "24022023", "1", "1"})
	want := "INPUT_DATA_ERROR - insufficient/invalid inputs"
	if got := registerForCourse([]string{"jkl@mail.com"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAddRegistrationOverflow(t *testing.T) {
	addCourse([]string{"Javascript", "Jacob", "23022023", "1", "1"})
	registerForCourse([]string{"jkl@mail.com", "OFFERING-Javascript-Jacob"})
	want := "COURSE_FULL_ERROR"
	if got := registerForCourse([]string{"opr@mail.com", "OFFERING-Javascript-Jacob"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAddRegistrationExpiredCourse(t *testing.T) {
	addCourse([]string{"Physics", "Payal", "18022023", "2", "4"})
	want := "COURSE_CANCELED"
	if got := registerForCourse([]string{"opr@mail.com", "OFFERING-Physics-Payal"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCancelUnAllotedRegistration(t *testing.T) {
	want := "REG-COURSE-jkl-OFFERING-Python-Lara CANCEL_ACCEPTED"
	if got := cancelCourseRegistration([]string{"REG-COURSE-jkl-OFFERING-Python-Lara"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCancelAllotedRegistration(t *testing.T) {
	addCourse([]string{"Javascript", "Jacob", "23022023", "1", "1"})
	registerForCourse([]string{"jkl@mail.com", "OFFERING-Javascript-Jacob"})
	allotCourse([]string{"OFFERING-Javascript-Jacob"})

	want := "REG-COURSE-jkl-OFFERING-Javascript-Jacob CANCEL_REJECTED"
	if got := cancelCourseRegistration([]string{"REG-COURSE-jkl-OFFERING-Javascript-Jacob"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCancelNonExistingRegistration(t *testing.T) {
	addCourse([]string{"Javascript", "Jacob", "23022023", "3", "7"})
	want := "INPUT_DATA_ERROR - registration not found"
	if got := cancelCourseRegistration([]string{"REG-COURSE-mno-OFFERING-Javascript-Jacob"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAllotCourseWithInvalidData(t *testing.T) {
	want := "INPUT_DATA_ERROR - insufficient/invalid inputs"
	if got := allotCourse([]string{}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAllotCourseWithNonExistingCourse(t *testing.T) {
	want := "INPUT_DATA_ERROR - course not found"
	if got := allotCourse([]string{"OFFERING-Python-John"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAllotCourse(t *testing.T) {
	addCourse([]string{"Java", "Elsa", "23022023", "1", "2"})
	registerForCourse([]string{"jkl@mail.com", "OFFERING-Java-Elsa"})

	want := "\nREG-COURSE-jkl-OFFERING-Java-Elsa jkl@mail.com OFFERING-Java-Elsa CONFIRMED"
	if got := allotCourse([]string{"OFFERING-Java-Elsa"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAllotCourseCancel(t *testing.T) {
	addCourse([]string{"Nodejs", "Nina", "23022023", "3", "7"})
	// registerForCourse([]string{"jkl@mail.com", "OFFERING-Nodejs-Nina"})
	want := "COURSE_CANCELED - zero registrations"
	if got := allotCourse([]string{"OFFERING-Nodejs-Nina"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAllotCourseExpiredCourse(t *testing.T) {
	addCourse([]string{"Chemistry", "Elsa", "17022023", "2", "2"})

	want := "COURSE_CANCELED - zero registrations" // cant add registration to test this
	if got := allotCourse([]string{"OFFERING-Chemistry-Elsa"}); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
