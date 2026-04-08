package handlers

import (
	"ascii-art-web/asciiart"
	"net/http"
	"strings"
	"text/template"
	errors"ascii-art-web/utils"
)

type Data struct{
	Text string
	Banner string
	Output string

}


func Parser(data *Data) error {

	inputLines := strings.Split(data.Text, "\n")
//Read the lines from the file
	file := "banners/" + data.Banner + ".txt"

    // Read the file
    lines, err := asciiarts.ReadLinesFromFile(file)
	if err != nil {
		return err
	}
	// Read file and map the characters.
	charmap := asciiarts.MapToASCII(lines)

	// output our character
	data.Output = asciiarts.OutputCharacters(inputLines, charmap)
	return nil
}
 
//Errt := template.parse

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errors.ErrorHandler(w,r,http.StatusNotFound,"Path is not-existent")
		return
	}

	var tmpl,tErr = template.ParseFiles("templates/index.html")
	if tErr != nil {
		errors.ErrorHandler(w, r, http.StatusNotFound, "template not found")
		return
	}

	switch r.Method{
	case http.MethodGet:
		tmpl.Execute(w, nil)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
           errors.ErrorHandler(w,r , http.StatusBadRequest,"Unable to parse form")
            return
        }
		
		Fdata := Data{  Text: r.FormValue("text-input"), Banner: r.FormValue("banner-style") }
		err = Parser(&Fdata)

		if err != nil{
			errors.ErrorHandler(w, r, http.StatusInternalServerError, err.Error())
			return
		}
		tmpl.Execute(w, Fdata)

	default:
		errors.ErrorHandler(w,r, http.StatusBadRequest, "Method not allowed")
        return
	}
}

func ClearOutput(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte(""))
}

