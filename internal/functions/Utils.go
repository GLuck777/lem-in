package functions

import (
	"strconv"
	"strings"
)

// Regroupe des petites fonctions de type fonctionnelles

// Enleve l'élément d'un tableau selon sont index
func Remove(slice [][]string, index int) [][]string {
	return append(slice[:index], slice[index+1:]...)
}

// Inverse un tableau
func Swap(slice [][]string) [][]string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// Met un tableau de len la plus petite len à la plus grande
func SortPaths(allPaths [][]string) [][]string {
	// var SortallPaths [][]string
	for range allPaths {
		for i := range allPaths {
			if i == len(allPaths)-1 {
				// fmt.Println("break2")
				break
			}
			if len(allPaths[i]) > len(allPaths[i+1]) {
				allPaths[i], allPaths[i+1] = allPaths[i+1], allPaths[i]
			}
		}
	}
	return allPaths
}

// de la plus grande len à la plus petite
func SortPathsInverse(allPaths [][]string) [][]string {
	// var SortallPaths [][]string
	for range allPaths {
		for i := range allPaths {
			if i == len(allPaths)-1 {
				// fmt.Println("break2")
				break
			}
			if len(allPaths[i]) < len(allPaths[i+1]) {
				allPaths[i], allPaths[i+1] = allPaths[i+1], allPaths[i]
			}
		}
	}
	return allPaths
}

// Sépare la ligne string en tableaux pour récuperer le mot entier en indice 0
func SplitString(linestr string) string {
	newtabline := strings.Split(linestr, " ")
	return newtabline[0]
}

func LenSplitString(linestr string) []string {
	newtabline := strings.Split(linestr, " ")
	return newtabline
}

// Compte le nombre de cell
func CounterCell(file []string) int {
	RoomCount := 0
	// Pour établir les tableaux de bool pour chaque chambre
	for linefile, p := range file {
		if p == "" {
			continue
		}
		// fmt.Println("ça : " + file[linefile])
		if linefile == 0 {
		} else {
			if []rune(p)[0] != '#' && string([]rune(p)[1]) != "-" {
				// fmt.Println(string([]rune(p)[0]))
				RoomCount++
			}
		}
	}
	return RoomCount
}

// Pour lemin Graph permet de créer le html pour les chambre dans l'index_page_tmpl
func (NbAnts *Room) SetPosition(Cells []Cells) []string {
	// Affichage des tableaux de connexions
	PX, PY := lePlusgrand(Cells)

	// Créer une matrice pour représenter le plan
	plan := make([][]string, PY+1)
	for i := range plan {
		plan[i] = make([]string, PX+1)
		for j := range plan[i] {
			// plan[i][j] = "  " // Trois espaces entre les chiffres
			plan[i][j] = "" // pas d'espaces entre les chiffres
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

	var PathURL []string
	// var PathURL string
	// Afficher le plan
	for i := PY; i >= 0; i-- {
		PathURL = append(PathURL, "<div class=\"Row\">\n")
		// PathURL += "<div class=\"row\">"
		for j := 0; j <= PX; j++ {
			if plan[i][j] == "" {
				PathURL = append(PathURL, "<div class=\"Cube\"></div>\n")
				// PathURL += "<div class=\"Cube\"></div>"
				// fmt.Print(plan[i][j])
			} else {
				if plan[i][j] == NbAnts.Start {
					PathURL = append(PathURL, "<div class=\"Cube rond start\" id=\""+plan[i][j]+"\">"+plan[i][j]+"</div>\n")
				} else if plan[i][j] == NbAnts.End {
					PathURL = append(PathURL, "<div class=\"Cube rond end\" id=\""+plan[i][j]+"\">"+plan[i][j]+"</div>\n")
				} else {
					PathURL = append(PathURL, "<div class=\"Cube rond\" id=\""+plan[i][j]+"\">"+plan[i][j]+"</div>\n")
				}
				// PathURL += "<div class=\"Cube\"></div>"
			}
		}
		PathURL = append(PathURL, "</div>\n")
		// PathURL += "</div>"
		// fmt.Println()
	}
	// fmt.Println("PathURL\n", PathURL)
	return PathURL
}

func lePlusgrand(Cells []Cells) (int, int) {
	var POX int
	var POY int
	var testX int
	var testY int
	// fmt.Println(Color_Red)
	for c, cell := range Cells {
		basicRoom := cell.(*BasicRoom)
		testX, _ = strconv.Atoi(basicRoom.PosX)
		testY, _ = strconv.Atoi(basicRoom.PosY)
		if c == 0 {
			POX, _ = strconv.Atoi(basicRoom.PosX)
			POY, _ = strconv.Atoi(basicRoom.PosY)
		}
		// fmt.Println("x enregistré", POX, "<", "x testé", basicRoom.PosX)
		if POX < testX {
			// fmt.Println("je garde", testX)
			POX = testX
		}
		// fmt.Println("y enregistré", POY, "<", "y testé", basicRoom.PosY)
		if POY < testY {
			// fmt.Println("je garde", testY)
			POY = testY
		}
	}
	// fmt.Println(ResetAll)
	// fmt.Printf("Resultat final Position X maximal: %s, et position Y maximale: %s\n", POX, POY)
	return POX, POY
}
