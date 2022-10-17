package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
)

var mot_a_trouver string
var mot_actuel string
var essaie int = 10
var liste_lettre []string

func main() {
	Initialisation()
	Affichage_espace()
	fmt.Println("Bienvenue dans le jeu du pendu !")
	fmt.Println("Bonne chance, vous avez 10 essaies")
	fmt.Println("\nNote : Le programme affiche des lettres dès le lancement, Toutefois il n'affiche pas pour autant toutes les occurences de ces lettres")
	for essaie > 0 {
		Affichage_mot()
		Affichage_liste_lettre()
		var lettre string = Entrée_utilisateur()
		Revelation_lettre(lettre)
		if mot_actuel == mot_a_trouver {
			fmt.Println("\n\nVous avez gagné !")
			fmt.Println("Le mot était bien :", mot_a_trouver)
			os.Exit(0)
		}
	}
	fmt.Println("\n\nVous avez perdu !")
	fmt.Println("Le mot était :", mot_a_trouver)
}

func Initialisation() {
	if len(os.Args) != 2 {
		fmt.Println("Merci d'indiquer le nom du fichier texte à utiliser : \ngo run main.go nom_du_fichier.txt")
		os.Exit(1)
	} else {
		Lecture_Fichier(os.Args[1])
	}
}

func Affichage_espace() {
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
}

func Affichage_mot() {
	fmt.Print("\n")
	for _, caractère := range mot_actuel {
		fmt.Print(strings.ToUpper(string(caractère)), " ")
	}
	fmt.Print("\n\n")
}

func Affichage_liste_lettre() {
	if len(liste_lettre) == 0 {
		return
	}
	fmt.Print("Liste des essais : ")
	for _, lettre := range liste_lettre {
		fmt.Print(lettre, " ")
	}
	fmt.Println()
}

func Affichage_pendu() {
	Affichage_espace()
	fichier, err := os.ReadFile("hangman.txt")
	if err != nil {
		fmt.Println("Impossible d'ouvrir le fichier")
		os.Exit(1)
	}
	var position int = 10 - (essaie + 1)
	for i := 0; i < 7; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(string(fichier[position*71+i*10+j]))
		}
	}
	fmt.Println()
}

func Entrée_utilisateur() string {
	var lettre string
	fmt.Print("Choix : ")
	fmt.Scanln(&lettre)
	if !Is_alpha(lettre) {
		fmt.Println("Merci de n'entrer que des lettres minusucules")
		return Entrée_utilisateur()
	}
	for _, lettre_essaye := range liste_lettre {
		if strings.ToUpper(lettre) == lettre_essaye {
			fmt.Println("Vous avez déjà essayé cette lettre, merci d'en choisir une autre")
			return Entrée_utilisateur()
		}
	}
	return strings.ToLower(lettre)
}

func Lecture_Fichier(nom_fichier string) {
	var mot string
	var liste_mots []string
	fichier, err := os.ReadFile(nom_fichier)
	if err != nil {
		fmt.Println("Impossible d'ouvrir le fichier")
		os.Exit(1)
	}
	for index, caractère := range fichier {
		if caractère == 10 { //lorsque l'on va à la ligne le mot est fini
			liste_mots = append(liste_mots, mot) //on l'append
			mot = ""                             //on rénitialise mot
		} else { //sinon on ajoute le caractère a mot
			mot += string(caractère)
		}
		if index == len(fichier)-1 { //on vérifie la fin
			liste_mots = append(liste_mots, mot)
		}
	}
	rand.Seed(int64(os.Getpid()))
	mot_a_trouver = liste_mots[rand.Intn(len(liste_mots))]
	for i := 0; i < len(mot_a_trouver); i++ {
		mot_actuel += "_"
	}
	for i := 0; i < len(mot_a_trouver)/2-1; i++ {
		mot_actuel = strings.Replace(mot_actuel, "_", string(mot_a_trouver[i]), 1)
	}
}

func Revelation_lettre(lettre string) {
	if len(lettre) != 1 {
		if lettre == mot_a_trouver {
			mot_actuel = mot_a_trouver
		} else {
			essaie -= 2
			Affichage_pendu()
			fmt.Println("Votre mot est incorrect, il vous reste", essaie, "essaies")
		}
	} else {
		var mot_temp string
		for index, caractère := range mot_a_trouver {
			if string(caractère) == lettre {
				mot_temp += lettre
			} else {
				mot_temp += string(mot_actuel[index])
			}
		}
		liste_lettre = append(liste_lettre, strings.ToUpper(lettre))
		sort.Strings(liste_lettre)
		if mot_temp == mot_actuel {
			essaie--
			Affichage_pendu()
			fmt.Println("La lettre n'est pas dans le mot, il vous reste", essaie, "essaies :")
		} else {
			mot_actuel = mot_temp
			if essaie != 10 {
				Affichage_pendu()
				fmt.Println("La lettre est dans le mot, il vous reste", essaie, "essaies :")
			} else {
				Affichage_espace()
				fmt.Println("La lettre est dans le mot, il vous reste", essaie, "essaies :")
			}
		}
	}
}

func Is_alpha(str string) bool {
	for _, letter := range str {
		if letter < 'a' || letter > 'z' {
			return false
		}
	}
	return true
}
