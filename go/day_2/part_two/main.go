package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Game struct {
	id       int8
	cubesSet []Cubes
}

func (g *Game) addCubes(cubes Cubes) {
	g.cubesSet = append(g.cubesSet, cubes)
}

func (g *Game) getNewMaxedValuesCubes() *Cubes {
	newCubes := Cubes{}

	for _, cubes := range g.cubesSet {
		if cubes.red > newCubes.red {
			newCubes.red = cubes.red
		}

		if cubes.green > newCubes.green {
			newCubes.green = cubes.green
		}

		if cubes.blue > newCubes.blue {
			newCubes.blue = cubes.blue
		}
	}

	return &newCubes
}

type Cubes struct {
	red   int8
	green int8
	blue  int8
}

func (c *Cubes) powerOfCubes() int {
	powerOfCubes := int(c.red) * int(c.blue) * int(c.green)

	return powerOfCubes
}

func main() {
	result := 0

	inputFilePath := path.Join("..", "real.input")
	filePointer, errFileOpen := os.Open(inputFilePath)
	check(errFileOpen)

	fileScanner := bufio.NewScanner(filePointer)

	for fileScanner.Scan() {
		game := NewGameFromText(fileScanner.Text())
		maxedValuesCube := game.getNewMaxedValuesCubes()
		powerOfCubes := maxedValuesCube.powerOfCubes()
		result += powerOfCubes
	}

	fmt.Println(result)
}

func NewGameFromText(text string) *Game {
	game := Game{}

	gameInfoArray := strings.Split(text, ":")
	gameIdText := strings.Replace(gameInfoArray[0], "Game ", "", -1)
	gameId, errGameIdAtoi := strconv.Atoi(gameIdText)
	check(errGameIdAtoi)

	game.id = int8(gameId)

	cubesSetsSplited := strings.Split(gameInfoArray[1], ";")
	for _, cubeSetText := range cubesSetsSplited {
		cubes := *NewCubesFromCubeSetText(cubeSetText)
		game.addCubes(cubes)
	}

	return &game
}

func NewCubesFromCubeSetText(text string) *Cubes {
	cubes := Cubes{}

	cubeTextSplited := strings.Split(text, ",")

	for _, cubeText := range cubeTextSplited {
		cubeTextTrimmed := strings.Trim(cubeText, " ")
		cubeInfo := strings.Split(cubeTextTrimmed, " ")

		quantity, cubeInfoAtoiErr := strconv.Atoi(cubeInfo[0])
		check(cubeInfoAtoiErr)

		if cubeInfo[1] == "blue" {
			cubes.blue = int8(quantity)
		}

		if cubeInfo[1] == "red" {
			cubes.red = int8(quantity)
		}

		if cubeInfo[1] == "green" {
			cubes.green = int8(quantity)
		}
	}

	return &cubes
}
