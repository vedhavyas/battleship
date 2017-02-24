package battleship

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

/*
Since it was not specifically mentioned in the problem as to how the coordinate system of the
board works. Hence, we are assuming that (0,0) lies in the top left corner and (m,m) in bottom right
Final depiction of the board looks as follows

      (0,0) (0,1) (0,2)
(1, 0)  _     _     _
(2, 0)  _     _     _

*/

/*
emptyLocation is the empty block in the board represented by "_"
initialPosition is a ship's initial position/non-hit ship after match represented by "B"
hit is a position of a destroyed ship represented as "X"
missedHit is the position where there is no ship docked but was targeted by Opposition represented by "O"
*/
const (
	emptyLocation   = "_"
	initialPosition = "B"
	hit             = "X"
	missedHit       = "O"
)

//location holds the x, y point of a specific location on board
type location struct {
	X int
	Y int
}

//playerBoard is a depiction of the battle board of a player
//boardRepresentation is the tabular representation of the board with ships
//gridSize is number of rows/columns in the board
//totalShips is the total number of ships the player has on board
//shipPositions holds the location of each and every ship on the board
type playerBoard struct {
	boardRepresentation [][]string
	gridSize            int
	totalShips          int
	shipPositions       []location
}

//String is the formatted version of Board of a Player
func (b playerBoard) String() string {
	var buf bytes.Buffer
	for _, row := range b.boardRepresentation {
		var rowBuf bytes.Buffer
		for _, point := range row {
			_, err := fmt.Fprintf(&rowBuf, "%s ", point)
			if err != nil {
				return err.Error()
			}
		}

		_, err := fmt.Fprintf(&buf, "%s\n", strings.TrimSpace(rowBuf.String()))
		if err != nil {
			return err.Error()
		}
	}
	return buf.String()
}

//player is the representation a single player
//board represents the battle board of the player
//moves are the player moves on the opposition player's board
//totalMissiles are the number of chances of the player
type player struct {
	board         playerBoard
	moves         []location
	totalMissiles int
}

//opponentMove marks and return if the opposition player's move was hit or not a hit
func (p *player) opponentMove(loc location) bool {
	var isHit bool
	newData := missedHit
	x, y := loc.X, loc.Y
	if p.board.boardRepresentation[x][y] == initialPosition {
		isHit = true
		newData = hit
	}
	p.board.boardRepresentation[x][y] = newData
	return isHit
}

//getOpponentHits returns the total number of successful hits by the opponent
func (p *player) getOpponentHits(oppositionMoves []location) int {
	var hits int
	for _, move := range oppositionMoves {
		if p.opponentMove(move) {
			hits++
		}
	}

	return hits
}

//parsePlayers will parse the raw data into player representation
//raw data should be of the following format
//Grid Size - 5
//Total Ships - 5
//Player1 Ship Positions - 1:1,2:0,2:3,3:4,4:3
//Player2 Ship Positions - 0:1,2:3,3:0,3:4,4:1
//Total Missiles for each player - 5
//Player1 moves on Player2 - 0:1,4:3,2:3,3:1,4:1
//Player2 moves on Player1 - 0:1,0:0,1:2,2:3,4:3
func parsePlayers(gameData []string) (*player, *player, error) {
	gridSize, err := strconv.Atoi(gameData[0])
	if err != nil {
		return nil, nil, err
	}

	totalShips, err := strconv.Atoi(gameData[1])
	if err != nil {
		return nil, nil, err
	}

	totalMissiles, err := strconv.Atoi(gameData[4])
	if err != nil {
		return nil, nil, err
	}

	p1, err := newPlayer(gameData[2], gameData[5], gridSize, totalShips, totalMissiles)
	if err != nil {
		return nil, nil, err
	}

	p2, err := newPlayer(gameData[3], gameData[6], gridSize, totalShips, totalMissiles)
	if err != nil {
		return nil, nil, err
	}
	return p1, p2, nil
}

//newPlayer returns a Player representation of the given raw data
func newPlayer(playerShipPositions, playerMoves string,
	gridSize, totalShips, totalMissiles int) (*player, error) {

	shipPositions, err := parseLocations(playerShipPositions)
	if err != nil {
		return nil, err
	}

	playerBoard := newBoard(gridSize, totalShips, shipPositions)

	moves, err := parseLocations(playerMoves)
	if err != nil {
		return nil, err
	}

	return &player{
		board:         playerBoard,
		moves:         moves,
		totalMissiles: totalMissiles,
	}, nil
}

//parseLocations will parse the raw location string into slice of Locations
//locationData should of the format
//"x1,y1:x2,y2:x3,y3:x4,y4"
func parseLocations(locationData string) ([]location, error) {
	var locs []location
	rows := strings.Split(locationData, ",")
	for _, row := range rows {
		location, err := newLocation(row)
		if err != nil {
			return nil, err
		}
		locs = append(locs, location)
	}

	return locs, nil
}

//newLocation returns Location representation of a given raw location
//location should be of format "x,y"
func newLocation(locData string) (location, error) {
	splits := strings.Split(locData, ":")
	if len(splits) != 2 {
		return location{}, fmt.Errorf("Not a Valid location - %s", locData)
	}
	x, err := strconv.Atoi(splits[0])
	if err != nil {
		return location{}, err
	}

	y, err := strconv.Atoi(splits[1])
	if err != nil {
		return location{}, err
	}
	return location{
		X: x,
		Y: y,
	}, nil
}

//newBoard returns a playerBoard representation of raw data
func newBoard(gridSize, totalShips int, shipPositions []location) playerBoard {
	boardRepresentation := make([][]string, gridSize)
	for i := range boardRepresentation {
		row := make([]string, gridSize)
		for j := range row {
			row[j] = emptyLocation
		}
		boardRepresentation[i] = row
	}

	for _, position := range shipPositions {
		boardRepresentation[position.X][position.Y] = initialPosition
	}

	return playerBoard{
		boardRepresentation: boardRepresentation,
		gridSize:            gridSize,
		totalShips:          totalShips,
		shipPositions:       shipPositions,
	}
}

//GameResult holds the results of the game
//Player{1/2}Board is the representation of the board of the player
//Player{1/2}Hits is the total hits of a player on his opponent
type GameResult struct {
	Player1Board string
	Player2Board string
	Player1Hits  int
	Player2Hits  int
	Result       string
}

//PlayGame initiates the battleship game between two players
//gameData is the players raw data of format
//Grid Size - 5
//Total Ships - 5
//Player1 Ship Positions - 1:1,2:0,2:3,3:4,4:3
//Player2 Ship Positions - 0:1,2:3,3:0,3:4,4:1
//Total Missiles for each player - 5
//Player1 moves on Player2 - 0:1,4:3,2:3,3:1,4:1
//Player2 moves on Player1 - 0:1,0:0,1:2,2:3,4:3
func PlayGame(gameData []string) (GameResult, error) {
	p1, p2, err := parsePlayers(gameData)
	if err != nil {
		return GameResult{}, nil
	}
	player1Hits := p2.getOpponentHits(p1.moves)
	player2Hits := p1.getOpponentHits(p2.moves)
	result := "It is a draw"
	if player1Hits > player2Hits {
		result = "Player1 wins"
	} else if player2Hits > player1Hits {
		result = "Player2 wins"
	}
	return GameResult{
		Player1Board: p1.board.String(),
		Player2Board: p2.board.String(),
		Player1Hits:  player1Hits,
		Player2Hits:  player2Hits,
		Result:       result,
	}, nil
}
