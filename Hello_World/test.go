package main

import(
  "fmt"
  "errors"
  "math"
)

type person struct{
  name string
  age int
}
func main(){
  y := 0
  fmt.Println(y)
  inc(&y)
  fmt.Println(y)

  p:=person{name: "John", age:16}
  fmt.Println(p)
  fmt.Println(p.name)
  fmt.Println(&p.name)
  fmt.Println(&p)
  z := []int{1,2,3,4,5}
  z = append(z , 20)
  z[0] = 10
  for index, value := range z{
    fmt.Println(index, value)
  }
  hashmap := make(map[string]int)
  hashmap["a"] = 1
  hashmap["b"] = 2
  hashmap["c"] = 3
  hashmap["d"] = 100
  x := "a"
  for key, value := range hashmap{
    fmt.Println(key, value)
  }
  fmt.Println(hashmap[x])
  for i:=0; i<5; i++{
    fmt.Println(i)
  }
  fmt.Println(sum(2, 3))

  result, error := sqrt(-1)
  if (error != nil){
    fmt.Println(error)
  } else{
    fmt.Println(result)
  }

  result, error = sqrt(16)
  if (error != nil){
    fmt.Println(error)
  } else{
    fmt.Println(result)
  }

}

func inc(x *int){
  *x++
}

func sum (x int, y int) int{
  return x+y
}

func sqrt( x float64) (float64, error){
  if x<0{
    return 0, errors.New("Negative numbers don't have real square root")
  } else{
    return math.Sqrt(x), nil
  }
}
