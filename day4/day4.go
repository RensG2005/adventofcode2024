package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)



func dfs(visited [140][140]int, grid [140][140]rune, searchWord [4]rune, i int, j int, index int, direction [2]int) int {
  if index == 4 { 
    return 1
  }
  if i<0 || i >= len(grid) || j < 0 || j >= len(grid[0]) ||  visited[i][j] == 1 || grid[i][j] != searchWord[index] {
    return 0
  }
  visited[i][j] = 1

  dx := direction[0]
  dy := direction[1]
  newI := i+dx
  newJ := j+dy
  if dfs(visited, grid, searchWord, newI, newJ, index+1, direction) == 1 {
    return 1
  }

  visited[i][j] = 0
  return 0
}

func part1(grid [140][140]rune, ans *int) {
  searchWord := [4]rune{'X', 'M', 'A', 'S'}
  visited := [140][140]int{}
   directions := [][2]int{
    {0, 1},
    {1, 0},
    {-1, 0},
    {0, -1},
    {-1, 1},
    {1, -1},
    {-1, -1},
    {1, 1},
  } 
  for i := 0; i < len(grid); i++ {
    for j := 0; j < len(grid[i]); j++ {
      if(grid[i][j] == searchWord[0]) {
        for k := 0; k < len(directions); k++ {
          *ans += dfs(visited, grid, searchWord, i, j, 0, directions[k])
        }
      }
    } 
  }
}

func main() {
  file, err := os.Open("input4.txt")
  if err != nil {
    log.Fatal(err)
  }

  grid := [140][140]rune{}
  ans := 0
  defer file.Close()
  scanner := bufio.NewScanner(file)
  
  row := 0 
  for scanner.Scan() {
    sequence := strings.Fields(scanner.Text())
    for i, char := range sequence[0] { 
      if i < 140 { 
        grid[row][i] = char
      }
    }
    row++
    if row >= 140 { // Ensure we don't exceed grid row limits
      break
    }
  }

  part1(grid, &ans)
  

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  fmt.Println(ans)
}
