package school

import (
	"fmt"
	school "school/cmd"
)

func main() {
	grades := map[string]map[string]int{
		"Alice": {
			"Math":    0,
			"English": 5,
			"Deutsch": 4,
		},
		"Bob": {
			"Math":    3,
			"English": 5,
			"Deutsch": 4,
		},
		"Sam": {
			"Math":    5,
			"English": 0,
			"Deutsch": 4,
		},
		"Tery": {
			"Math":    4,
			"English": 4,
			"Deutsch": 4,
		},
		"Marta": {
			"Math":    2,
			"English": 1,
			"Deutsch": 5,
		},
	}
	var analyzer school.SAnalyzer = school.School{Grades: grades}

	schoolAvgMark := analyzer.OverallAvgMark()
	fmt.Printf("School avg mark is %.2f.\n", schoolAvgMark)
	analyzer.AvgMarkStudent()
	analyzer.BestStudent()
	analyzer.BestSubject()
	analyzer.UnevenStudent()
	analyzer.UnevenSubject()
	analyzer.StudentInfo("Sam")

}
