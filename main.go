package luhn

func IsValidNumber(digits []int) bool {
	rd := make([]int, len(digits))
	copy(rd, digits)
	for i, j := 0, len(rd)-1; i < j; i, j = i+1, j-1 {
		rd[i], rd[j] = rd[j], rd[i]
	}

	var sum int
	alt := false
	for _, d := range rd {
		if alt {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		alt = !alt
	}
	return (sum % 10) == 0
}
