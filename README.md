# 🌟 Kiki Power 4 🌟

Un jeu de **Puissance 4** (Power 4) développé en Go avec une interface web au style rétro, inspiré de kiki.dyskolos.fr.

## 🎮 Fonctionnalités

- **3 niveaux de difficulté** : Facile (6x7), Normal (6x9), Difficile (7x8)
- **Jeu en local** : Tour par tour entre deux joueurs
- **Interface rétro** : Style années 90 avec animations CSS
- **Système de revanche** : Rejouer immédiatement après une partie

## 🚀 Installation et lancement

### Prérequis
- **Go 1.22.2** ou version supérieure

### Étapes

1. **Cloner le projet**
   ```bash
   git clone <url-du-repo>
   cd Power4-ProjGO-Dyskolos
   ```

2. **Lancer le serveur**
   ```bash
   go run server.go
   ```

3. **Ouvrir le navigateur**
   ```
   http://localhost:8080
   ```

## 🎯 Comment jouer

1. **Configurer la partie** : Entrez les noms des deux joueurs
2. **Choisir le niveau** : Sélectionnez la difficulté (taille de grille)
3. **Jouer** : Cliquez sur les colonnes pour déposer vos pions
4. **Gagner** : Alignez 4 pions horizontalement, verticalement ou en diagonale

## 📁 Structure du projet

```
├── server.go          # Serveur HTTP et routes
├── game/
│   ├── gametpt.go     # Logique du jeu Puissance 4
│   └── user.go        # Structure des joueurs
├── template/
│   ├── gametpt.html   # Interface de jeu
│   └── win.html       # Page de victoire
├── static/
│   └── styles.css     # Styles rétro 
├── src/               # Images des joueurs
└── index.html         # Page d'accueil
```


## 🏆 Niveaux de difficulté

| Niveau | Taille | Description |
|--------|--------|-------------|
| 🟢 Facile | 6x7 | Grille classique |
| 🟡 Normal | 6x9 | Plus de colonnes |
| 🔴 Difficile | 7x8 | Grille élargie |

## 🎪 Crédits

Développé avec ❤️ par [Dyskolos_](https://dyskolos.fr) dans l'esprit rétro de [kiki.dyskolos.fr](https://kiki.dyskolos.fr)