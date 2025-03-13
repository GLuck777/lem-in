package functions

import (
	"fmt"
	"os"
)

// Verifie les erreurs structurelles du fichier gestion d'Erreur fichier
func VerificationErrorFile(file []string) {
	for linefile, f := range file {
		if f == "" && linefile != 0 {
			continue
		}
		// fmt.Println("moment", string([]rune(f)[0]), "phrase", string(file[linefile]))
		// fmt.Println(string([]rune(f)[0]) == "#" && string(file[linefile]) != "##start" && string(file[linefile]) == "##end")
		if linefile == 0 {
			if VerifiLine(f) {
				fmt.Println(Color_Red, Text_Bold, "\nERROR: Invalid number of Ants", ResetBold, ResetAll)
				os.Exit(1)
			}
		} else {
			if file[linefile-1] == "##start" {
				if file[linefile] == "##end" {
					fmt.Println(Color_Red, Text_Bold, "\nERROR: invalid data format, no start room found", ResetBold, ResetAll)
					os.Exit(3)
				}
			}
			if string([]rune(f)[0]) == "#" && string(file[linefile]) != "##start" && string(file[linefile]) != "##end" {
				continue
			}
		}
	}
}

// Lors de la lecture du fichier verifie la ligne pour deceler des erreurs
func VerifiLine(line string) bool { // true if number ants problem
	switch {
	case line == "":
		return true
	case ([]rune(line)[0] == '#'):
		return true
	case line == "0":
		return true
	case VerifRune(line):
		fmt.Println("ma ligne", line)
		fmt.Println(Color_Blue, VerifRune(line), ResetAll)
		return true
	default:
		return false
	}
}

// Specificité de VerifiLine pour gérer des nombres a plusieurs lignes
func VerifRune(line string) bool {
	for i := 0; i < len(line); i++ {
		if []rune(line)[i] < 48 || []rune(line)[i] > 57 {
			return true
		}
	}
	return false
}

// Fonction Importante : Distingue les cellules Start et end pour gérer avec une autre fonction si colonies peut arriver jusqu'à la fin et montre les chemins possibles
func VerifStatus(Cells []Cells) [][]string {
	var StartRoomId int
	var StartRoomName string
	var EndRoomId int
	// var EndRoom string
	for _, cell := range Cells {
		// fmt.Println("Room:", cell.(*BasicRoom).name, "--status: ", m.Color_Green, cell.(*BasicRoom).status, m.ResetAll)
		switch cell.(*BasicRoom).Status {
		case "start":
			if VerifAdjacent(cell) {
				StartRoomId = cell.(*BasicRoom).Id
				StartRoomName = cell.(*BasicRoom).Name
				cell.(*BasicRoom).Present = true // Rend la Salle de départ occupé pour toutes les fourmis
			} else {
				fmt.Println(cell.(*BasicRoom))
				fmt.Println(Color_Red, Text_Bold, "\nERROR: Starting cell have no path", ResetBold, ResetAll)
				os.Exit(1)
			}
		case "end":
			if VerifAdjacent(cell) {
				EndRoomId = cell.(*BasicRoom).Id
				// EndRoom = cell.(*BasicRoom).name
			} else {
				fmt.Println(Color_Red, Text_Bold, "\nERROR: Ending cell have no path", ResetBold, ResetAll)
				os.Exit(1)
			}
		}
	}
	var allPaths [][]string                                              // le tableau de tableau de string affiché dans le for en bas
	currentPath := []string{StartRoomName}                               // Point de départ
	VerifTracking(Cells, StartRoomId, EndRoomId, currentPath, &allPaths) // Determine tous les chemins possible

	//////////A TESTER////////////////////
	if len(allPaths) < 1 {
		fmt.Println(Color_Red, Text_Bold, "Tracking room have no path until End", ResetBold, ResetAll)
		os.Exit(1)
	}
	/////////////////////////////////////
	allPaths = SortPaths(allPaths)
	// fmt.Println(Color_Black + "\nTous les chemins possibles du plus grand au plus petit:" + ResetAll)
	// for i, path := range allPaths {
	// 	fmt.Print(Color_Black)
	// 	fmt.Printf("Sort Path %d: %s, %v salles\n", i+1, path, len(path)-1)
	// 	fmt.Print(ResetAll)
	// }

	NewPaths := FilterPaths(allPaths, allPaths[0]) // Determine tous les chemins possibles et choisi le meilleur

	/*Adapté pour le [][]string*/
	// fmt.Println(Color_Orange + "\nTous les chemins Choisis:" + ResetAll)
	// for i, path := range NewPaths {
	// 	fmt.Print(Color_Black)
	// 	fmt.Printf("Chemin %d: %s, %v salles\n", i+1, path, len(path)-1)
	// 	fmt.Print(ResetAll)
	// }

	/*Adapté pour le [][][]string*/
	// for i, path := range NewPaths {
	// 	fmt.Printf("Chemin %d: longueur %s%v%s\n", i+1, Color_Blue, len(path), ResetAll)
	// 	for _, p := range path {
	// 		fmt.Print(Color_Black)
	// 		fmt.Println(p)
	// 		fmt.Print(ResetAll)
	// 	}
	// }

	return NewPaths
	// return allPaths
}

