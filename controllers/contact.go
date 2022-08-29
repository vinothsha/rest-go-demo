package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-go-demo/database"
	// "github.com/gin-gonic/gin"
)

type Contact_Form struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
	Subject string `json:"subject"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var Contact Contact_Form
	json.Unmarshal(requestBody, &Contact)
	ValidEmail := ValidateEmail(Contact.Email)
	ValidMobile := ValidateMobile(Contact.Contact)
	if ValidEmail && ValidMobile {
		// database.Connector.AutoMigrate(&Contact)
		database.Connector.Create(&Contact)
		w.WriteHeader(http.StatusCreated)
		resp := "Successfully Posted"
		fmt.Println(resp)
		json.NewEncoder(w).Encode(resp)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		resp := "Invalid Mobile or Email"
		fmt.Println(resp)
		json.NewEncoder(w).Encode(resp)
	}

}

// func CreateContact(c *gin.Context){
// 	var Contact Contact_Form
// 	if err:=c.ShouldBindJSON(&Contact);err!=nil{
// 		log.Fatalln(err)
// 	}
// 	database.Connector.AutoMigrate(&Contact)
// 	database.Connector.Create(&Contact)
// 	c.JSON(http.StatusOK,Contact)
// }

func Getall(w http.ResponseWriter, r *http.Request) {
	var AllFroms []Contact_Form
	database.Connector.Find(&AllFroms)
	json.NewEncoder(w).Encode(&AllFroms)
}
