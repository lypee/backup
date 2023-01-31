package multiply

func multiply(num1 string, num2 string) string {
	num1 = reverseStr(num1)
	num2 = reverseStr(num2)
	res := ""
	for idx , c := range num1{

	}
}



func reverseStr(str string) string {
	newStr := []byte{}
	for i := 0; i < len(str); i++ {
		newStr = append(newStr, str[i])
	}
	return string(newStr)
}

