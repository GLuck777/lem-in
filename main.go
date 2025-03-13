package main

import (
	"fmt"
	"net/http"
	"os"

	m "lemin/internal/functions"
	h "lemin/leminGraph/handlers"
)

const port = ":7653"

func main() {
	if len(os.Args) > 1 {
		var file string = os.Args[1]
		m.Le_min(file)
		// fmt.Println("LeminInfo parcours", LeminInfo.ParcoursHTML)
		// fmt.Println("LeminInfo soluce", LeminInfo.Soluce)
	} else {
		// Import the css
		// css := http.FileServer(http.Dir("csshandler"))
		// http.Handle("/csshandler/", http.StripPrefix("/csshandler", css))

		// Import the Directory assets
		assets := http.FileServer(http.Dir("assets"))
		http.Handle("/assets/", http.StripPrefix("/assets/", assets))

		// Page existing with her function
		http.HandleFunc("/", h.Index)
		// Server access
		fmt.Println(m.Color_Cyan, "http://localhost"+port+" - Server Started on port", port, m.ResetAll)
		http.ListenAndServe(port, nil)

	}
}
