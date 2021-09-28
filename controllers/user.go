package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/tunardev/dev-post-server/models"
	"github.com/tunardev/dev-post-server/utils"
	"gopkg.in/asaskevich/govalidator.v9"
	"gopkg.in/mgo.v2/bson"
)

func (c Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		defer r.Body.Close()
        body, err := io.ReadAll(r.Body)

	   if err != nil {
	  json.NewEncoder(w).Encode(utils.NewError(err))
	    return
	  }

	   var newUser models.User
	  err = json.Unmarshal(body, &newUser)
	  if err != nil {
	   json.NewEncoder(w).Encode(utils.NewError(err))
	   return
	  }

	newUser.Email = utils.NormalizeEmail(newUser.Email)

	if !govalidator.IsEmail(newUser.Email) {
		 json.NewEncoder(w).Encode(utils.NewError(err))
		 return
	}

	newUser.Password, err = utils.HashPassword(newUser.Password)
	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(err))
		return
	}

	newUser.Id = bson.NewObjectId()

	err = c.Database.DB("test").C("users").Insert(&newUser)
	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(err))
		return
	}

	session, err := c.Session.Get(r, "auth")
	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(err))
		return
	}

	session.Values["authenticated"] = true
	session.Values["userId"] = newUser.Id.Hex()

	session.Save(r, w)

	json.NewEncoder(w).Encode(newUser)
     return 
	} 

	json.NewEncoder(w).Encode(utils.NewError(errors.New("Wrong http method.")))
}

func (c Controller) SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "SignIn")
}

func (c Controller) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Logout")
}

func (c Controller) Me(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		session, err := c.Session.Get(r, "auth")
	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(err))
		return
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        json.NewEncoder(w).Encode(utils.NewError(errors.New("not authenticated")))
        return
	}

	userId, ok := session.Values["userId"].(string)
	if  !ok {
        json.NewEncoder(w).Encode(utils.NewError(errors.New("not authenticated")))
        return
	}

	var user models.User
	err = c.Database.DB("test").C("users").FindId(bson.ObjectIdHex(userId)).One(&user)

	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(err))
		return
	}
   json.NewEncoder(w).Encode(user)

   return
	}

	json.NewEncoder(w).Encode(utils.NewError(errors.New("Wrong http method.")))
}