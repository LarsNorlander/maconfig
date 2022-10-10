package parser

import (
	"strconv"
	"strings"
)

func ParseNumberList(str string) ([]int, error) {
	elems := strings.Split(str, ",")
	var result []int
	for _, elem := range elems {
		reps := 1
		if isRepetition(elem) {
			repDetails := strings.SplitN(elem, "*", 2)
			elem = repDetails[0] // remove the repetition info

			var err error // separated in order to prevent shadowing reps
			reps, err = strconv.Atoi(repDetails[1])
			if err != nil {
				return nil, err
			}
		}

		for i := 0; i < reps; i++ {
			if isNumberRange(elem) {
				nums, err := expandNumberRange(elem)
				if err != nil {
					return nil, err
				}
				result = append(result, nums...)
			} else {
				num, err := strconv.Atoi(elem)
				if err != nil {
					return nil, err
				}
				result = append(result, num)
			}
		}
	}

	return result, nil
}

func isNumberRange(str string) bool {
	return strings.Contains(str, "-")
}

func isRepetition(str string) bool {
	return strings.Contains(str, "*")
}

func expandNumberRange(str string) ([]int, error) {
	elems := strings.SplitN(str, "-", 2)
	left, err := strconv.Atoi(elems[0])
	if err != nil {
		return nil, err
	}
	right, err := strconv.Atoi(elems[1])
	if err != nil {
		return nil, err
	}

	var result []int
	// Not much of a range really, but I guess it could be legal
	if left == right {
		return append(result, left), nil
	} else if left > right {
		for i := left; i >= right; i-- {
			result = append(result, i)
		}
	} else {
		for i := left; i <= right; i++ {
			result = append(result, i)
		}
	}

	return result, nil
}
