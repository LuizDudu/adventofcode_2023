package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
func main() {
	maxCubes := Cubes{
		red:   12,
		green: 13,
		blue:  14,
	}

	inputPath := filepath.Join("..", "real.input")
	file, err := os.Open(inputPath)
	check(err)

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameFullText := scanner.Text()
		info := strings.Split(gameFullText, ":")
		cubesText := strings.Split(info[1], ";")

		allCubesFits := true
		for _, setOfCubes := range cubesText {
			newCubes := NewCubesFromSetOfCubes(setOfCubes)
			if newCubes.green > maxCubes.green || newCubes.red > maxCubes.red || newCubes.blue > maxCubes.blue {
				allCubesFits = false
				break
			}
		}

		if allCubesFits {
			gameText := strings.Split(info[0], " ")
			gameId, errAtoi := strconv.Atoi(gameText[1])
			check(errAtoi)
			result += gameId
		}
	}

	fmt.Println(result)
}

type Cubes struct {
	blue  int8
	red   int8
	green int8
}

func (cubes *Cubes) addCube(color string, quantity int8) {
	if color == "red" {
		cubes.red += quantity
		return
	}

	if color == "blue" {
		cubes.blue += quantity
		return
	}

	if color == "green" {
		cubes.green += quantity
		return
	}
}

func NewCubesFromSetOfCubes(setOfCubes string) *Cubes {
	cubes := Cubes{}

	cubesAsText := strings.Split(setOfCubes, ",")
	for _, cubeAsText := range cubesAsText {
		quantityColor := strings.Split(cubeAsText, " ")
		quantity, _ := strconv.Atoi(quantityColor[1])
		cubes.addCube(quantityColor[2], int8(quantity))
	}

	return &cubes
}
