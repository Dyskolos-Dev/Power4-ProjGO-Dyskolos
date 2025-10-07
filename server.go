package main

import (
	"html/template"
	"log"
	"net/http"
)

// User représente un utilisateur du jeu
type User struct {
	Name  string `json:"name"`
	Color string `json:"color"` // "red" ou "yellow"
}

// GameSession représente une session de jeu
type GameSession struct {
	Player1     User   `json:"player1"`
	Player2     User   `json:"player2"`
	CurrentTurn string `json:"currentTurn"` // nom du joueur dont c'est le tour
	GameMode    string `json:"gameMode"`    // "local" pour tour par tour local
}

// NewGameSession crée une nouvelle session de jeu
func NewGameSession(player1Name, player2Name string) *GameSession {
	return &GameSession{
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
	}
}

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

	// Récupérer les données du formulaire
	player1Name := r.FormValue("player1")
	player2Name := r.FormValue("player2")
	gameMode := r.FormValue("gamemode")

	// Validation des données
	if player1Name == "" || player2Name == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Créer une nouvelle session de jeu
	gameSession := NewGameSession(player1Name, player2Name)
	gameSession.GameMode = gameMode

	// Rediriger vers la page de jeu avec les données
	tmpl, err := template.ParseFiles("template/gametpt.html")
	if err != nil {
		log.Printf("Erreur lors du chargement du template: %v", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, gameSession)
	if err != nil {
		log.Printf("Erreur lors de l'exécution du template: %v", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/start-game", StartGame)

	// Servir les fichiers statiques
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Servir les images depuis le dossier /src
	imgFS := http.FileServer(http.Dir("src/"))
	http.Handle("/src/", http.StripPrefix("/src/", imgFS))

	log.Println("Server démarré sur le port :8080")
	log.Println("Routes disponibles:")
	log.Println("  GET  / - Page d'accueil")
	log.Println("  POST /start-game - Démarrer une partie")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
