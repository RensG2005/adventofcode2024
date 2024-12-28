package utils

import (
	"fmt"
	"reflect"
)


func SlicesEqual(slice1, slice2 []int) bool {
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

func PrettyPrint2D(slice interface{}) {
    // Check if the input is a slice of slices
    val := reflect.ValueOf(slice)
    if val.Kind() != reflect.Slice || val.Len() == 0 || val.Index(0).Kind() != reflect.Slice {
        fmt.Println("Input must be a 2D slice ([][]int or [][]string)")
        return
      }

    // Determine the type of the inner elements
    innerKind := val.Index(0).Index(0).Kind()

    switch innerKind {
    case reflect.Int:
      for i := 0; i < val.Len(); i++ {
            row := val.Index(i)
            for j := 0; j < row.Len(); j++ {
                fmt.Printf("%4d", row.Index(j).Int())
              }
            fmt.Println()
          }
    case reflect.String:
      for i := 0; i < val.Len(); i++ {
            row := val.Index(i)
            for j := 0; j < row.Len(); j++ {
                fmt.Printf("%s ", row.Index(j).String())
              }
            fmt.Println()
          }
    default:
      fmt.Println("Unsupported type: only [][]int or [][]string are supported")
    }
}

