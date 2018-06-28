package radio

import (
	"html/template"
	"log"
	"net/http"
)

type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Label      string
}

type PageVariables struct {
	PageTitle        string
	PageRadioButtons []RadioButton
	Answer           string
}

func LoadRadios(w http.ResponseWriter, r *http.Request) {
	Title := "Which do you prefer?"

	MyRadioButtons := []RadioButton{
		RadioButton{"animalselect", "cats", false, false, "Cats"},
		RadioButton{"animalselect", "dogs", false, false, "Dogs"},
	}

	MyPageVariables := PageVariables{
		PageTitle:        Title,
		PageRadioButtons: MyRadioButtons,
	}

	t, err := template.ParseFiles("radio/radio.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, MyPageVariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}

}

func LoadAnswer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	youranimal := r.Form.Get("animalselect")

	Title := "Your preferred animal"
	MyPageVariables := PageVariables{
		PageTitle: Title,
		Answer:    youranimal,
	}

	t, err := template.ParseFiles("radio/radio.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, MyPageVariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
