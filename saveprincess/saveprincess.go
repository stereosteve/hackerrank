package main

import (
  "os"
  "fmt"
  "bufio"
  "io"
)

type pos struct {
  x int
  y int
}

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT

 var size int
 _, err := fmt.Scanf("%d", &size)
 if err != nil { panic(err) }

 var mario_pos pos
 var princess_pos pos

 reader := bufio.NewReader(os.Stdin)
 for y := 0; y < size; y ++ {
   line, err := reader.ReadString('\n')
   if err != nil && err != io.EOF { panic(err) }
   for x, char := range line {
     if (char == 'm') {
       mario_pos.x = x
       mario_pos.y = y
     } else if (char == 'p') {
       princess_pos.x = x
       princess_pos.y = y
     }
   }
 }

 for (mario_pos.x < princess_pos.x) {
   fmt.Println("RIGHT")
   mario_pos.x++
 }
 for (mario_pos.x > princess_pos.x) {
   fmt.Println("LEFT")
   mario_pos.x--
 }
 for (mario_pos.y < princess_pos.y) {
   fmt.Println("DOWN")
   mario_pos.y++
 }
 for (mario_pos.y > princess_pos.y) {
   fmt.Println("UP")
   mario_pos.y--
 }

}

