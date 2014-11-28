package lights

import ("testing"
        . "gopkg.in/check.v1"
        "fmt"
       )
func Test(t *testing.T) { TestingT(t) }

type LSuite struct{
   Grid [][]uint8
   gridIn string 
   gameIn string
   smallGrid [][]uint8
   smallGridBlast [][]uint8
   smallGame Game
   smallGridMoves  []Move
   game Game
}
func (s* LSuite) SetUpTest(c *C){
  s.Grid = [][]uint8{
    {0,1,1,0,1,1,0,0},
    {0,1,0,1,0,1,0,1},
    {1,0,1,1,1,1,0,0},
    {0,0,0,0,1,1,1,1},
    {1,0,1,0,1,0,1,0},
    {0,0,0,0,0,0,0,1},
    {1,1,1,0,0,0,1,1},
    {0,0,1,0,1,0,0,0},
  }
  s.smallGrid = [][]uint8{
    {0,1,1,1},
    {0,1,0,0},
    {0,1,1,0},
    {0,1,1,0},
  } 
  
  s.smallGridMoves = []Move{
    Move{1,0,3},
    Move{1,1,2},
    Move{1,2,3},
    Move{1,3,2},
    Move{2,0,2},
    Move{2,2,2},
    Move{2,3,1},
    Move{3,0,1},
  } 
  
  s.smallGridBlast = [][]uint8{
    {0,0,0,1},
    {0,0,0,0},
    {0,1,1,0},
    {0,1,1,0},
  } 

    s.game = Game{s.Grid,1, 8,8}
    s.smallGame = Game{s.smallGrid,1, 4,4}
}
func (s* LSuite) SetUpSuite(c *C){
  s.gameIn = `1
01101100
01010101
10111100
00001111
10101010
00000001
11100011
00101000`
  
  s.gridIn = `01101100
01010101
10111100
00001111
10101010
00000001
11100011
00101000`

}

var _ = Suite(&LSuite{})


func (s *LSuite) TestParseGrid(c *C){
  output := ParseGrid(s.gridIn ,8,8)
  c.Assert(output, DeepEquals, s.Grid)
}

func (s *LSuite) TestParseGame(c *C){
  output := ParseGame(s.gameIn)
  c.Assert(output, DeepEquals, s.game)
}

func (s *LSuite) TestCanBlastFalse(c *C){
  output := s.smallGame.CanBlast(3,1)
  c.Assert(output, Equals, false)
}

func (s *LSuite) TestCanBlastTrue(c *C){
  output := s.smallGame.CanBlast(3,0)
  c.Assert(output, Equals, true)
}

func (s *LSuite) TestAt(c *C){
  output := s.smallGame.At(0,0)
  c.Assert(output, Equals, uint8(0x0))
  output = s.smallGame.At(3,0)
  c.Assert(output, Equals, uint8(0x1))
}

func (s *LSuite) TestPossibleMoves(c *C) {
  output := s.smallGame.PossibleMoves()
  c.Assert(output, DeepEquals, s.smallGridMoves)
}

func (s *LSuite) TestBlast(c *C){
  output := s.smallGame.Blast(1,0) 
  c.Assert(output.grid, DeepEquals, s.smallGridBlast)
  c.Assert(s.smallGame.grid, DeepEquals, s.smallGrid)
  fmt.Println(s.smallGame)
}

