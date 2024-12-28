package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func move(grid []string, i int, j int, curDirection int) int {
  direction := [][2]int {
    {-1, 0},
    {0, 1},
    {1, 0},
    {0, -1},
  }

  runes := []rune(grid[i])
  runes[j] = 'X' 
  grid[i] = string(runes)

  newX := i + direction[curDirection][0]
  newY := j + direction[curDirection][1]

  if newX < 0 || newY < 0 || newX > len(grid) - 1 || newY > len(grid[0]) - 1 { 
    return 0
  }

  if(grid[newX][newY] == '#') {
    curDirection = (curDirection + 1) % 4
    newX  = i + direction[curDirection][0]
    newY  = j + direction[curDirection][1]
  }

  return move(grid, newX, newY, curDirection)
}

func part1(grid []string, ans *int) {
  for i := 0; i<len(grid); i++ {
    for j := 0; j<len(grid[i]); j++ {
      if(string(grid[i][j]) == "^") {
        move(grid, i, j, 0)
      }
    }
  }
  for i := 0; i<len(grid); i++ {
    fmt.Println(grid[i])
    for j := 0; j<len(grid[i]); j++ {
      if(string(grid[i][j]) == "X") {
        *ans += 1
      }
  }
  }
}

func main() {
  //file, err := os.Open("input6.txt")
  file, err := os.Open("test.txt")

  if err != nil {
    log.Fatal(err)
  }

  var grid []string
  ans := 0
  defer file.Close()
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    rule := strings.Fields(scanner.Text())
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    grid = append(grid, rule[0])
  }
  //challenge
  //part2(rules, intList, &ans, false)
  part1(grid, &ans)

  //answer
  fmt.Println(ans)
}
