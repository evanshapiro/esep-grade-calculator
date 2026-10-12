
# ESEP Grade Calculator

A simple Go library for calculating students' final course grades based on assignments, exams, and essays.

## Requirements

- Go 1.24.4 or later
- Git (to clone the repository)

The project uses Go's built-in testing package and does not require additional third-party dependencies.

## Getting Started

Clone the repository:

```bash
git clone https://github.com/YOUR-USERNAME/esep-grade-calculator.git
cd esep-grade-calculator/grade-calculator
```

Replace `YOUR-USERNAME` with the owner of your fork.

## How It Works

The grade calculator supports three types of grades:

- Assignments: 50%
- Exams: 35%
- Essays: 15%

The intended final numerical grade is calculated by averaging the grades in each category and multiplying those averages by their respective weights.

Final Grade = (Assignment Average × 0.50) + (Exam Average × 0.35) + (Essay Average × 0.15)

The calculator converts the numerical result to a letter grade using the following scale:

| Numerical Grade | Letter Grade |
|---|---|
| 90–100 | A |
| 80–89 | B |
| 70–79 | C |
| 60–69 | D |
| Below 60 | F |

## Using the Calculator

The library provides the following methods:

- `NewGradeCalculator()` creates a grade calculator.
- `AddGrade(name, grade, gradeType)` adds a grade.
- `GetFinalGrade()` calculates and returns the final letter grade.

Example usage within the Go package:

```go
calculator := NewGradeCalculator()

calculator.AddGrade("Assignment 1", 95, Assignment)
calculator.AddGrade("Exam 1", 90, Exam)
calculator.AddGrade("Essay 1", 85, Essay)

finalGrade := calculator.GetFinalGrade()
```

The example shows the intended API usage. The starter implementation is still under development.

## Running Unit Tests

From the `grade-calculator` directory, run:

```bash
go test ./...
```

For detailed results showing individual tests:

```bash
go test -v ./...
```

The initial repository includes three unit tests:

- `TestGetGradeA`
- `TestGetGradeB`
- `TestGetGradeF`

Some starter tests fail because the implementation and test expectations require corrections.

## Measuring Code Coverage

Run the tests with coverage enabled:

```bash
go test -coverprofile=coverage.out ./...
```

Display coverage information by function:

```bash
go tool cover -func=coverage.out
```

Generate an HTML coverage report:

```bash
go tool cover -html=coverage.out
```

Code coverage measures which portions of the program are executed during testing. High coverage does not necessarily mean all possible behaviors have been tested correctly.

## Project Files

- `README.MD` — Project documentation.
- `go.mod` — Go module configuration.
- `grade_calculator.go` — Grade calculation implementation.
- `grade_calculator_test.go` — Unit tests.

## Additional Resources

- Go Documentation: https://go.dev/doc/
- Go Testing Tutorial: https://go.dev/doc/tutorial/add-a-test
- Go Testing Package: https://pkg.go.dev/testing
```
