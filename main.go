package main

import(
	"fmt"
	"math/rand/v2"
	// "time"
)

const (
	line  = 5
	column = 5
	sleep  = 100
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
			if rand.IntN(4) == 1{
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

func main(){
	world := GenerateWorld()
	Seed(world)
	PrintWorld(world)
	newWorld := GenerateWorld()
	for indexRow, row := range world {
		for indexCol, _ := range row {
			// fmt.Println(indexRow, indexCol)
			newWorld[indexRow][indexCol] = Rules(world, indexRow, indexCol)
		}
	}
	fmt.Println()
	PrintWorld(newWorld)
}