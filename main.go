package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Point struct {
	x, y int
}

type Rectangle struct {
	start, end Point
	id         int
}

type Rects struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

type InputRectangles struct {
	Rects []Rects `json:"rects"`
}

func newRectangle(id, x, y, w, h int) Rectangle {
	return Rectangle{Point{x, y}, Point{x + w, y + h}, id}
}

func isIntersect(r1, r2 Rectangle) bool {
	// check if the r2 is on side of r1
	if r2.end.x < r1.start.x || r2.start.x > r1.end.x {
		return false
	}

	// check if the r2 is on the top or botton of r1
	if r2.end.y < r1.start.y || r2.start.y > r1.end.y {
		return false
	}

	return true
}

func checkIntersection(arrRect []Rectangle, rect Rectangle) bool {
	for _, r := range arrRect {
		if !isIntersect(r, rect) {
			return false
		}
	}

	return true
}

func printIntersected(arrRect []Rectangle, rect Rectangle) {
	if len(arrRect) == 1 {
		fmt.Print("Between ", arrRect[0].id, " and ", rect.id)
	} else {
		fmt.Print("Between ")
		for i := 0; i < len(arrRect); i++ {
			fmt.Print(arrRect[i].id, ",")
		}
		fmt.Print(" and ", rect.id)
	}

	startPoint, width, height := findIntersectionPoints(arrRect, rect)

	fmt.Println(" at (", startPoint.x, "\b,", startPoint.y, "), w=", width, "\b, h=", height, "\b.")
}

func min(vals []int) int {
	min := 999

	for _, val := range vals {
		if val < min {
			min = val
		}
	}

	return min
}

func max(vals []int) int {
	max := 0

	for _, val := range vals {
		if val > max {
			max = val
		}
	}

	return max
}

func findIntersectionPoints(rectArr []Rectangle, rect Rectangle) (Point, int, int) {
	x1s := []int{}
	y1s := []int{}

	x2s := []int{}
	y2s := []int{}

	for i := 0; i < len(rectArr); i++ {
		x1s = append(x1s, rectArr[i].start.x)
		y1s = append(y1s, rectArr[i].start.y)

		x2s = append(x2s, rectArr[i].end.x)
		y2s = append(y2s, rectArr[i].end.y)
	}

	x1s = append(x1s, rect.start.x)
	y1s = append(y1s, rect.start.y)

	x2s = append(x2s, rect.end.x)
	y2s = append(y2s, rect.end.y)

	startPoint := Point{max(x1s), max(y1s)}
	endPoint := Point{min(x2s), min(y2s)}

	width := endPoint.x - startPoint.x
	height := endPoint.y - startPoint.y

	return startPoint, width, height
}

func convertInputRectangles(inputRectangles InputRectangles) ([]Rectangle, error) {
	var rectangles []Rectangle

	inputRectangleSize := len(inputRectangles.Rects)
	if inputRectangleSize > 10 {
		return rectangles, errors.New("Input has more than 10 Rectangles!")
	}

	for i := 0; i < len(inputRectangles.Rects); i++ {
		rect := newRectangle(i+1, inputRectangles.Rects[i].X,
			inputRectangles.Rects[i].Y,
			inputRectangles.Rects[i].W,
			inputRectangles.Rects[i].H)

		rectangles = append(rectangles, rect)
	}

	return rectangles, nil
}

func findRectangleIntersection(inputRectangles_ InputRectangles) {

	inputRectangles, err := convertInputRectangles(inputRectangles_)

	if err != nil {
		return
	}

	var existingRectangles [100][]Rectangle
	var existingRectanglesSize = 0
	var inputRectanglesSize = 4

	inputRectangles[0] = newRectangle(1, 100, 100, 250, 80)
	inputRectangles[1] = newRectangle(2, 120, 200, 250, 150)
	inputRectangles[2] = newRectangle(3, 140, 160, 250, 100)
	inputRectangles[3] = newRectangle(4, 160, 140, 350, 190)

	for i := 0; i < inputRectanglesSize; i++ {
		if i != 0 {

			existingRectanglesCpy := existingRectangles
			existingRectanglesCpySize := existingRectanglesSize

			for j := 0; j < existingRectanglesCpySize; j++ {
				if checkIntersection(existingRectanglesCpy[j], inputRectangles[i]) {

					printIntersected(existingRectanglesCpy[j], inputRectangles[i])

					tmp := existingRectanglesCpy[j]
					tmp = append(tmp, inputRectangles[i])
					existingRectangles[existingRectanglesSize] = tmp
					existingRectanglesSize++
				}
			}
		}

		existingRectangles[existingRectanglesSize] = []Rectangle{inputRectangles[i]}
		existingRectanglesSize++
	}
}

func importJSON() InputRectangles {
	data, err := ioutil.ReadFile("input.json")

	var inputRectangles InputRectangles

	err = json.Unmarshal(data, &inputRectangles)

	if err != nil {
		fmt.Println("error :", err)
	}

	return inputRectangles
}

func main() {
	inputRectangles := importJSON()

	findRectangleIntersection(inputRectangles)
}
