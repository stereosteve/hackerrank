package main

import (
  "fmt"
  "bufio"
  "os"
  "io"
)

type Point struct {
  x int
  y int
}

func main() {

  botPos := Point{}
  reader := bufio.NewReader(os.Stdin)

  line, _ := reader.ReadString('\n')
  fmt.Sscanf(line, "%d %d", &botPos.x, &botPos.y)

  for x := 0; ; x++ {
    line, err := reader.ReadString('\n')
    if err != nil && err != io.EOF { panic(err) }
    for y := 0; y < len(line); y++ {
      ch := line[y]
      if (ch == 'd') {
        if (x == botPos.x && y == botPos.y) {
          fmt.Println("CLEAN")
        } else if (x > botPos.x) {
          fmt.Println("DOWN")
        } else if (x < botPos.x) {
          fmt.Println("UP")
        } else if (y > botPos.y) {
          fmt.Println("RIGHT")
        } else if (y < botPos.y) {
          fmt.Println("LEFT")
        }
      }
    }
    if err == io.EOF {
      break
    }
  }
}
