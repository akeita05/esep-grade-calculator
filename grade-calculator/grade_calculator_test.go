package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

    gradeCalculator := NewGradeCalculator()

    gradeCalculator.AddGrade("assignment", 70, Assignment)
    gradeCalculator.AddGrade("exam", 70, Exam)
    gradeCalculator.AddGrade("essay", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

    if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
    }
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

    gradeCalculator := NewGradeCalculator()

    gradeCalculator.AddGrade("assignment", 60, Assignment)
    gradeCalculator.AddGrade("exam", 60, Exam)
    gradeCalculator.AddGrade("essay", 60, Essay)

	actual_value := gradeCalculator.GetFinalGrade()	

    if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
    }
}

func TestRealGetGradeF(t *testing.T) {
	expected_value := "F"

    gradeCalculator := NewGradeCalculator()

    gradeCalculator.AddGrade("assignment", 50, Assignment)
    gradeCalculator.AddGrade("exam", 50, Exam)
    gradeCalculator.AddGrade("essay", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()	
    if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
    } 
}

func TestGradeTypeString(t *testing.T) {
    tests := []struct {
        gradeType GradeType
        expected  string
    }{
        {Assignment, "assignment"},
        {Exam, "exam"},
        {Essay, "essay"},
    }
	
    for _, test := range tests {
        if actual := test.gradeType.String(); actual != test.expected {
            t.Errorf("Expected %s; got %s", test.expected, actual)
        }
    }

}

//Part 7

func TestPassFailGradePass(t *testing.T) {
	expected := "Pass"

	gradeCalculator := NewGradeCalculator(true)

	gradeCalculator.AddGrade("assignment", 90, Assignment)
	gradeCalculator.AddGrade("exam", 90, Exam)
	gradeCalculator.AddGrade("essay", 90, Essay)

	actual := gradeCalculator.GetFinalGrade()

	if actual != expected {
		t.Errorf("Expected '%s'; got '%s' instead", expected, actual)
	}
}

func TestPassFailGradeAtThreshold(t *testing.T) {
	expected := "Pass"

	gradeCalculator := NewGradeCalculator(true)

	gradeCalculator.AddGrade("assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam", 70, Exam)
	gradeCalculator.AddGrade("essay", 70, Essay)

	actual := gradeCalculator.GetFinalGrade()

	if actual != expected {
		t.Errorf("Expected '%s'; got '%s' instead", expected, actual)
	}
}

func TestPassFailGradeBelowThreshold(t *testing.T) {
	expected := "Fail"

	gradeCalculator := NewGradeCalculator(true)

	gradeCalculator.AddGrade("assignment", 69, Assignment)
	gradeCalculator.AddGrade("exam", 69, Exam)
	gradeCalculator.AddGrade("essay", 69, Essay)

	actual := gradeCalculator.GetFinalGrade()

	if actual != expected {
		t.Errorf("Expected '%s'; got '%s' instead", expected, actual)
	}
}

func TestPassFailGradeFail(t *testing.T) {
	expected := "Fail"

	gradeCalculator := NewGradeCalculator(true)

	gradeCalculator.AddGrade("assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam", 50, Exam)
	gradeCalculator.AddGrade("essay", 50, Essay)

	actual := gradeCalculator.GetFinalGrade()

	if actual != expected {
		t.Errorf("Expected '%s'; got '%s' instead", expected, actual)
	}
}
