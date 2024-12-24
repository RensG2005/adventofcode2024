package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func main() {
  regex := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
  regex2 := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|don't\(\)|do\(\)`)
  numberRegex := regexp.MustCompile(`[0-9]+`)
   file, err := os.Open("input3.txt")
   if err != nil {
       log.Fatal(err)
     }
   defer file.Close()

   var strBuilder strings.Builder
   scanner := bufio.NewScanner(file)
   for scanner.Scan() {
       strPart := strings.Fields(scanner.Text())
       strBuilder.WriteString(strings.Join(strPart, ""))
     }
   if err := scanner.Err(); err != nil {
       log.Fatal(err)
     }

   str := strBuilder.String()
   part1(str, regex, numberRegex)
  
   part2(str, regex2, numberRegex)
 }

 func part1(str string, re *regexp.Regexp, numberRegex *regexp.Regexp) {
   ans := 0
   matches := re.FindAllStringSubmatch(str, -1)
   for i := 0; i < len(matches); i++ {
       numbers := numberRegex.FindAllString(matches[i][0], -1)
       n0, err := strconv.Atoi(numbers[0])
       if err != nil {
           continue
         }
       n1, err := strconv.Atoi(numbers[1])
       if err != nil {
           continue
         }
       ans += (n0 * n1)
     }
   fmt.Println(ans)
}

func part2(str string, regex *regexp.Regexp, numberRegex *regexp.Regexp) {
  ans := 0
  pause := true
  matches := regex.FindAllStringSubmatch(str, -1)
  for i := 0; i < len(matches); i++ {
    if(matches[i][0] == "do()") {
      pause = true
      continue
    }
    if(matches[i][0] == "don't()") {
      pause = false
      continue
    }
     numbers := numberRegex.FindAllString(matches[i][0], -1)
    if len(numbers) != 2 { continue }
     n0, err := strconv.Atoi(numbers[0])
     if err != nil {
       continue
     }
     n1, err := strconv.Atoi(numbers[1])
     if err != nil {
       continue
     }
    if(pause) {
     ans += (n0 * n1)
    }
   }
  fmt.Println(ans)
}
