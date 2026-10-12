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

func TestGetGradeHighScores(t *testing.T) {
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

func TestGetFinalGradeBoundaries(t *testing.T) {
	tests := []struct {
		name     string
		score    int
		expected string
	}{
		{"A at 90", 90, "A"},
		{"B at 89", 89, "B"},
		{"B at 80", 80, "B"},
		{"C at 79", 79, "C"},
		{"C at 70", 70, "C"},
		{"D at 69", 69, "D"},
		{"D at 60", 60, "D"},
		{"F at 59", 59, "F"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			calculator := NewGradeCalculator()

			calculator.AddGrade("assignment", test.score, Assignment)
			calculator.AddGrade("exam", test.score, Exam)
			calculator.AddGrade("essay", test.score, Essay)

			actual := calculator.GetFinalGrade()

			if actual != test.expected {
				t.Errorf("Expected %s, got %s",
					test.expected, actual)
			}
		})
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
		actual := test.gradeType.String()

		if actual != test.expected {
			t.Errorf("Expected %s, got %s",
				test.expected, actual)
		}
	}
}

func TestPassFailMode(t *testing.T) {
	tests := []struct {
		name       string
		assignment int
		exam       int
		essay      int
		expected   string
	}{
		{"passing at 70", 70, 70, 70, "Pass"},
		{"failing at 69", 69, 69, 69, "Fail"},
		{"weighted passing grade", 80, 60, 70, "Pass"},
		{"weighted failing grade", 60, 70, 80, "Fail"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			calculator := NewGradeCalculator(PassFail)

			calculator.AddGrade("assignment", test.assignment, Assignment)
			calculator.AddGrade("exam", test.exam, Exam)
			calculator.AddGrade("essay", test.essay, Essay)

			actual := calculator.GetFinalGrade()

			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}
