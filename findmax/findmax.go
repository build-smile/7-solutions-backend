package findmax

import (
	"encoding/json"
	"fmt"
	"os"
)

func Execute() {
	// เปิดไฟล์
	file, err := os.Open("findmax/t.json")
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
	fmt.Println("Max Path Sum:", sum)
}

// ฟังก์ชันคำนวณหาค่ามากสุดของเส้นทาง
func maxPathSum(triangle [][]int) int {
	// คัดลอกข้อมูลมาเก็บไว้เพื่อหลีกเลี่ยงการเปลี่ยนแปลงต้นฉบับ
	n := len(triangle)
	dp := make([][]int, n)
	for i := range triangle {
		dp[i] = make([]int, len(triangle[i]))
		copy(dp[i], triangle[i])
	}

	// ใช้วิธี Bottom-Up: เริ่มจากแถวรองสุดท้าย แล้วไล่ขึ้นไปจนถึงแถวบนสุด
	for row, it := range dp {
		for col := 0; col < len(it); col++ {
			// อัปเดตค่าของแต่ละตำแหน่ง โดยเลือกค่ามากสุดจากลูกทางซ้ายหรือขวา
			dp[row][col] += max(dp[row+1][col], dp[row+1][col+1])
		}
	}
	// ค่ามากสุดของเส้นทางจะอยู่ที่ dp[0][0]
	return dp[0][0]
}

// ฟังก์ชันหาค่าสูงสุดระหว่างสองค่า
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
