package triangle

import (
	"encoding/json"
	"fmt"
	"os"
)

func Execute() {
	// เปิดไฟล์
	file, err := os.Open("triangle/hard.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode JSON
	var data [][]int
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	sum := maxPathSum(data)
	fmt.Println("Sum:", sum)
}

func maxPathSum(triangle [][]int) int {
	for row := len(triangle) - 2; row >= 0; row-- {
		rows := len(triangle[row])
		for col := 0; col < rows; col++ {
			triangle[row][col] += max(triangle[row+1][col], triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}
