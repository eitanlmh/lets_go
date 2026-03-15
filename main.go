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
	for key := range students {
		for range 2 {
			fmt.Println("enter a new grade for:", key)
			grade := 0
			fmt.Scan(&grade)
			fmt.Println("entered grade:", grade)
			addGrade(students, key, grade)
		}
		fmt.Println("all of", key, "'s grades are:", students[key])
	}

	name := ""
	fmt.Println("enter a student name:")
	fmt.Scan(&name)
	avg := studentAverage(students, name)
	fmt.Println(name, "s avergae", avg)

	all_student_avg := overallAverage(students)
	fmt.Println("overall avergae", all_student_avg)

}
