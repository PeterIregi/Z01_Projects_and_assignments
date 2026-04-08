package handlers

import(
	"net/http"
	"strings"
	errors"ascii-art-web/utils"
)

func AsciiArt(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
            errors.ErrorHandler(w, r, http.StatusBadRequest, "Unable to parse form")
            return
        }
		
		Fdata := Data{  Text: r.FormValue("text-input"), Banner: r.FormValue("banner-style") }
		err = Parser(&Fdata)

		if err != nil{
			errors.ErrorHandler(w, r, http.StatusInternalServerError, "Error generating ASCII art")
			return
		}

		safeOutput := strings.ReplaceAll(Fdata.Output, " ", "&nbsp;")
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(safeOutput))

	default:
		errors.ErrorHandler(w, r, http.StatusMethodNotAllowed,"Method not allowed",)
        return
	}
}