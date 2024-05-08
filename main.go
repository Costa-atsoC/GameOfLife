package main

import(
	"fmt"
	"math/rand/v2"
	// "os"
	"flag"
	"time"
)

var line = 1
var column = 1
var steps = 1
var sleep = 1

const (
	alive  = "🟩"
	dead   = "🟥"
) 	

type World[][] bool

// It's in the name honestly
func GenerateWorld() World{
	world := make(World, line);
	for i := range world {
		world[i] = make([]bool, column)
	}
	return world
}

//This is the function that will randomnly generate numbers
func Seed(w World){
	for _,row := range w {
		for col := range row{
			if rand.IntN(3) == 1{ //In here we change the probability of the cubes being alive
				row[col] = true
			}
		}
	}
}

func PrintWorld(w World){
	for _,row := range w {
		for col := range row{
			switch row[col]{
			case true:	
				fmt.Print(alive)
			case false:
				fmt.Print(dead)
			}
		}
		fmt.Println()
	}
}

func Rules(w World, x, y int) bool{
	fam := 0

	for i := x-1; i <= x+1; i++{
		for j := y-1; j <= y+1; j++{
			if i == x && j == y{
				continue
			}
			if i >= 0 && j >= 0 && i < line && j < column && w[i][j] {
				fam++
			}
		}
	}

	if fam < 2 || fam >= 4 && w[x][y]{
		return false
	}

	if w[x][y] == false && fam == 3{
		return true
	}
	if (fam == 2 || fam == 3) && w[x][y]{
		return true
	}
	return false
}

func CheckBoard(w World) bool{
	for _,row := range w {
		for col := range row{
			if row[col]{
				return true
			}
		}
	}
	return false
}

func main(){
	// Flags
	widthFlag := flag.Int("x", 5, "The width of the board.")
	heighFlag := flag.Int("y", 5, "The width of the board.")
	stepFlag := flag.Int("s", 1, "The number of iterations the program will go through.")
	timeFlag := flag.Int("t", 1000, "The time the program will sleep between the boards. This time is in milliseconds.")

	flag.Parse()

	line = *widthFlag
	column = *heighFlag
	steps = *stepFlag
	sleep = *timeFlag

	//To know how many steps it took until the world dies
	numberSteps := 0

	world := GenerateWorld()
	Seed(world)
	PrintWorld(world)
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	for i := 0; i< steps; i++ {
		newWorld := GenerateWorld()
		for indexRow, row := range world {
			for indexCol, _ := range row {
				newWorld[indexRow][indexCol] = Rules(world, indexRow, indexCol)
			}
		}
		fmt.Println()
		PrintWorld(newWorld)
		world = newWorld
		numberSteps++

		if!CheckBoard(newWorld) {
			fmt.Println("It took ", numberSteps, " until the world died.")
			break 
		}
		time.Sleep(time.Duration(sleep) * time.Millisecond)

	}
}
