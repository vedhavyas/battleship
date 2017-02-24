package battleship

import "testing"

func TestParsePlayers(t *testing.T) {

	tests := []struct {
		gameData    []string
		p1          player
		p2          player
		shouldPanic bool
	}{
		{
			gameData: []string{
				"5",
				"5",
				"1:1,2:0,2:3,3:4,4:3",
				"0:1,2:3,3:0,3:4,4:1",
				"5",
				"0:1,4:3,2:3,3:1,4:1",
				"0:1,0:0,1:2,2:3,4:3",
			},

			p1: player{
				totalMissiles: 5,
				board: playerBoard{
					gridSize:   5,
					totalShips: 5,
				},
			},
			p2: player{
				totalMissiles: 5,
				board: playerBoard{
					gridSize:   5,
					totalShips: 5,
				},
			},
		},
		{
			gameData: []string{
				"6",
				"5",
				"1:1,2:0,2:3,3:4,4:3",
				"0:1,2:3,3:0,3:4,4:1",
				"5",
				"0:1,4:3,2:3,3:1,4:1",
				"0:1,0:0,1:2,2:3,4:3",
			},

			p1: player{
				totalMissiles: 5,
				board: playerBoard{
					gridSize:   6,
					totalShips: 5,
				},
			},
			p2: player{
				totalMissiles: 5,
				board: playerBoard{
					gridSize:   6,
					totalShips: 5,
				},
			},
		},
		{
			gameData: []string{
				"7",
				"4",
				"1:1,2:0,2:3,3:4",
				"0:1,2:3,3:0,3:4",
				"5",
				"0:1,4:3,2:3,3:1,4:1",
				"0:1,0:0,1:2,2:3,4:3",
			},

			p1: player{
				totalMissiles: 5,
				board: playerBoard{
					gridSize:   7,
					totalShips: 4,
				},
			},
			p2: player{
				totalMissiles: 5,
				board: playerBoard{
					gridSize:   7,
					totalShips: 4,
				},
			},
		},
		{
			gameData: []string{
				"7",
				"3",
				"1:1,2:0,2:3,3:4",
				"0:1,2:3,3:0,3:4",
				"5",
				"0:1,4:3,2:3,3:1,4:1",
				"0:1,0:0,1:2,2:3,4:3",
			},
			shouldPanic: true,
		},
	}

	for _, testCase := range tests {
		eP1, eP2, err := parsePlayers(testCase.gameData)
		if err != nil {
			if !testCase.shouldPanic {
				t.Error("Parse should not have Failed")
			}

			return
		}

		if testCase.p1.totalMissiles != eP1.totalMissiles {
			t.Error("P1 Missiles did not match")
		}

		if testCase.p1.board.gridSize != eP1.board.gridSize {
			t.Error("P1 GridSize did not match")
		}

		if testCase.p2.totalMissiles != eP2.totalMissiles {
			t.Error("P2 Missiles did not match")
		}

		if testCase.p2.board.gridSize != eP2.board.gridSize {
			t.Error("P2 GridSize did not match")
		}
	}

}

func TestPlayGame(t *testing.T) {
	tests := []struct {
		gameData    []string
		gameResult  GameResult
		shouldPanic bool
	}{
		{
			gameData: []string{
				"5",
				"5",
				"1:1,2:0,2:3,3:4,4:3",
				"0:1,2:3,3:0,3:4,4:1",
				"5",
				"0:1,4:3,2:3,3:1,4:1",
				"0:1,0:0,1:2,2:3,4:3",
			},
			gameResult: GameResult{
				Player1Hits: 3,
				Player2Hits: 2,
				Result:      "Player1 wins",
			},
		},
		{
			gameData: []string{
				"6",
				"5",
				"0:1,2:3,3:0,3:4,4:1",
				"1:1,2:0,2:3,3:4,4:3",
				"5",
				"0:1,0:0,1:2,2:3,4:3",
				"0:1,4:3,2:3,3:1,4:1",
			},

			gameResult: GameResult{
				Player1Hits: 2,
				Player2Hits: 3,
				Result:      "Player2 wins",
			},
		},
		{
			gameData: []string{
				"7",
				"5",
				"1:1,2:0,2:3,3:4,4:3",
				"0:1,2:3,3:0,3:4,4:1",
				"5",
				"0:1,4:3,2:3,3:1,4:1",
				"0:1,0:0,1:1,2:3,4:3",
			},

			gameResult: GameResult{
				Player1Hits: 3,
				Player2Hits: 3,
				Result:      "It is a draw",
			},
		},
		{
			gameData: []string{
				"7",
				"3",
				"1:1,2:0,2:3,3:4",
				"0:1,2:3,3:0,3:4",
				"5",
				"0:1,4:3,2:3,3:1,4:1",
				"0:1,0:0,1:2,2:3,4:3",
			},
			shouldPanic: true,
		},
	}

	for _, testCase := range tests {
		gameResult, err := PlayGame(testCase.gameData)
		if err != nil {
			if !testCase.shouldPanic {
				t.Error("Parse Should not have Failed")
			}

			return
		}
		if gameResult.Player1Hits != testCase.gameResult.Player1Hits {
			t.Error("Player1 hits did not match")
		}

		if gameResult.Player2Hits != testCase.gameResult.Player2Hits {
			t.Error("Player2 hits did not match")
		}

		if gameResult.Result != testCase.gameResult.Result {
			t.Error("Game Result did not match")
		}
	}

}

func TestNewLocation(t *testing.T) {
	tests := []struct {
		locData  string
		loc      location
		hasError bool
	}{
		{
			locData: "1:3",
			loc: location{
				x: 1,
				y: 3,
			},
		},
		{
			locData: "2:3",
			loc: location{
				x: 2,
				y: 3,
			},
		},
		{
			locData: "8:5",
			loc: location{
				x: 8,
				y: 5,
			},
		},
		{
			locData: "10:13",
			loc: location{
				x: 10,
				y: 13,
			},
		},
		{
			locData:  "1:a",
			hasError: true,
		},
		{
			locData:  "1",
			hasError: true,
		},
	}

	for _, testCase := range tests {
		loc, err := newLocation(testCase.locData)
		if err != nil {
			if !testCase.hasError {
				t.Error("Should not have raised an error")
			}

			return
		}

		if loc.x != testCase.loc.x {
			t.Error("X coordinate did not match")
		}

		if loc.y != testCase.loc.y {
			t.Error("Y coordinate did not match")
		}
	}
}

func TestParseLocations(t *testing.T) {
	tests := []struct {
		locData  string
		locs     int
		hasError bool
	}{
		{
			locData: "1:3,2:3,1:4,1:5,10:24",
			locs:    5,
		},
		{
			locData: "5:3,4:2,5:1,1:5",
			locs:    4,
		},
		{
			locData: "1:3,2:3,1:4,1:5,10:24,5:1,1:5",
			locs:    7,
		},
		{
			locData: "10:24",
			locs:    1,
		},
		{
			locData: "1:3,10:24",
			locs:    2,
		},
		{
			locData:  "13,2:3,1:4,1:5,10:24",
			hasError: true,
		},
	}

	for _, testCase := range tests {
		locs, err := parseLocations(testCase.locData)
		if err != nil {
			if !testCase.hasError {
				t.Error("Should not have failed")
			}

			return
		}

		if len(locs) != testCase.locs {
			t.Error("Parsed location number did not match")
		}
	}
}
