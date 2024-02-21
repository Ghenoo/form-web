package handlers

import (
	"fmt"
	"form-web/controllers"
	"net/http"
)

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprint(w, "erro ao fazer parse do form: %v", err)
			return
		}

		err := controllers.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))
		if err != nil {
			fmt.Fprint(w, "erro ao salvar dados: %v", err)
			return
		}
	}

	http.ServeFile(w, r, "handlers/templates/subscription_forms.html")
}
