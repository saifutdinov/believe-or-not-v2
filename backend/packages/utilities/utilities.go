package utilities

func Switch[valueT any](condition bool, valueIfTrue, valueIfFalse valueT) valueT {
	if condition {
		return valueIfTrue
	}
	return valueIfFalse
}
