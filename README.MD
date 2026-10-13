# ESEP Grade Calculator

## Overview

The ESEP Grade Calculator is a Go library that calculates a student's final numerical grade and converts it into a letter grade. It supports three categories of grades: assignments, exams, and essays.

Each category contributes a weighted percentage to the final grade:

- Assignments: 50%
- Exams: 35%
- Essays: 15%

The final numerical grade is calculated by finding the average of each category and multiplying it by that category's weight. The weighted values are then added together.

The letter grade is determined using the following scale:

| Letter Grade | Numerical Grade |
|---|---|
| A | 90–100 |
| B | 80–89 |
| C | 70–79 |
| D | 60–69 |
| F | Below 60 |

A grade of C or better (70 or higher) is considered passing.

## Requirements

- Go 1.24 or newer
- Git, if cloning the repository

## Getting Started

Clone the repository and navigate to the Go module directory:

```bash
git clone https://github.com/audgster/esep-grade-calculator.git
cd esep-grade-calculator/grade-calculator
```

## Running Unit Tests

Run the existing unit tests with:

```bash
go test ./...
```

To see the individual tests and their results, use:

```bash
go test -v ./...
```

## Measuring Code Coverage

Generate a code coverage report with:

```bash
go test -coverprofile=coverage.txt ./...
```

Display coverage statistics for individual functions:

```bash
go tool cover -func=coverage.txt
```

Generate an HTML coverage report:

```bash
go tool cover -html=coverage.txt
```

## Project Structure

- `README.MD`: Project overview, setup instructions, and testing documentation.
- `go.mod`: Defines the Go module and its Go version requirement.
- `grade_calculator.go`: Contains the grade calculator implementation.
- `grade_calculator_test.go`: Contains the unit tests for the grade calculator.

## Development Notes

Run the unit tests before and after making code changes to verify the calculator's behavior. When modifying the implementation, maintain or expand the test suite to check correctness and code coverage.