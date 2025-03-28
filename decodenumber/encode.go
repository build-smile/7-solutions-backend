package decodenumber

import (
	"fmt"
	"slices"
	"strconv"
)

func Execute(encoded string) {
	result := decodeToInt(encoded)
	setValPosVal(result)
	strResult := arrToStr(result)
	fmt.Println(strResult)
}

func setValPosVal(result []int) {
	minVal := slices.Min(result)
	shift := -minVal
	for i := range result {
		result[i] += shift
	}
}

func decodeToInt(encoded string) []int {
	n := len(encoded) + 1
	result := make([]int, n)
	result[0] = 0
	for i := 0; i < len(encoded); i++ {
		nextInd := i + 1
		if encoded[i] == 'L' { // ซ้ายมากกว่าขวา
			result[nextInd] = result[i] - 1 // -1 เพื่อให้ขวาน้อยกว่าซ้าย
		} else if encoded[i] == 'R' { //ขวามากกว่าซ้าย
			result[nextInd] = result[i] + 1 // +1 เพื่อให้ขวามากกว่าซ้าย
		} else { // '=' case
			result[nextInd] = result[i]
		}
	}
	return result
}

func arrToStr(result []int) string {
	strResult := ""
	for _, v := range result {
		strResult += strconv.Itoa(v)
	}
	return strResult
}
