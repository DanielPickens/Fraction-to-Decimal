



func fractionToDecimal(numerator int, denominator int) string {
 res := ""
	if numerator == 0 {
		return "0"
	}
	if numerator < 0 && denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	}
	if numerator < 0 {
		res += "-"
		numerator = -numerator
	}
	if denominator < 0 {
		res += "-"
		denominator = -denominator
	}
	res += strconv.Itoa(numerator / denominator)
	numerator = numerator % denominator
	if numerator == 0 {
		return res
	}
	res += "."
	m := map[int]int{}
	for numerator != 0 {
		if i, ok := m[numerator]; ok {
			res = res[:i] + "(" + res[i:] + ")"
			return res
		}
		m[numerator] = len(res)
		numerator *= 10
		res += strconv.Itoa(numerator / denominator)
		numerator = numerator % denominator
	}
	return res
}


/*
Success
Details 
Runtime: 0 ms, faster than 100.00% of Go online submissions for Fraction to Recurring Decimal.
Memory Usage: 2.9 MB, less than 44.44% of Go online submissions for Fraction to Recurring Decimal.
*/
