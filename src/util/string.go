package util

type String struct {
	str string
}

/*
func allToString[T int | byte | float64](val T) string {
	switch reflect.TypeOf(val).String() {
	case "int":
		return strconv.Itoa(int(val))
	case "byte":
		return string(byte(val))
	case "float64":
		return strconv.FormatFloat(float64(val), 'f', 10, 64)
	}
}
*/
