#BattleShips

* Battleship is a game played between 2 players. Each player will be initialised with an M * M
grid with S number of the ships placed at specified positions on the grid.

* One battleship occupies a single position on the Grid.

* Objective of the game is to destroy opponent's ships. Each player will
be given T number of missiles.

* Based on hits/misses of a missile on the opponent, either of the player might be vistorious or the game might end as a draw.


##Installation
```
go get github.com/vedhavyas/battleship/cmd
cd $GOPATH/bin
./battleship --help
```

##Input
Program requires an `Input file` and an `Output file`.

###Input File
Input file should be of the following format

1. GridSize (0<M<10)
2. Number of ships for each player
3. Player 1 ship positions, should be of format x1,y1:x2,y2:x3,y3:x4,y4
4. Player 2 ship positions, should be of format x1,y1:x2,y2:x3,y3:x4,y4
5. Total number of missiles for each player (0<T<100)
6. Player 1 Moves, should be of format x1,y1:x2,y2:x3,y3:x4,y4
7. Player 2 Moves, should be of format x1,y1:x2,y2:x3,y3:x4,y4

Since it was not specifically mentioned in the problem as to how the coordinate system of the
board works. Hence, we are assuming that (0,0) lies in the top left corner and (m,m) in bottom right.
Final depiction of the board looks as follows

```
     (0,0) (0,1) (0,2)
(1,0)  _     _     _
(2,0)  _     _     _
```

###Output file
Program requires an output file to write the result back to the file.
If file does not exist, one will be created.
If file exists, then the data is overwritten

##Output
Programs sample output in the `Output File` will look as follows.
```
Player1
O O _ _ _
_ B O _ _
B _ _ X _
_ _ _ _ B
_ _ _ X _
```
```
Player2
_ X _ _ _
_ _ _ _ _
_ _ _ X _
B O _ _ B
_ X _ O _
```

```
P1: 3
P2: 2
Player1 wins
```

* O - represents a missed hit by the opponent
* B - represents a safe/un-hit ship of the player
* X - represents a confirmed hit by the opponent
* _ - represnets an empty location