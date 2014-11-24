package lights

import(
  "strings"
  "fmt"
  "strconv"
  "log"
)

type Game struct {
  grid [][]uint8 
  player uint8
}

func runeToUint(in rune)(out uint8){
  switch {
    case '0' == in:
      return 0
    case '1' == in:
      return 1
  }
  return 9
}

func ParseGrid(input string, width int, height int) [][]uint8 {
  rows := strings.Split(input, "\n")
  
  slc := make([][]uint8, height)
  for i := range slc {
    slc[i] = make([]uint8, width)
    fmt.Printf("%v", rows[i])
    for pos, char := range rows[i] {
       
      slc[i][pos] = runeToUint(char) 
    }
  }
  return slc 
}

func ParseGame(input string, width int, height int) Game {
  playerStr := strings.Split(input, "\n")[0]
  player, err := strconv.Atoi(playerStr)
  if err != nil {
    log.Fatal(err)
  } 
  grid := ParseGrid(input, width, height)
  return Game{grid, uint8(player)}
}