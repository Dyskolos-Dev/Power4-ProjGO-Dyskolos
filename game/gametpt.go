package game

import "fmt"

type GameState struct {
	Board       [][]int `json:"board"`
	Player1     User    `json:"player1"`
	Player2     User    `json:"player2"`
	CurrentTurn string  `json:"currentTurn"`
	GameMode    string  `json:"gameMode"`
	Level       string  `json:"level"`
	Rows        int     `json:"rows"`
	Cols        int     `json:"cols"`
	Winner      string  `json:"winner"`
	GameOver    bool    `json:"gameOver"`
	LastMove    struct {
		Row int `json:"row"`
		Col int `json:"col"`
	} `json:"lastMove"`
}

func NewGameState(player1Name, player2Name, level string) *GameState {
	var rows, cols int
	switch level {
	case "facile":
		rows, cols = 6, 7
	case "normal":
		rows, cols = 6, 9
	case "difficile":
		rows, cols = 7, 8
	default:
		rows, cols = 6, 7
	}

	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, cols)
	}

	return &GameState{
		Board: board,
		Player1: User{
			Name:  player1Name,
			Color: "red",
		},
		Player2: User{
			Name:  player2Name,
			Color: "yellow",
		},
		CurrentTurn: player1Name,
		GameMode:    "local",
		Level:       level,
		Rows:        rows,
		Cols:        cols,
		Winner:      "",
		GameOver:    false,
	}
}

func (gs *GameState) DropPiece(col int) bool {
	if col < 0 || col >= gs.Cols || gs.GameOver {
		return false
	}

	for row := gs.Rows - 1; row >= 0; row-- {
		if gs.Board[row][col] == 0 {
			playerNum := 1
			if gs.CurrentTurn == gs.Player2.Name {
				playerNum = 2
			}

			gs.Board[row][col] = playerNum
			gs.LastMove.Row = row
			gs.LastMove.Col = col

			if gs.CheckWin(row, col, playerNum) {
				gs.Winner = gs.CurrentTurn
				gs.GameOver = true
			} else if gs.IsBoardFull() {
				gs.Winner = "Match nul"
				gs.GameOver = true
			} else {
				if gs.CurrentTurn == gs.Player1.Name {
					gs.CurrentTurn = gs.Player2.Name
				} else {
					gs.CurrentTurn = gs.Player1.Name
				}
			}
			return true
		}
	}
	return false
}

func (gs *GameState) CheckWin(row, col, player int) bool {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
	}

	for _, dir := range directions {
		count := 1

		for i := 1; i < 4; i++ {
			newRow := row + dir[0]*i
			newCol := col + dir[1]*i
			if newRow >= 0 && newRow < gs.Rows && newCol >= 0 && newCol < gs.Cols &&
				gs.Board[newRow][newCol] == player {
				count++
			} else {
				break
			}
		}

		for i := 1; i < 4; i++ {
			newRow := row - dir[0]*i
			newCol := col - dir[1]*i
			if newRow >= 0 && newRow < gs.Rows && newCol >= 0 && newCol < gs.Cols &&
				gs.Board[newRow][newCol] == player {
				count++
			} else {
				break
			}
		}

		if count >= 4 {
			return true
		}
	}
	return false
}

func (gs *GameState) IsBoardFull() bool {
	for col := 0; col < gs.Cols; col++ {
		if gs.Board[0][col] == 0 {
			return false
		}
	}
	return true
}

type GameManager struct {
	sessions map[string]*GameState
}

func NewGameManager() *GameManager {
	return &GameManager{
		sessions: make(map[string]*GameState),
	}
}

func (gm *GameManager) CreateGame(player1Name, player2Name, gameMode, level string) string {
	gameState := NewGameState(player1Name, player2Name, level)
	gameState.GameMode = gameMode

	sessionID := generateSessionID(player1Name, player2Name, len(gm.sessions))
	gm.sessions[sessionID] = gameState

	return sessionID
}

func (gm *GameManager) GetGame(sessionID string) (*GameState, bool) {
	game, exists := gm.sessions[sessionID]
	return game, exists
}

func (gm *GameManager) MakeMove(sessionID string, col int) (bool, error) {
	game, exists := gm.sessions[sessionID]
	if !exists {
		return false, NewGameError("Session de jeu introuvable")
	}

	success := game.DropPiece(col)
	if !success {
		return false, NewGameError("Coup invalide")
	}

	return true, nil
}

func generateSessionID(player1, player2 string, sessionCount int) string {
	return player1 + "-" + player2 + "-" + fmt.Sprintf("%d", sessionCount)
}

type GameError struct {
	Message string
}

func NewGameError(message string) *GameError {
	return &GameError{Message: message}
}

func (ge *GameError) Error() string {
	return ge.Message
}
