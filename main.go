package main

import (
	"encoding/json"
	"errors"
	"flag"
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

func printIntersected(arrRect []Rectangle, rect Rectangle, intersectionIndex int) {
	if len(arrRect) == 1 {
		fmt.Print("\t", intersectionIndex, ": Between ", arrRect[0].id, " and ", rect.id)
	} else {
		fmt.Print("\t", intersectionIndex, ": Between ")
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

	fmt.Println("Input:")

	for i := 0; i < len(inputRectangles.Rects); i++ {
		rect := newRectangle(i+1, inputRectangles.Rects[i].X,
			inputRectangles.Rects[i].Y,
			inputRectangles.Rects[i].W,
			inputRectangles.Rects[i].H)

		fmt.Println("\t", i, ": Rectangle at (", rect.start.x, ",", rect.start.y, "), w=", rect.start.x+rect.end.x, ", h=", rect.start.y+rect.end.y)

		rectangles = append(rectangles, rect)
	}

	return rectangles, nil
}

func findRectangleIntersection(inputRectangles_ InputRectangles) {

	inputRectangles, err := convertInputRectangles(inputRectangles_)

	if err != nil {
		fmt.Println(err)
		return
	}

	var existingRectangles [100][]Rectangle
	var existingRectanglesSize = 0
	inputRectanglesSize := len(inputRectangles)

	fmt.Println("Intersections:")
	intersectionIndex := 0

	for i := 0; i < inputRectanglesSize; i++ {
		if i != 0 {

			existingRectanglesCpy := existingRectangles
			existingRectanglesCpySize := existingRectanglesSize

			for j := 0; j < existingRectanglesCpySize; j++ {
				if checkIntersection(existingRectanglesCpy[j], inputRectangles[i]) {

					printIntersected(existingRectanglesCpy[j], inputRectangles[i], intersectionIndex)
					intersectionIndex++

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

func importJSON(filename string) (InputRectangles, error) {
	data, err := ioutil.ReadFile(filename)

	var inputRectangles InputRectangles

	err = json.Unmarshal(data, &inputRectangles)

	if err != nil {
		fmt.Println("error : File reading error, please check your input file")
		return inputRectangles, err
	}

	return inputRectangles, nil
}

func checkArguments() (string, error) {
	var inputfilename string
	flag.StringVar(&inputfilename, "file", "input.json", "JSON file input")
	flag.Parse()

	if inputfilename == "" {
		return inputfilename, errors.New("No file input was provided!")
	}

	return inputfilename, nil
}

func main() {

	jsonFileInput, err := checkArguments()
	if err != nil {
		fmt.Println(err)
		return
	}

	inputRectangles, err := importJSON(jsonFileInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	findRectangleIntersection(inputRectangles)
}
