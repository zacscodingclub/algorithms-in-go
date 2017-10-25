package testutil

func Expect(a, b interface{}) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func ExpectSlice(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
