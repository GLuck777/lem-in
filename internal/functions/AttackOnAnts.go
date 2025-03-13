package functions

import (
	"fmt"
	"strconv"
	"strings"
)

// Donne nom aux Fourmis pour la fonction finale
func (nbAnts *Room) CreationAnts(allPaths [][]string) []Colony {
	var compteur int
	Ant := &Ants{}
	var Colony []Colony
	nbPath := len(allPaths)

	for i := 1; i <= nbAnts.N; i++ {
		if compteur == nbPath {
			compteur = 0
		}
		Ant.Name = "L" + strconv.Itoa(i)
		Ant.Cell = nbAnts.Start
		Ant.Path = allPaths[compteur]
		Colony = append(Colony, NewAnt(Ant.Name, Ant.Cell, Ant.Path))
		// fmt.Printf("%s-%s %s\n", Ant.Name, Ant.Cell, Ant.Path) // affichage important pour voir la path de ta fourmis
		// fmt.Println(m.Color_Orange+"Ant", i, ":"+Ant.name+"-"+Ant.cell, m.ResetAll)
		compteur++
	}
	// fmt.Println("Nombre de fourmis: ", nbAnts.n)
	return Colony
}

func (NbAnts *Room) Le_min(Cells []Cells, allPaths [][]string) []string {
	var SortieResultat []string
	Colony := NbAnts.CreationAnts(allPaths) // creation des fourmis
	///////////////////////////////
	// for _, f := range Colony { // montre les fourmis, le nom et la salle où elles se trouvent
	// fmt.Print("Fourmis: ", f.(*Ants).Name, " - Dans la salle: ", f.(*Ants).Cell, "à le chemin:",f.(*Ants).Path,",\n")
	// }
	///////////////////////////////
	// for _, cell := range Cells {
	// 	// fmt.Println("cell together", cell)
	// 	fmt.Println("Salle", cell.(*BasicRoom).Name, "fourmis presents ?", cell.(*BasicRoom).Present)
	// }
	///////////////////////////////
	// Afficher tous les chemins possibles
	// fmt.Println(m.Color_Orange + "\nTous les chemins possibles entre la salle de départ et la salle de fin :" + m.ResetAll)
	for i, path := range allPaths {
		fmt.Printf("Chemin %d: %s, %v salles\n", i+1, path, len(path)-1)
	}
	///////////////////////////////

	// Créez une variable pour compter le nombre de fourmis ayant atteint la fin
	var fourmisTerminees int
	var tour int
	var Compteur int
	var salle string
	var tempPath []string

	// Parcourez chaque étape du chemin
	for {
		tour++
		// fmt.Printf("\n%sTour %d%s\n", Text_Bold, tour, ResetBold) // Écrit les tours
		Compteur++
		tempPath = nil
		for _, fourmis := range Colony {
			if fourmis.(*Ants).Cell == NbAnts.End { // je me débarasse de mes fourmis
				continue
			}
			if fourmis == Colony[len(Colony)-1] {
				if fourmisTerminees == len(Colony)-2 && len(fourmis.(*Ants).Path)-2 == len(allPaths[0]) {
					fourmis.(*Ants).Path = allPaths[0]
				}
			}
			salle = FindRoom(fourmis.(*Ants).Cell, fourmis.(*Ants).Path, NbAnts.End)
			if VerifFourmis(Colony, fourmis.(*Ants).Cell, fourmis.(*Ants).Path, NbAnts.End) {
				if tempPath == nil || !VerifPath(fourmis.(*Ants).Path, tempPath) { // à améliorer ? non un seul chemin aussi court ne pourra exister, tempPath == nil là juste pour gagner du temps
					// fmt.Printf("%s-%s ", fourmis.(*Ants).Name, salle)
					SortieResultat = append(SortieResultat, fourmis.(*Ants).Name+"-"+salle+" ")
					fourmis.(*Ants).Cell = salle
				}

				if fourmis.(*Ants).Cell == NbAnts.End && len(fourmis.(*Ants).Path) == 2 {
					tempPath = fourmis.(*Ants).Path
				}
			}

			// Ajout une temporisation pour simuler un tour au besoin
			// time.Sleep(500 * time.Millisecond)

			/* Vérifiez si la salle où se trouve la fourmis est la salle de fin !
			Associé au if NbAnts.End au début de la boucle pour ne pas avoir de doublon*/
			if fourmis.(*Ants).Cell == NbAnts.End {
				// fmt.Printf("Fourmis %s a atteint la fin.\n", fourmis)
				fourmisTerminees++
			}
		}
		// fmt.Println()
		SortieResultat = append(SortieResultat, "\n")

		// Condition de sortie : Si toutes les fourmis ont atteint la fin
		if fourmisTerminees == len(Colony) {
			// fmt.Printf("%d fourmis ont atteint la fin.\n", fourmisTerminees)
			SortieResultat = append(SortieResultat, "Fini en "+strconv.Itoa(tour)+" tour(s)")
			// fmt.Println("Fini en", tour, "tour(s)")
			var Res []string
			var entree string
			for i, ligne := range SortieResultat {
				if ligne == "\n" {
					entree = strings.TrimSpace(entree)
					Res = append(Res, entree)
					Res = append(Res, "\n")
					entree = ""
				} else {
					entree += ligne + " "
				}
				if i == len(SortieResultat)-1 {
					Res = append(Res, entree)
				}
			}
			// return SortieResultat
			return Res
		}
	}
}

/*
Fonction basique importante : verifie le chemin actuel où se trouve ta fourmis, verifie si le chemin qui suit est occupé par une fourmis et
permet ou non son avancement sur son chemin
*/
func VerifFourmis(Colony []Colony, FourmisCell string, allPaths []string, SalleFin string) bool {
	var temp int
	for cpt, v := range allPaths {
		if FourmisCell == v {
			temp = cpt
			break
		}
	}
	for _, fourmis := range Colony {
		if allPaths[temp+1] == SalleFin {
			return true
		}
		if fourmis.(*Ants).Cell == allPaths[temp+1] {
			return false
		}
	}

	return true
}

/*
Fonction plus utilisé: (voir git master/le-minTerminé ; Commit: 2b2678ca28c9674b3c6b127f8608b2074be794fa)
Permet de retrouver le chemin de la fourmis grace à sa chambre actuelle et ressort le nombre int correspondant au chemin associé à FindRoom
*/
func VerifNumberPath(FourmisCell string, allPaths [][]string) int {
	var temp int = 0
	var finished bool = false
	for cpt, Path := range allPaths {
		for r, room := range Path {
			if r == 0 || r == len(Path)-1 {
				continue
			}
			if FourmisCell == room {
				temp = cpt
				finished = true
				// fmt.Println(Color_Cyan+" ===>Temp", temp, ResetAll) ///////////////
				break
			}
		}
		if finished {
			break
		}
	}
	return temp
}

/*
Fonction plus utilisé: Permet de déterminer la salle qui suit celle dans laquelle ta fourmis se trouve et la ressort
associé à VerifNumberPath
*/
func FindRoom(FourmisCell string, allPaths []string, SalleFin string) string {
	var temp string
	for cpt, v := range allPaths {
		// fmt.Println("Cell fourmis", FourmisCell, "Contre", v) ///////////////
		if v == SalleFin {
			temp = allPaths[cpt]
			return temp
		}
		if FourmisCell == v {
			temp = allPaths[cpt+1]
			// fmt.Println("Temp", temp) ///////////////
			break
		}
	}
	return temp
}
