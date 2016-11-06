package main

import (
	"fmt"
)

func main() {
	nums := []int{9, 3, 5, 7, 4, 1, 2, 8, 0, 6}
	
  //nums = bubbleSort(nums)
  //nums = insertionSort(nums)
  nums = mergeSort(nums)
  
  fmt.Println(nums)
}

func bubbleSort(datas []int) ([]int) {
  for i:=0; i<len(datas)-2; i++ {
    for j:=1; j<len(datas)-i ;j++ {
      if datas[j-1] > datas[j] {
        tmp := datas[j]
        datas[j] = datas[j-1]
        datas[j-1] = tmp
      }
    }
  }
  return datas
}

func insertionSort(datas []int) ([]int) {
  for i:=1; i<len(datas); i++ {
    tmp := datas[i]
    if datas[i-1] > tmp {
      j := i
      for {
        datas[j] = datas[j-1]
        j--
        if (j<1 || datas[j-1] < tmp){
          break
        }
      }
      datas[j] = tmp
    }
  }
  return datas
}

func mergeSort(datas []int) ([]int) {
  list := make([]int, len(datas))
  copy(list, datas)
  
  half := func(lst []int)([]int, []int){
    a := lst[:len(lst)/2]
    b := lst[len(lst)/2:]
    return a, b
  }
  
  merge := func(lst1 []int, lst2 []int)([]int){
    len1, len2 := len(lst1), len(lst2)
    result := make([]int, len1+len2)
    i, j := 0, 0

    for k:=0; k<len(result); k++ {
      if(j>=len(lst2) || (i<len(lst1) && lst1[i] < lst2[j])){
        result[k] = lst1[i]
        i++
      } else {
        result[k] = lst2[j]
        j++
      }
    }
    return result
  }
  
  var mg func([]int)([]int)
  mg = func(lst []int)([]int){
    if len(lst) <= 1 {
      return lst
    }
    x, y := half(lst)
    return merge(mg(x), mg(y))
  }

  return mg(list)
}

func shellSort(datas []int) ([]int) {
  return datas
}