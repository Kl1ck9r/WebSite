package request

import (
	//"github.com/cmd/package/server"
   	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request){
	//var Error utils.HandlerStateCode

	//if r.URL.Path!="/"{
	//	http.Error(w,"404 not found",http.StatusNotFound)
	//}

	switch r.Method{

	case "POST":
	

	case "GET":
	

	case "DELETE":
		

	case "PUT":
	
	
	default:
		
	}

}