// Inscrit dans un tableau de tableau de string tous les chemins empruntable par les fourmis
func VerifTracking(Cells []Cells, StartRoomId, EndRoomId int, currentPath []string, allPaths *[][]string) {
	// Si la salle actuelle est la salle de fin, ajouter le chemin actuel à la liste des chemins possibles
	// fmt.Println("Start:", StartRoomId, "End Place:", EndRoomId)
	if StartRoomId == EndRoomId {
		*allPaths = append(*allPaths, append([]string{}, currentPath...))
		return
	}

	for _, cell := range Cells {
		// Ignorer la salle de départ
		if cell.(*BasicRoom).Status == "start" {
			continue
		}

		// Vérifier si la salle actuelle est connectée à la cellule en cours
		if cell.IsConnectedTo(StartRoomId) {
			// Vérifier si la salle en cours est déjà dans le chemin
			isInPath := false
			for _, p := range currentPath {
				// fmt.Println("Verification: Chemin de départ-", p, "Chemin à verifier-", cell.(*BasicRoom).name)
				if p == cell.(*BasicRoom).Name {
					// fmt.Println(m.Color_Cyan, "hahahaha", !isInPath, m.ResetAll)
					isInPath = true
					break
				}
			}

			// Si la salle en cours n'est pas déjà dans le chemin, continuer à explorer
			// fmt.Println(m.Color_Cyan, "hohohoho", !isInPath, m.ResetAll)
			if !isInPath {
				newPath := append([]string{}, currentPath...)
				newPath = append(newPath, cell.(*BasicRoom).Name)
				// fmt.Println("Resultat:", newPath)
				// Explorer récursivement
				VerifTracking(Cells, cell.(*BasicRoom).Id, EndRoomId, newPath, allPaths)
			}
		}
	}
}

// Verifie si les salle start et end on au moins une salle adjacente, evite de faire perdre du temps
func VerifAdjacent(tabAdjacent Cells) bool {
	var boolAdjacent bool = false
	for _, adjacent := range tabAdjacent.(*BasicRoom).Adjacent {
		if adjacent {
			boolAdjacent = true
			return boolAdjacent
		}
	}
	return boolAdjacent
}

// Recupère tous les chemins via *allPaths pour déterminer les intersections et filtrer les chemins à emprunter
func FilterPaths(allPaths [][]string, Paths []string) [][]string {
	var NewNewPaths [][][]string
	var NewPaths [][]string
	var ChoosenPath [][]string

	for o, path := range allPaths {
		NewPaths = nil
		// fmt.Println(Color_Red, path, "indice o", o, ResetAll)
		var intersecting bool
		NewPaths = append(NewPaths, path)
		for _, pathmatch := range allPaths[o+1:] {
			// fmt.Println(Color_Blue, "path actuel", path, ResetAll, Color_Green, "Face au path actuel", pathmatch, ResetAll)
			if VerifPath(path, pathmatch) {
				// fmt.Println("je skip car c'est le même chemin")
				continue
			}

			intersecting = false
			// fmt.Println("================> Fight !! <================")
			for i := 1; i < len(path)-1; i++ {
				for k := 1; k < len(pathmatch)-1; k++ {
					if path[i] == pathmatch[k] {
						// fmt.Println("=====>VRAI : Pour l'indice i", i, "Croisement entre", path[i], "et", pathmatch[k])
						intersecting = true
						break
					}
				}

				if intersecting {
					// fmt.Println("ils se croisent donc je break !, je passe aux autres nombre du chemin")
					break
				}
			}

			if !intersecting {
				// fmt.Println("je sors de la double boucle i et k et intersecting est toujours", intersecting)

				if VerifPathKai(pathmatch, NewPaths) { // Renvoi true si pas de croisement avec les chemins précédemment

					NewPaths = append(NewPaths, pathmatch)
				}
			}
		}
		NewNewPaths = append(NewNewPaths, NewPaths)
	}
	ChoosenPath = ChooseBestPath(NewNewPaths)
	// ChoosenPath = SortPathsInverse(ChoosenPath)
	return ChoosenPath
	// return NewNewPaths
}

/*
Recoupe le tableau [][][]string qui regroupait les différents possibilités de chemins.
Récupère l'ensemble de chemin avec le plus de chemin différent disponible
si la longueur de l'ensemble de chemin par rapport a celui tester alors un verification du chemin le plus court est effectué
*/
func ChooseBestPath(EnsemblePaths [][][]string) [][]string {
	var NewPaths [][]string
	var temp int
	var temp2 int

	for i, Paths := range EnsemblePaths {
		if i == 0 {
			NewPaths = Paths
			continue
		}
		if len(NewPaths) < len(Paths) {
			NewPaths = Paths
		} else if len(NewPaths) == len(Paths) {
			for _, v := range NewPaths {
				temp += len(v)
			}
			for _, w := range Paths {
				temp2 += len(w)
			}
			if temp > temp2 {
				NewPaths = Paths
			}
		}
	}
	return NewPaths
}

// Verifie si un chemin est similaire à un autre (Recycler dans AttackOnAnts)
func VerifPath(path, chemin []string) bool {
	var samePath bool = true
	if len(path) != len(chemin) {
		samePath = false
		return samePath
	}
	for i := 0; i < len(path); i++ {
		if path[i] == chemin[i] {
			samePath = true
		} else {
			samePath = false
			return samePath
		}
	}
	return samePath
}

/*
VerifPath revisité pour faire une vérification des chemins qui précède celui qui est testé dans FilterPath pour décider si on réalise l'ajout de ce chemin ou si il sera ignoré.
Fonction importante du filtre
*/
func VerifPathKai(pathtest []string, tabPath [][]string) bool {
	for i, path := range pathtest {
		if i == 0 || i == len(pathtest)-1 { // Je prend pas le start ni le end
			continue
		}
		for _, ensChemin := range tabPath {
			for z, pathverif := range ensChemin {
				if z == 0 || z == len(ensChemin)-1 { // Je prend pas le start ni le end
					continue
				}
				if path == pathverif {
					return false // Il y a un croisement, retourne false
				}
			}
		}
	}
	return true // Aucun croisement trouvé
}
