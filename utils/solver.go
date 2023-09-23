package utils

import (
	"bufio"
	"log"
	"os"

	"github.com/01-edu/z01"
)

func Solver(field [10][10]rune, words []string) {
	// On print la matrice
	printArrays(field, 1, 1)
	printSentences("Mots à trouver =", 0, 1)

	//Création des probas / cols
	var cols [][]rune
	for i := 0; i < len(field[:]); i++ {
		var temp []rune
		for _, v := range field {
			temp = append(temp, v[i])
		}
		cols = append(cols, temp)
	}

	//Création des probas / diag1
	var diag1 [][]rune
	for z := len(field[:]) - 1; z >= 0; z-- {
		for i := 0; i < 10-z; i++ {
			for j := i; j < 10-z; j++ {
				var tab []rune
				for k := i; k <= j; k++ {
					tab = append(tab, field[k][z+k])
				}
				diag1 = append(diag1, tab)
			}
		}
	}

	//Création des probas / diag2
	var diag2 [][]rune
	for z := len(field[:]) - 1; z >= 0; z-- {
		for i := 0; i <= z; i++ {
			for j := i; j <= z; j++ {
				var tab []rune
				for k := i; k <= j; k++ {
					tab = append(tab, field[z-k][k])
				}
				diag2 = append(diag2, tab)
			}
		}
	}

	//Différents tableaux avec les mots trouvés
	var findInRows [][]rune
	var findInCols [][]rune
	var findInDiag1 [][]rune
	var findInDiag2 [][]rune

	//On test pour chaques mots si il est dans un des tableau de probas
	for _, word := range words {
		var notFinded bool = true
		for _, rows := range field {
			if searchInArray(rows[:], word) {
				notFinded = false
				findInRows = append(findInRows, []rune(word))
			}
		}
		if notFinded {
			for _, col := range cols {
				if searchInArray(col[:], word) {
					notFinded = false
					findInCols = append(findInCols, []rune(word))

				}
			}
		}
		if notFinded {
			for _, diag := range diag1 {
				if searchInArray(diag[:], word) && notFinded {
					notFinded = false
					findInDiag1 = append(findInDiag1, []rune(word))
				}
			}
		}
		if notFinded {
			for _, diagR := range diag2 {
				if searchInArray(diagR[:], word) && notFinded {
					notFinded = false
					findInDiag2 = append(findInDiag2, []rune(word))
				}
			}
		}
	}
	// On print les tableaux de mots trouvés
	printFinded(findInRows)
	printFinded(findInCols)
	printFinded(findInDiag1)
	printFinded(findInDiag2)
	printSentences("\n", 0, 0)
}

// On recherche le word dans un tableau de proba
func searchInArray(line []rune, word string) bool {
	for i := 0; i < len(line); i++ {
		var temp string
		wordlen := len([]rune(word))
		for e := 0; e < wordlen; e++ {
			if i+e < len(line) {
				temp += string(line[i+e])
			}
		}
		if temp == word {
			return true
		}
	}
	return false
}

// On récupéres tout les mots (ligne) du fichier txt et return un tableau de string
func searchWordsInFiles(file string) []string {
	var words []string
	readFile, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}

	readFile.Close()

	return words
}

// On print les differentes ligne de la matrice formaté
func printArrays(t [10][10]rune, spacesBefor int, spacesAfter int) {
	for i := 0; i < spacesBefor; i++ {
		z01.PrintRune('\n')
	}

	for i := 0; i < len(t); i++ {
		for e := 0; e < len(t[i][:]); e++ {
			if e == 0 {
				z01.PrintRune('[')
			}

			z01.PrintRune(t[i][e])

			if e != len(t[i][:])-1 {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune(']')
			}
		}
		z01.PrintRune('\n')
	}
	for i := 0; i < spacesAfter; i++ {
		z01.PrintRune('\n')
	}
}

// Fonction pour écrire une phrase grâce à Z01
func printSentences(s string, spacesBefor int, spacesAfter int) {
	for i := 0; i < spacesBefor; i++ {
		z01.PrintRune('\n')
	}

	for _, v := range s {
		z01.PrintRune(v)
	}

	for i := 0; i < spacesAfter; i++ {
		z01.PrintRune('\n')
	}
}

// On print une matrice formaté
func printFinded(t [][]rune) {
	for i := 0; i < len(t); i++ {
		if i == 0 {
			z01.PrintRune('[')
		}
		for e := 0; e < len(t[i][:]); e++ {
			z01.PrintRune(t[i][e])
		}
		if i == len(t)-1 {
			z01.PrintRune(']')
		} else {
			z01.PrintRune(' ')
		}
	}
}
