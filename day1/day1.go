package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
  left := []int{}
  right := []int{}
  file, err := os.Open("input1.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    words := strings.Fields(scanner.Text())
    word1, err := strconv.Atoi(words[0])
    if err != nil {
      return
    }
    left = append(left, word1)
    word2, err := strconv.Atoi(words[1])
    if err != nil {
        return
    }
    right= append(right, word2)
    strconv.Atoi(words[1])
  }
  part1(left, right)
  part2(left,right) 
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

func part1(left []int, right []int) {
    sort.Ints(left)
    sort.Ints(right)
    sum := 0
    i := 0
    for i < len(left) {
        diff := int(math.Abs(float64(left[i] - right[i])))
        sum += diff
        i++
    }
    fmt.Println(sum)
}

func part2(left []int, right []int) {
  ans := 0
  for i := 0; i < len(left); i++ {
    for j := 0; j < len(right) && right[j] <= left[i]; j++ {
      if left[i] == right[j] {
        ans += left[i]
      }
    }
  }
  fmt.Println(ans)
}

