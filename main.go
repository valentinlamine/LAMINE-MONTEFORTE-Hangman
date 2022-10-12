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
	fmt.Println("\n\n\nBonne chance, vous avez 10 essaies")
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

func Affichage_mot() {
	for _, caractère := range mot_actuel {
		fmt.Print(strings.ToUpper(string(caractère)), " ")
	}
	fmt.Print("\n\n")
}

func Affichage_liste_lettre() {
	if len(liste_lettre) == 0 {
		return
	}
	fmt.Print("Lettres déjà essayées : ")
	for _, lettre := range liste_lettre {
		fmt.Print(lettre, " ")
	}
	fmt.Println()
}

func Affichage_pendu() {
	//se déclenche lorsque l'utilisateur se trompe
	//le pendu est représenté par des caractères ASCII
	//le pendu se trouve dans le fichier hangman.txt
	//le fichier contient 10 positions, une pour chaque essaie
	//chaque position contient 7 lignes, finissant par un saut de ligne
	//chaque ligne contient 9 caractères, finissant par un saut de ligne

	//on ouvre le fichier
	fichier, err := os.ReadFile("hangman.txt")
	if err != nil {
		fmt.Println("Impossible d'ouvrir le fichier")
		os.Exit(1)
	}
	//on récupère la position du pendu
	var position int = 10 - essaie
	//on récupère la ligne du pendu
	var ligne int = 0 + position*7
	//on affiche le pendu
	for i := 0; i < 7; i++ {
		fmt.Println(string(fichier[ligne+i*9 : ligne+i*9+9]))
	}
}

func Entrée_utilisateur() string {
	var lettre string
	fmt.Print("Choix : ")
	fmt.Scanln(&lettre)
	if len(lettre) != 1 || lettre < "a" || lettre > "z" {
		fmt.Println("Merci d'entrer une lettre")
		return Entrée_utilisateur()
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
	for _, lettre_essaye := range liste_lettre {
		if strings.ToUpper(lettre) == lettre_essaye {
			essaie--
			fmt.Println("Vous avez déjà essayé cette lettre, il vous reste", essaie, "essaies")
			Affichage_pendu()
		}
	}
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
		fmt.Println("La lettre n'est pas dans le mot, il vous reste", essaie, "essaies")
		Affichage_pendu()
	} else {
		mot_actuel = mot_temp
	}
}
