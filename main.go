package main

import (
	"fmt"
)

func main() {
	f := FunctionImpl{
		domain: []int{2},
		funcValues: func(x []float64) []float64 {
			return []float64{
				x[0] * x[0], // 예: f(x) = x^2
			}
		},
	}

	// 점 정의
	point := Point{
		Coordinates: []float64{2.0},
	}

	// 미분계수 계산
	diffMatrix := Differential(f, point)

	// 결과 출력
	fmt.Println("Differential Matrix:")
	for _, row := range diffMatrix {
		fmt.Println(row)
	}
}
