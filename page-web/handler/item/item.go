package item

import (
	"fmt"
	"net/http"
)

func GetItem(w http.ResponseWriter,r *http.Request)  {

	fmt.Println("=========================")
	r.ParseForm()

	id := r.Form.Get("id")

	fmt.Println("id:",id)
}