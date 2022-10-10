package main

import (
	"fmt"
	"os"
)

//

var liste_mots []string
var essaie int = 10

func main() {
	Initialisation()
	fmt.Println(liste_mots)
}

func Initialisation() {
	if len(os.Args) != 2 {
		fmt.Println("Merci d'indiquer le nom du fichier texte à utiliser : \ngo run main.go nom_du_fichier.txt")
		os.Exit(1)
	} else {
		Lecture_Fichier(os.Args[1])
	}
}

func Lecture_Fichier(nom_fichier string) {
	var mot string
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
}
