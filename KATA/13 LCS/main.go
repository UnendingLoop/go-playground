package kata

func LCS(s1, s2 string) string {
	size1, size2 := len(s1), len(s2)
	matrix := make([][]int, size1+1)
	for i := range matrix {
		matrix[i] = make([]int, size2+1)
	}
	for i := 1; i <= size1; i++ {
		for j := 1; j <= size2; j++ {
			if s1[i-1] == s2[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = max(matrix[i][j-1], matrix[i-1][j])
			}
		}
	}
	answer := ""
	for i, j := size1, size2; i > 0 && j > 0; {
		switch {
		case s1[i-1] == s2[j-1]:
			answer = string(s1[i-1]) + answer
			i--
			j--
		case matrix[i-1][j] >= matrix[i][j-1]:
			i--
		default:
			j--
		}
	}
	return answer
}

func max(a, b int) int {
	switch {
	case a > b:
		return a
	case b > a:
		return b
	default:
		return a
	}

}

/*
func LCS(s1, s2 string) string {
	output := ""
	for _, i := range s2 {
		letter := string(i)
		if strings.Contains(s1, letter) {
			output += letter
		}
	}
	return output
}
*/
