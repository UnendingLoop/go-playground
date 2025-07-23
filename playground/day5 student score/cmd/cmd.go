package school

import (
	"fmt"
	"school/logger"
	"slices"
	"sort"
)

// SAnalyzer - interface for methods of type "school", which is a map - key is a student, value is a map of subjects and student's grades
type SAnalyzer interface {
	AvgMarkStudent()
	OverallAvgMark() float64
	BestStudent()
	BestSubject()
	UnevenStudent()
	UnevenSubject()
	StudentInfo(name string)
}

// School - structure for storing grades of students
type School struct {
	Grades map[string]map[string]int
	Logger logger.Logger
}

// AvgMarkStudent - returns avg mark of a student for all subjects
func (journal School) AvgMarkStudent() {
	numberOfSubjects := 0
	for personName, subjects := range journal.Grades {
		sum := 0
		avg := 0.0
		numberOfSubjects = len(subjects)
		for _, grade := range subjects {
			sum += grade
		}
		avg = float64(sum) / float64(numberOfSubjects)
		fmt.Printf("Student %s has an avg mark of %.2f.\n", personName, avg)
	}
}

// OverallAvgMark - index that shows an avg mark among the whole school
func (journal School) OverallAvgMark() float64 {
	counter := 0
	sum := 0
	for _, person := range journal.Grades {
		for _, subject := range person {
			sum += subject
			counter++
		}
	}
	result := float64(sum) / float64(counter)
	return result
}

// BestStudent - finds and returns the most successful student
func (journal School) BestStudent() {
	type student struct {
		name    string
		avgMark float64
	}
	list := make([]student, 0)
	maxAvgMark := 0.0

	for name, subjects := range journal.Grades {
		sum := 0
		for _, marks := range subjects {
			sum += marks
		}
		avg := float64(sum) / float64(len(subjects))

		switch {
		case maxAvgMark < avg:
			maxAvgMark = avg
			list = list[:0]
			list = append(list, student{name: name, avgMark: avg})
		case maxAvgMark == avg:
			list = append(list, student{name: name, avgMark: avg})
		}

	}

	for _, val := range list {
		fmt.Printf("The best student is %s, avg mark is %.2f.\n", val.name, val.avgMark)
	}

}

// BestSubject - finds the subject best studied among students
func (journal School) BestSubject() {
	type rating struct {
		sumMarks int
		count    int
	}

	list := make(map[string]rating)
	ratingResult := make([]struct {
		subjectname string
		avgMark     float64
	}, 0)
	maxAvg := 0.0

	for _, subjects := range journal.Grades { //достаем из исходной карты значения и распределяем в новую карту, где ключи - предметы
		for subject, mark := range subjects {
			candidate := list[subject]
			candidate.sumMarks += mark
			candidate.count++
			list[subject] = candidate
		}
	}

	for name, marks := range list { //вычисляем среднее и помещаем в результирующий слайс только предмет с максимальным средним значением
		avg := (float64(marks.sumMarks) / float64(marks.count))
		switch {
		case avg > maxAvg:
			ratingResult = append(ratingResult[:0], struct {
				subjectname string
				avgMark     float64
			}{subjectname: name, avgMark: avg})
		case avg == maxAvg:
			ratingResult = append(ratingResult, struct {
				subjectname string
				avgMark     float64
			}{subjectname: name, avgMark: avg})

		}

	}

	for _, i := range ratingResult {
		journal.Logger.Printf("The most successfully learnt subject is %s, its avg score is %.2f.\n", i.subjectname, i.avgMark)
	}
}

// UnevenStudent - shows a student with the biggest difference in their grades
func (journal School) UnevenStudent() {
	type student struct {
		name  string
		marks []int
		diff  int
	}
	unevenStudents := make([]student, 0)
	maxDiff := 0
	for name, subjects := range journal.Grades {
		var candidate student
		for _, mark := range subjects {
			candidate.name = name
			candidate.marks = append(candidate.marks, mark)
		}
		sort.Slice(candidate.marks, func(i, j int) bool {
			return candidate.marks[i] < candidate.marks[j]
		})
		candidate.diff = candidate.marks[len(candidate.marks)-1] - candidate.marks[0]
		switch {
		case candidate.diff > maxDiff:
			unevenStudents = append(unevenStudents[:0], candidate)
			maxDiff = candidate.diff
		case candidate.diff == maxDiff:
			unevenStudents = append(unevenStudents, candidate)

		}
	}
	for _, value := range unevenStudents {
		fmt.Printf("Student with most uneven marks is %s, their marks are %v.\n", value.name, value.marks)
	}
}

// UnevenSubject - shows a subject with most differentiated grades among students
func (journal School) UnevenSubject() {
	type subject struct {
		name       string
		marksArray []int
	}
	list := make(map[string]subject)

	for _, subjects := range journal.Grades { //делаем карту предмет-оценки - достаем оценки из каждого студента
		for name, mark := range subjects {
			subject := list[name]
			subject.name = name
			subject.marksArray = append(subject.marksArray, mark)
			list[name] = subject
		}
	}
	maxDiff := 0
	result := []subject{}
	for _, marks := range list {
		slices.Sort(marks.marksArray)
		diff := marks.marksArray[len(marks.marksArray)-1] - marks.marksArray[0] //вычисляем и добавляем недостающую инфу в карту

		switch {
		case maxDiff < diff:
			maxDiff = diff
			result = append(result[:0], marks)
		case maxDiff == diff:
			result = append(result, marks)
		}
	}

	for _, value := range result {
		fmt.Printf("The subject with most differentiated marks among students is %s, the marks are %v.\n", value.name, value.marksArray)
	}
}

// StudentInfo - returns info about student using their name as string
func (journal School) StudentInfo(name string) {
	// Показывает оценки конкретного ученика, среднюю оценку и предмет с самой высокой оценкой
	type subject struct {
		name string
		mark int
	}
	subjects := []subject{}
	candidate, exists := journal.Grades[name]
	if !exists {
		journal.Logger.Printf("Student with name '%s' is not found!\n", name)
		return
	}
	markSum := 0
	marks := []int{}
	for item, mark := range candidate {
		subjects = append(subjects, subject{name: item, mark: mark})
		markSum += mark
		marks = append(marks, mark)
	}
	avgMark := float64(markSum) / float64(len(candidate))
	sort.Slice(subjects, func(i, j int) bool {
		return subjects[i].mark < subjects[j].mark
	})
	journal.Logger.Printf("Information about student: %s\n", name)
	journal.Logger.Printf("Their marks: %v\n", marks)
	journal.Logger.Printf("Their avg score: %.2f\n", avgMark)
	journal.Logger.Printf("Most successful in subject: %s\n", subjects[len(subjects)-1].name)
}
