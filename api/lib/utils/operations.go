package utils

func Ternary(comparison bool, isTrue interface{}, isFalse interface{}) (outcome interface{}) {
	if comparison {
		return isTrue
	}
	return isFalse
}
