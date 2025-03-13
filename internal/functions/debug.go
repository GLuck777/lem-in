package functions

import (
	"fmt"
	"strconv"
)

// DEBUG : affiche les informations du fichier dans le terminal
func EcritureFichier(file []string) {
	var Startant int

	for linefile, f := range file {
		if f == "" && linefile != 0 {
			continue
		}
		if linefile == 0 {
			if f == "" {
				fmt.Println("Pas de fourmis renseigné")
			} else {
				Startant, _ = strconv.Atoi(f)
				fmt.Println(Color_Black+Text_Underline+"Nombre de fourmis"+ResetUnderline+": "+ResetAll+Color_Blue, Startant, ResetAll)
			}
		} else {
			if file[linefile-1] == "##start" {
				fmt.Println(Color_Black+Text_Underline+"Point de départ"+ResetUnderline+": "+ResetAll+Color_Green, string([]rune(f)[0]), ResetAll)

				fmt.Println(Color_Red, f, ResetAll)
			} else if file[linefile-1] == "##end" {
				fmt.Println(Color_Black+Text_Underline+"Arrivée"+ResetUnderline+": "+ResetAll+Color_Red, string([]rune(f)[0]), ResetAll)
				fmt.Println(Color_Red, f, ResetAll)
			} else {
				fmt.Println(Color_Red, f, ResetAll)
			}
		}
	}
}

// Verifie les tableaux boolean de chaque cell pour en verifier les relations entre elles
func TabConnexion(Cells []Cells) {
	// Affichage des tableaux de connexions
	fmt.Println(Color_Green, "Cell connection arrays:")
	for _, cell := range Cells {
		basicRoom := cell.(*BasicRoom)
		fmt.Printf("Cell %s: %+v\n", basicRoom.Name, basicRoom.Adjacent)
	}
	fmt.Println(ResetAll)
}

func TabPosition(Cells []Cells) {
	// Affichage des tableaux de connexions
	PX, PY := lePlusgrand(Cells)
	// fmt.Println("x max: ", PX, " y max: ", PY) ////////////////

	// fmt.Println(Color_Green, "Cell Position Cell:")
	// for _, cell := range Cells {
	// 	basicRoom := cell.(*BasicRoom)
	// 	fmt.Printf("Cell %s en Position X: %s, en position Y: %s\n", basicRoom.Name, basicRoom.PosX, basicRoom.PosY)
	// }
	// fmt.Println(ResetAll)

	// Créer une matrice pour représenter le plan
	plan := make([][]string, PY+1)
	for i := range plan {
		plan[i] = make([]string, PX+1)
		for j := range plan[i] {
			plan[i][j] = "  " // Trois espaces entre les chiffres
			// plan[i][j] = "" // pas d'espaces entre les chiffres
		}
	}
	for _, cell := range Cells {
		basicRoom := cell.(*BasicRoom)
		char := basicRoom.Name
		x, _ := strconv.Atoi(basicRoom.PosX)
		y, _ := strconv.Atoi(basicRoom.PosY)

		if basicRoom.Status == "start" {
			Salle := char
			plan[y][x] += Salle
		} else if basicRoom.Status == "end" {
			Salle := (char)
			plan[y][x] += Salle
		} else {
			Salle := (char)
			plan[y][x] += Salle
		}
	}
	// Afficher le plan
	for i := PY; i >= 0; i-- {
		for j := 0; j <= PX; j++ {
			fmt.Print(plan[i][j])
		}
		fmt.Println()
	}
}
