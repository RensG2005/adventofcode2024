package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func removeNthElementTemp(slice []int, n int) []int {
    copySlice := append([]int(nil), slice...)

  if n < 0 || n >= len(copySlice) {
      return copySlice
    }

    return append(copySlice[:n], copySlice[n+1:]...)
}

func main() {
  file, err := os.Open("input2.txt")
  if err != nil {
    log.Fatal(err)
  }

  ans := 0
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    sequence := strings.Fields(scanner.Text())
    var intSeq []int
    for _, str := range sequence {
      num, err := strconv.Atoi(str)
      if err != nil {
        continue
      }
      intSeq = append(intSeq, num)
    }

    //part 1
    if(checkSequence(intSeq)) {
      ans++
    } else {
      //part2
      for i := 0; i < len(intSeq); i++ {
        if(checkSequence(removeNthElementTemp(intSeq, i))) {
          ans++
          break
        }
      }
    }
    ans += checkSequenceTwo(intSeq)
  }
  
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  fmt.Println(ans)
}

func checkSequence(sequence []int) bool {
  increasing := false
  decreasing := false
  for i := 0; i < len(sequence) - 1; i++ {
    if (sequence[i] > sequence[i + 1]) {
      if increasing {
        return false
      }
      if (sequence[i] > sequence[i+1] + 3) { return false }
      decreasing = true
    } else if (sequence[i] < sequence[i + 1]) {
      if decreasing {
        return false
      }
      if (sequence[i] < sequence[i+1] - 3) { return false }
      increasing = true
    } else if (sequence[i] == sequence[i+1]) {
      return false
    }
  }
  fmt.Println(sequence)
  return true
}
