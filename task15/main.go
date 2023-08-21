package main

import (
	 "os"
)

// var justString string
 

func createHugeString(size int,f *os.File) error {
    buffer := make([]byte, size)
   _,err :=f.Write(buffer)
   return err
}


func someFunc() {
    f, err := os.Create("justString.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()
      err = createHugeString(1 << 30, f)  // лучше и быстрее держать большую строку в файле а не в глобальной переменной
	  if err != nil {
		panic(err)
	  }
}

// func createHugeString(size int) string {
//     buffer := make([]rune, size)
//     return string(buffer)
// }

// func someFunc() {
// 	v := []rune(createHugeString(1 << 30))  
// 	v =  v[:100]
// 	justString = string(v)
// }

func main() {
	someFunc()
}

 