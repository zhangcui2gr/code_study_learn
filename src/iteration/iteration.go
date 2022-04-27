package iteration

func Repeat(s string, num int) string {
	var str string
	for i := 0; i < num; i++ {
		str += s
	}
	return str

}
