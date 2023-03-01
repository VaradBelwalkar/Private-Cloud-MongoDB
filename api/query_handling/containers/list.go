package containers

import (

	"context"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	ca "github.com/VaradBelwalkar/Private-Cloud-MongoDB/api/container_apis"
	as "github.com/VaradBelwalkar/Private-Cloud-MongoDB/api/auth_service"
	"github.com/docker/docker/client"
)



func Container_List(w http.ResponseWriter, r *http.Request) {
	check,username:=as.Verify_Auth(w,r)
	if check!=true{
		return
	}

 	ctx := context.Background()
	//Extracting required string from the request Structure
	//Get the requested Image from from the request-URL and pass it to the Container handler
	containerInfo,err:=ca.OwnedContainerInfo(ctx,username)	
	if err!=200{
		if err ==500{
		w.WriteHeader(http.StatusInternalServerError)
		return
		}
		return
	}
	b, _ := json.Marshal(containerInfo)
	w.Write(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Send reponse in the body
	

}
