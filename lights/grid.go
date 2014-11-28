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
  width  uint
  height uint
}

type Move struct {
  x, y uint
  value int 
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

func ParseGrid(input string, width uint, height uint) [][]uint8 {
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

func ParseGame(input string) Game {
  lines     := strings.Split(input, "\n")
  playerStr := lines[0]
  height    := uint(len(lines[1:]))
  width     := uint(len(lines[1]))
  player, err := strconv.Atoi(playerStr)
  if err != nil {
    log.Fatal(err)
  } 
  grid := ParseGrid(strings.Join(lines[1:],"\n"), width, height)
  return Game{grid, uint8(player), width, height}
}

func flip(bit uint8) uint8{
  if bit == 1{
    return 0
  } else {
    return 1
  }
  
}

func (game Game) Blast(x,y uint)(Game){
  g := game.copyGame()
  g.grid[y][x] = 0
  g.grid[y][x+1] = flip(game.At(x+1,y)) 
  g.grid[y+1][x] = flip(game.At(x,y+1)) 
  return g
}

func (game Game) BlastValue(x,y uint)(int){
  return int(game.At(x,y) + game.At(x+1,y) + game.At(x,y+1))
}

func (game Game) CanBlast(x,y uint)(bool){
  fmt.Println(x, y, game.At(x,y))
  
  if game.At(x,y) == 0 {
    return false
  } else {
    return true
  }
}

func (game Game) PossibleMoves()([]Move){
  fmt.Printf("%v", game)
  slc := make([]Move, 0)
  for x := uint(0); x < game.width; x++ {
    for y := uint(0); y < game.height; y++ {
      if game.CanBlast(uint(x),uint(y)) && game.BlastValue(uint(x), uint(y) ) > 0 {
        slc =  append(slc,Move{uint(x),uint(y), game.BlastValue(uint(x),uint(y))})
      }
    }
  }
  return slc
}
func (game Game) copyGame()(Game){
  slc := make([][]uint8, game.width)  
  for x := range(game.grid){
    slc[x] = make([]uint8, game.height)
    copy(slc[x],game.grid[x])
  }
  return Game{slc, game.player, game.width, game.height} 
}

func (game Game) At(x,y uint)(uint8){
  if x < game.width && y < game.height {
    return game.grid[y][x]
  } else {
    return 0
  }
}