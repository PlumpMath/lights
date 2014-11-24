package lights

import ("testing"
        . "gopkg.in/check.v1"
       )
func Test(t *testing.T) { TestingT(t) }

type LSuite struct{
   Grid [][]uint8
   gridIn string 
   gameIn string
   game Game
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
  
  s.game = Game{s.Grid,1}
}

var _ = Suite(&LSuite{})


func (s *LSuite) TestParse(c *C){
  output := ParseGrid(s.gridIn ,8,8)
  c.Assert(output, DeepEquals, s.Grid)
}

