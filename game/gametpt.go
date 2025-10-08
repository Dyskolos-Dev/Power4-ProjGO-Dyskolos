package game

import "fmt"

const (
	ROWS = 6
	COLS = 7
)

// GameState représente l'état full d'une game de P4
type GameState struct {
	Board       [][]int `json:"board"` // 0 = vide, 1 = joueur 1 (rouge), 2 = joueur 2 (jaune)
	Player1     User    `json:"player1"`
	Player2     User    `json:"player2"`
	CurrentTurn string  `json:"currentTurn"`
	GameMode    string  `json:"gameMode"`
	Winner      string  `json:"winner"` // "" si pas de gagnant, nom du gagnant sinon
	GameOver    bool    `json:"gameOver"`
	LastMove    struct {
		Row int `json:"row"`
		Col int `json:"col"`
	} `json:"lastMove"`
}

// NewGameState crée un new plato
func NewGameState(player1Name, player2Name string) *GameState {
	board := make([][]int, ROWS)
	for i := range board {
		board[i] = make([]int, COLS)
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
		Winner:      "",
		GameOver:    false,
	}
}

// DropPiece fait tomber un pion dans la cln choisi
func (gs *GameState) DropPiece(col int) bool {
	if col < 0 || col >= COLS || gs.GameOver {
		return false
	}

	// Trouver la 1e case libr dans la col (en partant du bas)
	for row := ROWS - 1; row >= 0; row-- {
		if gs.Board[row][col] == 0 {
			// Déterminer le no du playr
			playerNum := 1
			if gs.CurrentTurn == gs.Player2.Name {
				playerNum = 2
			}

			gs.Board[row][col] = playerNum
			gs.LastMove.Row = row
			gs.LastMove.Col = col

			// Vérifier si victoir
			if gs.CheckWin(row, col, playerNum) {
				gs.Winner = gs.CurrentTurn
				gs.GameOver = true
			} else if gs.IsBoardFull() {
				gs.Winner = "Match nul"
				gs.GameOver = true
			} else {
				// Changer de tr
				if gs.CurrentTurn == gs.Player1.Name {
					gs.CurrentTurn = gs.Player2.Name
				} else {
					gs.CurrentTurn = gs.Player1.Name
				}
			}
			return true
		}
	}
	return false // Colonne full
}

func (gs *GameState) CheckWin(row, col, player int) bool {
	directions := [][]int{
		{0, 1},  // horizontal
		{1, 0},  // vertical
		{1, 1},  // diagonal /
		{1, -1}, // diagonal \
	}

	for _, dir := range directions {
		count := 1

		for i := 1; i < 4; i++ {
			newRow := row + dir[0]*i
			newCol := col + dir[1]*i
			if newRow >= 0 && newRow < ROWS && newCol >= 0 && newCol < COLS &&
				gs.Board[newRow][newCol] == player {
				count++
			} else {
				break
			}
		}

		for i := 1; i < 4; i++ {
			newRow := row - dir[0]*i
			newCol := col - dir[1]*i
			if newRow >= 0 && newRow < ROWS && newCol >= 0 && newCol < COLS &&
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
	for col := 0; col < COLS; col++ {
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

func (gm *GameManager) CreateGame(player1Name, player2Name, gameMode string) string {
	gameState := NewGameState(player1Name, player2Name)
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
