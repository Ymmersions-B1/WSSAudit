package main

import (
	"fmt"
	"log"
	"main/Correction"
	"main/utils"
	"os"
	"strings"
)

// /!\ Le dossier Correction est un exemple de répo étudiant
// 0 . Ne pas oublier le "go mod tidy" ;)
// 1 . Clonner le répo de l'étudiant dans ce projet
// 2 . Dans le répo étudiant, suprimez le go.mod & go.sum
// 3 . Modifiez les variables wordsDir & studentSolver
// 4 . Dans le code des étudiants, modifiez le nouveau chemin d'accès au fichier words.txt
// 5 . Lancer le fichier main
// 6 . Comparez les résultats
func main() {
	//Chemin d'accès au fichier words
	wordsDir := "./Correction/words.txt"

	//Fonction solver de l'étudiant
	studentSolver := Correction.Solver

	startTests(studentSolver, wordsDir)
}

func startTests(studentSolver func([10][10]rune), wordsPath string) {
	fmt.Println("[-----------------[TEST 1]-----------------]")
	replaceFileContent(wordsPath, utils.Test1.Words)
	utils.Solver(utils.Test1.Grid, utils.Test1.Words)
	fmt.Println("[---------------[Student reult]----------------]")
	studentSolver(utils.Test1.Grid)
	fmt.Println("[---------------[FIN TEST 1]----------------]\n")

	fmt.Println("[-----------------[TEST 2]-----------------]")
	replaceFileContent(wordsPath, utils.Test2.Words)
	utils.Solver(utils.Test2.Grid, utils.Test2.Words)
	fmt.Println("[---------------[Student reult]----------------]")
	studentSolver(utils.Test2.Grid)
	fmt.Println("[---------------[FIN TEST 2]----------------]\n")

	fmt.Println("[-----------------[TEST 3]-----------------]")
	replaceFileContent(wordsPath, utils.Test3.Words)
	utils.Solver(utils.Test3.Grid, utils.Test3.Words)
	fmt.Println("[---------------[Student reult]----------------]")
	studentSolver(utils.Test3.Grid)
	fmt.Println("[---------------[FIN TEST 3]----------------]\n")
}

func replaceFileContent(wordsPath string, lines []string) {
	// Joindre les lignes en une seule chaîne avec des sauts de ligne
	content := []byte(strings.Join(lines, "\n"))

	// Écrire le contenu dans le fichier
	err := os.WriteFile(wordsPath, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
