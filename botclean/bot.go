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
  cost int
}

func main() {

  botPos := Point{0,0,0}
  var lowestCost *Point

  reader := bufio.NewReader(os.Stdin)

  line, _ := reader.ReadString('\n')
  fmt.Sscanf(line, "%d %d", &botPos.x, &botPos.y)

  for x := 0; ; x++ {
    line, err := reader.ReadString('\n')
    if err != nil && err != io.EOF { panic(err) }
    for y := 0; y < len(line); y++ {
      ch := line[y]
      if (ch == 'd') {
        cost := abs(botPos.x - x) + abs(botPos.y - y)
        point := Point{x, y, cost}
        if (cost == 0) {
          fmt.Println("CLEAN")
          os.Exit(0)
        } else if lowestCost == nil || point.cost < lowestCost.cost {
          lowestCost = &point
        }
      }
    }
    if err == io.EOF {
      break
    }
  }
  if botPos.x < lowestCost.x {
    fmt.Println("DOWN")
  } else if botPos.x > lowestCost.x {
    fmt.Println("UP")
  } else if botPos.y > lowestCost.y {
    fmt.Println("LEFT")
  } else if botPos.y < lowestCost.y {
    fmt.Println("RIGHT")
  }
}

func abs(i int) int {
  if (i < 0) {
    return -1 * i
  }
  return i
}
