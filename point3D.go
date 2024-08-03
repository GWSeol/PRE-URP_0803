package main

import (
	"fmt"
	"math"
)

type Point3D struct {
	X, Y, Z float64
}

func Distance(p1, p2 Point3D) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dz := p2.Z - p1.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func main() {
	var x1, y1, z1, x2, y2, z2 float64

	fmt.Print("좌표입력 x1 y1 z1: ")
	_, err := fmt.Scanf("%f %f %f", &x1, &y1, &z1)
	if err != nil {
		fmt.Println("에러:", err)
		return
	}

	fmt.Print("좌표입력 x2 y2 z2 : ")
	_, err = fmt.Scanf("%f %f %f", &x2, &y2, &z2)
	if err != nil {
		fmt.Println("에러:", err)
		return
	}

	point1 := Point3D{X: x1, Y: y1, Z: z1}
	point2 := Point3D{X: x2, Y: y2, Z: z2}

	distance := Distance(point1, point2)

	fmt.Printf("두 점 사이의 거리: %.2f\n", distance)
}
