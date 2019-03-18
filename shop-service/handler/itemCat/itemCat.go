package itemCat

import (
	"fmt"
	"net/http"
)

/*
--request
id=3

--response

*/
func FindOne(w http.ResponseWriter,r *http.Request)  {

	//r.ParseForm()
	//id := utils.StringToInt64(r.Form.Get("id"))

	fmt.Println("======================")
}
