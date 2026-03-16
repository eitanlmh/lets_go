package main

import (
	"fmt"
)

func addGrade(students map[string][]int, name string, grade int) {
	students[name] = append(students[name], grade)
}

func studentAverage(students map[string][]int, name string) float64 {
	grade_sum := 0.0
	for _, grade := range students[name] {
		grade_sum += float64(grade)
	}
	return grade_sum / float64(len(students[name]))
}

func overallAverage(students map[string][]int) float64 {
	average_sum := 0.0
	for key, _ := range students {
		average_sum += studentAverage(students, key)
	}
	return average_sum / float64(len(students))
}

func main() {

	students := map[string][]int{
		"eitan": {},
		"saar":  {},
	}
	choice := 0
outer:
	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Add a grade to a student")
		fmt.Println("2. Print grade avg of a student")
		fmt.Println("3. Print overall avg grade of all students")
		fmt.Println("4. Exit the program")
		student_name := ""
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("enter a student name")
			fmt.Scan(&student_name)
			grade := 0
			fmt.Println("enter grade")
			fmt.Scan(&grade)
			addGrade(students, student_name, grade)
		case 2:
			fmt.Println("enter a student name:")
			fmt.Scan(&student_name)
			avg := studentAverage(students, student_name)
			fmt.Println(student_name, "s avergae", avg)
		case 3:
			all_student_avg := overallAverage(students)
			fmt.Println("overall avergae", all_student_avg)
		case 4:
			break outer
		}
	}

}
