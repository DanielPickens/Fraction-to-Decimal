// First, we can check to see if the number is an integer (remainder of numerator/denominator is 0), and if it is,
//we can simply return the number.
//Otherwise we'll check to see if one of the numbers but not both numbers is a negative. 
//Next we add a - if this is the case.
// Then we can add our integer to the result because we know it cannot be repeating.
// Now, we can declare a new map for long remainder to it's position in result.
// From here, we can get the first remainder and add a 0 to it by multiplying it by 10.
// We know the next digit is the remainder/denominator. 
//If we have already seen the remainder we add the ( to the position we saw it, 
//and then add the ) at the end and return the value. Otherwise, we put the remainder at the current length and add our next digit, we can then get our next remainder.



//time complexity: O(len(frac))
//space complexity: O(n)

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
