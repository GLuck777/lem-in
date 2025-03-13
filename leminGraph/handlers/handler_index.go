package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	m "lemin/internal/functions"
)

// Function for the main page
func Index(w http.ResponseWriter, r *http.Request) {
	// Url Verification
	if r.URL.Path != "/" {
		// fmt.Println(m.Color_Cyan, "Error executing template for index page", m.ResetAll)
		Error(w, r, http.StatusNotFound)
		return
	}
	var file string
	file = ""
	file = r.FormValue("file")
	fmt.Println(file)
	LeminInfo := &m.LeminGraph{}
	if file != "" {
		LeminInfo = m.Le_min(file)
		// fmt.Println("Lemin info", LeminInfo)
	}
	// fmt.Println("Lemin_info")
	// fmt.Println(LeminInfo)
	// fmt.Println("Fin Lemin_info")
	if LeminInfo != nil {
		// fmt.Println("Infos Nombre", LeminInfo.NBANTS)
		/* Récuperer sur le site
		https://stackoverflow.com/questions/73167569/how-to-make-html-templates-recognize-a-string-as-html-in-go*/
		// template.HTML takes only one string, hence we loop over the entry slice
		// and store the slice values in htmlvalues of `HTML` type
		var htmlvalues []template.HTML

		for _, n := range LeminInfo.ParcoursHTML {
			htmlEncapsulate := template.HTML(n)
			htmlvalues = append(htmlvalues, htmlEncapsulate)
		}
		data := m.ImprimeGraph{
			ParcoursHTML: htmlvalues,
			Soluce:       LeminInfo.Soluce,
			Connexion:    LeminInfo.Connexion,
			NBANTS:       LeminInfo.NBANTS,
			StartRoom:    LeminInfo.StartRoom,
			Title: file,
		}
		/*
			Load Index
		*/

		tmpl := template.Must(template.ParseFiles("assets/templates/index_page.tmpl"))
		tmpl.Execute(w, data)
		// errTmpl := tmpl.ExecuteTemplate(w, "index_page.tmpl", data)
		// if errTmpl != nil {
		// 	fmt.Println(http.StatusInternalServerError)
		// 	return
		// }

	}
}

// Function to verify the method passed
// func method(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "POST":
// 		// nothing
// 	case "GET":
// 		// nothing
// 	default:
// 		Error(w, r, http.StatusBadRequest)
// 		return
// 	}
// }
