package main

import (
	"html/template"
	"log"
	"net/http"
	"power4/game"
	"strconv"
)

var gameManager = game.NewGameManager()

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func StartGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	player1Name := r.FormValue("player1")
	player2Name := r.FormValue("player2")
	gameMode := r.FormValue("gamemode")
	level := r.FormValue("level")
	if player1Name == "" || player2Name == "" || level == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	sessionID := gameManager.CreateGame(player1Name, player2Name, gameMode, level)

	http.Redirect(w, r, "/game/"+sessionID, http.StatusSeeOther)
}

func Rematch(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	oldSessionID := r.FormValue("sessionId")
	gameState, exists := gameManager.GetGame(oldSessionID)
	if !exists {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	newSessionID := gameManager.CreateGame(gameState.Player1.Name, gameState.Player2.Name, gameState.GameMode, gameState.Level)

	http.Redirect(w, r, "/game/"+newSessionID, http.StatusSeeOther)
}

func ShowGame(w http.ResponseWriter, r *http.Request) {
	sessionID := r.URL.Path[len("/game/"):]

	gameState, exists := gameManager.GetGame(sessionID)
	if !exists {
		http.NotFound(w, r)
		return
	}

	if gameState.GameOver {
		tmpl, err := template.ParseFiles("template/win.html")
		if err != nil {
			log.Printf("Erreur lors du chargement du template win: %v", err)
			http.Error(w, "Erreur interne", http.StatusInternalServerError)
			return
		}
		data := struct {
			*game.GameState
			SessionID string
		}{
			GameState: gameState,
			SessionID: sessionID,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Erreur lors de l'exécution du template win: %v", err)
			http.Error(w, "Erreur interne", http.StatusInternalServerError)
		}
		return
	}
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"subtract": func(a, b int) int {
			return a - b
		},
		"mod": func(a, b int) int {
			return a % b
		},
		"list": func(args ...int) []int {
			return args
		},
		"makeRange": func(n int) []int {
			arr := make([]int, n)
			for i := 0; i < n; i++ {
				arr[i] = i
			}
			return arr
		},
	}
	tmpl, err := template.New("gametpt.html").Funcs(funcMap).ParseFiles("template/gametpt.html")
	if err != nil {
		log.Printf("Erreur lors du chargement du template: %v", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}
	data := struct {
		*game.GameState
		SessionID string
	}{
		GameState: gameState,
		SessionID: sessionID,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Erreur lors de l'exécution du template: %v", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
	}
}

func MakeMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
	sessionID := r.FormValue("sessionId")
	colStr := r.FormValue("column")

	col, err := strconv.Atoi(colStr)
	if err != nil {
		http.Error(w, "Colonne invalide", http.StatusBadRequest)
		return
	}
	success, gameErr := gameManager.MakeMove(sessionID, col)
	if !success {
		http.Error(w, gameErr.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/game/"+sessionID, http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/start-game", StartGame)
	http.HandleFunc("/rematch", Rematch)
	http.HandleFunc("/game/", ShowGame)
	http.HandleFunc("/move", MakeMove)
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	imgFS := http.FileServer(http.Dir("src/"))
	http.Handle("/src/", http.StripPrefix("/src/", imgFS))
	log.Println("Server démarré sur le port :8080")
	log.Println("Routes disponibles:")
	log.Println("  GET  / - Page d'accueil")
	log.Println("  POST /start-game - Démarrer une partie")
	log.Println("  POST /rematch - Revanche")
	log.Println("  GET  /game/{id} - Page de jeu")
	log.Println("  POST /move - Jouer un coup")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
