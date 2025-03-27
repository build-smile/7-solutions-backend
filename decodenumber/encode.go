package decodenumber

import (
	"fmt"
	"strconv"
)

func Execute(encode string) {
	fmt.Println(decodeMinSum(encode))
}

func decodeMinSum(encoded string) string {
	n := len(encoded) + 1
	result := make([]int, n)

	// ตั้งค่าตัวแรกเป็นค่าต่ำสุดที่เป็นไปได้
	result[0] = 0

	// คำนวณค่าที่น้อยที่สุดที่ตรงกับเงื่อนไขของแต่ละคู่
	for i := 0; i < len(encoded); i++ {
		if encoded[i] == 'L' {
			result[i+1] = result[i] - 1
		} else if encoded[i] == 'R' {
			result[i+1] = result[i] + 1
		} else { // '=' case
			result[i+1] = result[i]
		}
	}

	// Shift ค่าให้เป็นบวกทั้งหมด (ทำให้เลขน้อยที่สุดเป็น 0)
	minVal := 0
	for _, v := range result {
		if v < minVal {
			minVal = v
		}
	}
	shift := -minVal
	for i := range result {
		result[i] += shift
	}

	// แปลงผลลัพธ์เป็น string
	strResult := ""
	for _, v := range result {
		strResult += strconv.Itoa(v)
	}

	return strResult
}
