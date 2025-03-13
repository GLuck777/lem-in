package functions

import (
	"fmt"
	"os"
	"strings"
)

func LectureFichier(file string) []string {
	file1 := "internal/colony/" + file
	content, err := os.ReadFile(file1)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return nil
	}
	// fmt.Println(content)
	Tabfile := strings.Split(strings.ReplaceAll(string(content), "\r", ""), "\n")
	// affiche ton tableau en string exactement de la meme manière que dans le fichier
	// fmt.Println(Tableau)
	return Tabfile
}
