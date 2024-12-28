package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strconv"
	"strings"
)

func contains(slice []int, item int) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

func addUnique(slice []int, value int) []int {
    for _, v := range slice {
        if v == value {
            // Skip adding if the value is already in the slice
            return slice
        }
    }
    return append(slice, value)
}

func part1(rules map[int][]int, intList [][]int, ans *int) int {
  for i := 0; i<len(intList); i++ {
    wrong := false
    seen := []int{} 
    for j := len(intList[i]) - 1; j>=0; j-- {
      for _, v := range rules[intList[i][j]] {
        seen = addUnique(seen, v)
      }
      if(contains(seen, intList[i][j])) {
        wrong = true
        continue
      }
    }
    if(!wrong) {
      middleIndex := len(intList[i]) / 2
      *ans += intList[i][middleIndex]
    }

  }
  return 0
}

func main() {
  file, err := os.Open("input5.txt")
  if err != nil {
    log.Fatal(err)
  }

  ans := 0
  defer file.Close()
  scanner := bufio.NewScanner(file)
  rules := make(map[int][]int) 
  intList := [][]int{}
  second := false
  
  for scanner.Scan() {
    rule := strings.Fields(scanner.Text())
    if(len(rule) == 0) { second = true; continue }
    if(second) {
      var row []int
		  for _, part := range strings.Split(rule[0], ",") {
			if num, err := strconv.Atoi(part); err == nil {
				row = append(row, num)
        }
      }
      intList = append(intList, row)
      } else {
      rule1, err := strconv.Atoi(strings.Split(rule[0], "|")[0])
      if err != nil {log.Fatal(err)}
      rule2, err := strconv.Atoi(strings.Split(rule[0], "|")[1])
      if err != nil {log.Fatal(err)}
      rules[rule1] = append(rules[rule1], rule2)
    }
    
    if err := scanner.Err(); err != nil {
      log.Fatal(err)
    }
  }

  part1(rules, intList, &ans)
  fmt.Println(ans)
}

