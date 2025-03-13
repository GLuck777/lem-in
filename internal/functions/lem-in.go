package functions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	// m "lemin/assets/functions"
)

var (
	file         []string
	startliaison int
	CellNumber   int

	Sentence     []string
	roomName     string
	posX         string
	posY         string
	ParcoursHTML []string // chemin html pour mes chambres
	Soluce       []string
)

func Le_min(Nfile string) *LeminGraph {
	idRoom := 0
	Verifroom := &Room{}
	var Cells []Cells
	// filez = "exemple00.txt"
	if Nfile != "" {
		if strings.HasSuffix(Nfile, ".txt") {
			file = LectureFichier(Nfile)
		} else {
			fmt.Println(Color_Red, Text_Bold, "Error reading format file", ResetBold, ResetAll)
			os.Exit(1)
		}
	} else {
		fmt.Println("Expected 1 Argument for filename")
		fmt.Println("Error reading format file .txt")
		os.Exit(1)
	}
	// EcritureFichier(file) // lit le fichier d'information

	VerificationErrorFile(file)
	RoomCount := CounterCell(file) // Compteur de chambres pour le tableau bool
	for linefile, f := range file {
		switch {
		case f == "" && linefile != 0:
			continue
		case linefile == 0:
			if VerifiLine(f) {
				// fmt.Println("invalid number of Ants")
				// os.Exit(1)
			} else {
				Verifroom.N, _ = strconv.Atoi(f)
				// fmt.Println(Color_Black+Text_Underline+"Number of ants"+ResetUnderline+":"+ResetAll, Verifroon)
			}
		case linefile != 0:
			if file[linefile-1] == "##start" {
				// if string([]rune(f)[1]) != "" {
				mot := SplitString(f)
				Verifroom.Start = mot
				// } else {
				// VerifrooStart = string([]rune(f)[0])
				// }
				// fmt.Println(Color_Black+Text_Underline+"Start room"+ResetUnderline+":"+ResetAll, Verifroostart)
			}

			if file[linefile-1] == "##end" {
				// if string([]rune(f)[1]) != "" {
				mot := SplitString(f)
				Verifroom.End = mot
				// } else {
				// VerifrooEnd = string([]rune(f)[0])
				// }
				// fmt.Println(Color_Black+Text_Underline+"End room"+ResetUnderline+":"+ResetAll, Verifrooend)
			}

			// fmt.Println("Pour", f, "on a", len(LenSplitString(f)), "de longueur")
			if []rune(f)[0] != '#' && len(LenSplitString(f)) != 1 { //&& string([]rune(f)[1]) != "-"
				// idRoom commence a 0
				// if string([]rune(f)[1]) != "" {
				Sentence = LenSplitString(f)
				roomName = Sentence[0]
				posX = Sentence[1]
				posY = Sentence[2]
				// } else {
				// roomName = string([]rune(f))
				// }
				if string([]rune(f)[1]) != "" {
					mot := LenSplitString(f)[0]
					if mot == Verifroom.Start {
						Cells = append(Cells, NewRoom(idRoom, roomName, RoomCount, "start", posX, posY))
						// fmt.Println(Color_Red, Cells[len(Cells)-1], ResetAll) Ultime test pour verifier si le start est donnée à une salle ou non
					} else if mot == Verifroom.End {
						Cells = append(Cells, NewRoom(idRoom, roomName, RoomCount, "end", posX, posY))
						// fmt.Println(Color_Blue, Cells[len(Cells)-1], ResetAll) Verif02
					} else {
						Cells = append(Cells, NewRoom(idRoom, roomName, RoomCount, "intermediaire", posX, posY))
						// fmt.Println(Cells[len(Cells)-1]) // Verif03
					}
				} else if string([]rune(f)[0]) == Verifroom.Start {
					Cells = append(Cells, NewRoom(idRoom, roomName, RoomCount, "start", posX, posY))
				} else if string([]rune(f)[0]) == Verifroom.End {
					Cells = append(Cells, NewRoom(idRoom, roomName, RoomCount, "end", posX, posY))
				} else {
					Cells = append(Cells, NewRoom(idRoom, roomName, RoomCount, "intermediaire", posX, posY))
				}
				idRoom++
			}
			// if string([]rune(f)[1]) == "-" && len(SplitString(f)) == 1 {
			// 	startliaison = linefile
			// 	break
			// }
		}
	}

	// Initialisation des chambres

	// fmt.Println("\n"+Color_Yellow, "Initializing links between rooms...", ResetAll)
	var TableConnexion []string
	startliaison = 0
	for startliaison < len(file) {
		parts := strings.Split(file[startliaison], "-")
		if len(parts) == 2 {
			firstcell := parts[0]  // 1
			secondcell := parts[1] // 2
			// fmt.Printf("%s-%s\n", firstcell, secondcell)
			TableConnexion = append(TableConnexion, firstcell+" est connectée à "+secondcell)
			for _, cell := range Cells {
				if cell.(*BasicRoom).Name == firstcell {
					cell.ConnectTo(secondcell, Cells) // ici on connecte la première cellule avec la seconde
				} else if cell.(*BasicRoom).Name == secondcell {
					cell.ConnectTo(firstcell, Cells) // on fait de même entre la second et la première
				}
			}
			/*fmt.Print(Color_Yellow)
			fmt.Printf("Room %s is now connected to %s\n", firstcell, secondcell)
			fmt.Print(ResetAll)*/
		}
		startliaison++
	}
	// Affiche les positions des chambres dans le cmd
	// TabPosition(Cells) // voir position des salles en x et y
	//////////////////////////////////

	// Affiche les Chambres connectées
	// TabConnexion(Cells) // voir connexion entre les chambres
	//////////////////////////////////

	/* multi fonction, sert de point central pour les fonctions :
	- VerifTracking qui permet de récupérer toutes les chemin possibles
	- SortPaths qui permet de ranger mon [][]string selon la longueur des chemin dans l'ordre décroissant
	- FilterPaths qui permet de récuperer uniquement les chemins qui ne se croisent pas entre eux
	*/
	allPaths := VerifStatus(Cells)

	// Mise en place du systeme d'avancement des fourmis
	Soluce = Verifroom.Le_min(Cells, allPaths)
	for _, sR := range Soluce {
		fmt.Printf("%s", sR)
	}
	fmt.Println() // affichage présentable
	///////////////////////////////////////////////////////////////////////
	// Crée les positions des chambres dans une string pour la page web
	ParcoursHTML = Verifroom.SetPosition(Cells)
	// fmt.Println("ParcoursHTML", ParcoursHTML)
	//////////////////////////////////
	Graphique := &LeminGraph{}

	Graphique.ParcoursHTML = ParcoursHTML
	Graphique.Soluce = Soluce
	Graphique.StartRoom = Verifroom.Start
	Graphique.EndRoom = Verifroom.End
	Graphique.NBANTS = Verifroom.N // je voudrais utiliser Colony qui se trouve dans AttackOnAnts pour avoir
	// Graphique.Infos = Verifroom
	Graphique.Connexion = TableConnexion
	// Colony si on veux les fourmis en détail mais sinon utiliser "Soluce"

	return Graphique
}
