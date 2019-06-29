package number


// https://en.m.wikipedia.org/wiki/Binary_GCD_algorithm
func gcd(u int, v int) int {

	// gcd(0, v) = v
	if u == 0 {
		return v
	}
	if v == 0 {
		return u
	}

	// If u and v are both even, then gcd(u, v) = 2·gcd(u/2, v/2)
	if u % 2 == 0 && v % 2 == 0 {
		return 2 * gcd(u/ 2, v/ 2)
	}

	// If u is even and v is odd, then gcd(u, v) = gcd(u/2, v)
	if u% 2 == 0 && v % 2 != 0 {
		return gcd(u/ 2, v)
	}

	// if u is odd and v is even, then gcd(u, v) = gcd(u, v/2)
	if u % 2 != 0 && v % 2 == 0 {
		return gcd(u, v/ 2)
	}

	// If u and v are both odd, and u ≥ v, then gcd(u, v) = gcd((u − v)/2, v)
	if u < v {
		u,v = v, u
	}
	return gcd((u - v) / 2, v)
}