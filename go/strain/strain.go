package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(check func(int) bool) Ints {
	var result Ints

	for _, v := range i {
		if check(v) {
			result = append(result, v)
		}
	}

	return result
}

func (i Ints) Discard(check func(int) bool) Ints {
	var result Ints

	for _, v := range i {
		if !check(v) {
			result = append(result, v)
		}
	}

	return result
}

func (l Lists) Keep(check func([]int) bool) Lists {
	var result Lists

	for _, v := range l {
		if check(v) {
			result = append(result, v)
		}
	}

	return result
}

func (s Strings) Keep(check func(string) bool) Strings {
	var result Strings

	for _, v := range s {
		if check(v) {
			result = append(result, v)
		}
	}

	return result
}
