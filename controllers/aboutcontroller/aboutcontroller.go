package aboutcontroller

import (
	"go-web-native/models/productmodel"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.Getall()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/about/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}


