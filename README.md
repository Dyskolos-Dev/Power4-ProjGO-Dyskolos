# 🌟 Kiki Power 4 🌟

Un jeu de **Puissance 4** (Power 4) développé en Go avec une interface web au style rétro, inspiré de kiki.dyskolos.fr.

🎮 **Jeu en ligne** : [p4.dyskolos.fr](https://p4.dyskolos.fr)

## 🎮 Fonctionnalités

- **3 niveaux de difficulté** : Facile (6x7), Normal (6x9), Difficile (7x8)
- **2 modes de jeu** :
  - **Mode Classique** : Gravité normale du haut vers le bas
  - **Mode Reverse (Gravité Inversée)** : La gravité change tous les 5 tours ! Les pions tombent alternativement vers le bas puis vers le haut
- **Jeu en local** : Tour par tour entre deux joueurs
- **Interface rétro** : Style années 90 avec animations CSS
- **Indicateurs visuels** : Changement de couleur du plateau selon la gravité (bleu → cyan pour normale, violet → magenta pour inversée)
- **Système de revanche** : Rejouer immédiatement après une partie

## 🚀 Installation et lancement

### Prérequis
- **Go 1.22.2** ou version supérieure

### Étapes

1. **Cloner le projet**
   ```bash
   git clone https://github.com/Dyskolos-Dev/Power4-ProjGO-Dyskolos.git
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
3. **Choisir le mode** : Mode classique ou mode reverse (gravité inversée)
4. **Jouer** : Cliquez sur les colonnes pour déposer vos pions
5. **Gagner** : Alignez 4 pions horizontalement, verticalement ou en diagonale

### Mode Reverse (Gravité Inversée) 🔄
Dans ce mode bonus, la gravité du jeu change toutes les 5 tours, créant une expérience de jeu totalement unique :

- **Gravité normale** (🔽) : Les pions tombent du haut vers le bas
  - Plateau en dégradé bleu → cyan
  - Boutons avec flèche vers le bas
- **Gravité inversée** (🔼) : Les pions "remontent" du bas vers le haut
  - Plateau en dégradé violet → magenta
  - Boutons avec flèche vers le haut
- Un indicateur visuel affiche le tour actuel et le nombre de tours avant le prochain changement de gravité
- Stratégie : Anticipez les changements de gravité pour bloquer votre adversaire !




## 🏆 Niveaux de difficulté

| Niveau | Taille | Description |
|--------|--------|-------------|
| 🟢 Facile | 6x7 | Grille classique |
| 🟡 Normal | 6x9 | Plus de colonnes |
| 🔴 Difficile | 7x8 | Grille élargie |

## 🎪 Crédits

Développé avec ❤️ par [Dyskolos_](https://dyskolos.fr) dans l'esprit rétro de [kiki.dyskolos.fr](https://kiki.dyskolos.fr)