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
	for essaie > 0 { //boucle principale du jeu, s'arrête lorsque l'on perd
		Affichage_mot()
		Affichage_liste_lettre()
		Revelation_lettre(Entrée_utilisateur())
		if mot_actuel == mot_a_trouver { //condition de victoire
			fmt.Println("\n\nVous avez gagné !\nLe mot était bien :", mot_a_trouver)
			os.Exit(0) //sortie du programme
		}
	}
	fmt.Println("\n\nVous avez perdu !\nLe mot était :", mot_a_trouver)
}

func Initialisation() {
	if len(os.Args) != 2 { //vérifie qu'il y a bien un argument
		fmt.Println("Merci d'indiquer le nom du fichier texte à utiliser : \ngo run main.go nom_du_fichier.txt")
		os.Exit(1)
	} else {
		Lecture_Fichier(os.Args[1])
	}
	Affichage_espace()
	fmt.Println("Bienvenue dans le jeu du pendu !")
	fmt.Println("Bonne chance, vous avez 10 essaies")
	fmt.Println("\nNote : Le programme affiche des lettres dès le lancement, Toutefois il n'affiche pas pour autant toutes les occurences de ces lettres")
}

func Affichage_espace() { //pour faire un affichage propre
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
}

func Affichage_mot() {
	fmt.Println()
	for _, caractère := range mot_actuel {
		fmt.Print(strings.ToUpper(string(caractère)), " ") // Permet d'afficher les lettres en majuscule avec un espace entre chaque
	}
	fmt.Print("\n\n")
}

func Affichage_liste_lettre() {
	if len(liste_lettre) == 0 {
		return
	}
	fmt.Print("Liste des essais : ")
	for _, lettre := range liste_lettre {
		fmt.Print(lettre, " ") //affiche la liste des lettres déjà essayées
	}
	fmt.Println()
}

func Affichage_pendu() {
	Affichage_espace()
	fichier, err := os.ReadFile("hangman.txt") //on lit le fichier
	if err != nil {                            //si il y a une erreur
		fmt.Println("Impossible d'ouvrir le fichier")
		os.Exit(1) //on quitte le programme
	}
	var position int = 10 - (essaie + 1) //on calcule la position de la ligne à afficher
	for i := 0; i < 7; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(string(fichier[position*71+i*10+j])) //on affiche la ligne
		}
	}
	fmt.Println()
}

func Entrée_utilisateur() string {
	var lettre string
	fmt.Print("Choix : ")
	fmt.Scanln(&lettre)      //on récupère la lettre entrée par l'utilisateur
	if !Est_lettre(lettre) { //vérifie que l'utilisateur a bien entré que des lettres
		fmt.Println("Merci de n'entrer que des lettres minusucules")
		return Entrée_utilisateur() //on relance la fonction
	}
	for _, lettre_essaye := range liste_lettre { //vérifie que l'utilisateur n'a pas déjà essayé cette lettre
		if strings.ToUpper(lettre) == lettre_essaye {
			fmt.Println("Vous avez déjà essayé cette lettre, merci d'en choisir une autre")
			return Entrée_utilisateur() //on relance la fonction
		}
	}
	return strings.ToLower(lettre) //on retourne la lettre en minuscule
}

func Lecture_Fichier(nom_fichier string) {
	var mot string
	var liste_mots []string
	fichier, err := os.ReadFile(nom_fichier) //on lit le fichier
	if err != nil {                          //si il y a une erreur
		fmt.Println("Impossible d'ouvrir le fichier")
		os.Exit(1) //on quitte le programme
	}
	for index, caractère := range fichier {
		if caractère == 10 { //lorsque l'on va à la ligne le mot est fini
			liste_mots = append(liste_mots, mot) //on l'ajoute à la liste
			mot = ""                             //on rénitialise mot
		} else { //sinon on ajoute le caractère a mot
			mot += string(caractère)
		}
		if index == len(fichier)-1 { //on vérifie la fin
			liste_mots = append(liste_mots, mot)
		}
	}
	rand.Seed(int64(os.Getpid()))                          //on initialise le générateur de nombre aléatoire
	mot_a_trouver = liste_mots[rand.Intn(len(liste_mots))] //on choisit un mot aléatoire
	for i := 0; i < len(mot_a_trouver); i++ {
		mot_actuel += "_" //on initialise le mot actuel avec des _
	}
	for i := 0; i < len(mot_a_trouver)/2-1; i++ { //on remplace des _ par des lettres
		mot_actuel = strings.Replace(mot_actuel, "_", string(mot_a_trouver[i]), 1)
	}
}

func Revelation_lettre(lettre string) {
	if len(lettre) != 1 { //vérifie que l'utilisateur a rentré une lettre ou un mot
		if lettre == mot_a_trouver { //vérifie que le mot entré est le bon
			mot_actuel = mot_a_trouver //on met le mot actuel à jour
		} else {
			essaie -= 2       //on enlève 2 essaies
			Affichage_pendu() //on affiche le pendu
			fmt.Println("Votre mot est incorrect, il vous reste", essaie, "essaies")
		}
	} else {
		var mot_temporaire string
		for index, caractère := range mot_a_trouver { //on parcourt le mot à trouver
			if string(caractère) == lettre { //si la lettre est dans le mot
				mot_temporaire += lettre //on ajoute la lettre au mot temporaire
			} else {
				mot_temporaire += string(mot_actuel[index]) //sinon on ajoute le caractère du mot actuel
			}
		}
		liste_lettre = append(liste_lettre, strings.ToUpper(lettre)) //on ajoute la lettre à la liste des lettres essayées
		sort.Strings(liste_lettre)                                   //on trie la liste
		if mot_temporaire == mot_actuel {                            //si le mot temporaire est égal au mot actuel
			essaie-- //on enlève un essaie
			Affichage_pendu()
			fmt.Println("La lettre n'est pas dans le mot, il vous reste", essaie, "essaies :")
		} else {
			mot_actuel = mot_temporaire //on met le mot actuel à jour
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

func Est_lettre(str string) bool { //vérifie que la chaine de caractère ne contient que des lettres
	for _, lettre := range str {
		if lettre < 'a' || lettre > 'z' {
			return false
		}
	}
	return true
}
