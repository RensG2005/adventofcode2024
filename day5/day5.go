package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// func contains(slice []int, item int, integerList []int) int{
//  for _, v := range slice {
//      if v == item {
//          for j, k := range integerList {
//            if k == v {
//              return j
 //           }
 //         }
 //     }
  //}
  //return -1
//}
func contains(slice []int, target int) bool {
    for _, v := range slice {
        if v == target {
            return true
          }
      }
    return false
}

func addUnique(slice []int, value int) []int {
    for _, v := range slice {
        if v == value {
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
  //    if(contains(seen, intList[i][j], intList[i]) > 0) {
   //     wrong = true
    //    continue
    //  }
    }
    if(!wrong) {
      middleIndex := len(intList[i]) / 2
      *ans += intList[i][middleIndex]
    }

  }
  return 0
}

func reorderUpdate(update []int, rules map[int][]int) []int {
    orderedUpdate := make([]int, len(update))
    copy(orderedUpdate, update)

    for swapped := true; swapped; {
        swapped = false
        for i := 0; i < len(orderedUpdate)-1; i++ {
            post := orderedUpdate[i]
            pre := orderedUpdate[i+1]
            if contains(rules[pre], post) && !contains(rules[post], pre) {
                orderedUpdate[i+1], orderedUpdate[i] = orderedUpdate[i], orderedUpdate[i+1]
                swapped = true
              }
          }
      }
    return orderedUpdate
}

func slicesEqual(slice1, slice2 []int) bool {
    if len(slice1) != len(slice2) {
        return false
      }
    for i := range slice1 {
        if slice1[i] != slice2[i] {
            return false
          }
      }
    
    return true
}

func part2Search(intList []int, rules map[int][]int, ans *int, wrong bool)  int {
  //seen := []int{} 
  //for j := len(intList) - 1; j >= 0; j-- {
  //  for _, v := range rules[intList[j]] {
  //    seen = addUnique(seen, v) 
  //  }
  //}
  //for j := len(intList) - 1; j >= 0; j-- {
  //  if(contains(seen, intList[j], intList) > 0) {
  //    fmt.Println(seen, intList)
  //  }
  //}
  //return 0

  reorderd := reorderUpdate(intList, rules)
  if(!slicesEqual(intList, reorderd)) {
    fmt.Println(reorderd)
    middleIndex := len(reorderd) / 2
    *ans += reorderd[middleIndex]
  }
  return 0
}

func part2(rules map[int][]int, intList [][]int, ans *int, wrong bool) int {
  for i := 0; i<len(intList); i++ {
    part2Search(intList[i], rules, ans, wrong) 
  }
  return 0
}

func main() {
  file, err := os.Open("input5.txt")
  //file, err := os.Open("test.txt")
  if err != nil {
    log.Fatal(err)
  }

  ans := 0
  defer file.Close()
  scanner := bufio.NewScanner(file)
  rules := make(map[int][]int) 
  intList := [][]int{}
  second := false
 
  //parse input
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
  //challenge
  part2(rules, intList, &ans, false)
  //part1(rules, intList, &ans)

  //answer
  fmt.Println(ans)
}

