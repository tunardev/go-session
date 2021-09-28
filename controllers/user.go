package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/tunardev/go-session/models"
	"github.com/tunardev/go-session/utils"
	"gopkg.in/asaskevich/govalidator.v9"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (c Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		defer r.Body.Close()
        body, err := io.ReadAll(r.Body)

	   if err != nil {
	  json.NewEncoder(w).Encode(utils.NewError(err, false))
	    return
	  }

	   var newUser models.User
	  err = json.Unmarshal(body, &newUser)
	  if err != nil {
	   json.NewEncoder(w).Encode(utils.NewError(err, false))
	   return
	  }

	newUser.Email = utils.NormalizeEmail(newUser.Email)

	if !govalidator.IsEmail(newUser.Email) {
		 json.NewEncoder(w).Encode(utils.NewError(err, false))
		 return
	}

	newUser.Password, err = utils.HashPassword(newUser.Password)
	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(err, false))
		return
	}

	newUser.Id = bson.NewObjectId()
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = newUser.CreatedAt

	err = c.Database.DB("test").C("users").Insert(&newUser)
	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(err, false))
		return
	}

	session, err := c.Session.Get(r, "auth")
	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(err, false))
		return
	}

	session.Values["authenticated"] = true
	session.Values["userId"] = newUser.Id.Hex()
 
	if err = session.Save(r, w); err != nil {
		json.NewEncoder(w).Encode(utils.NewError(err, false))
	}	

	json.NewEncoder(w).Encode(newUser)
     return 
	} 

	json.NewEncoder(w).Encode(utils.NewError(errors.New("Wrong http method."), false))
}

func (c Controller) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		defer r.Body.Close()
        body, err := io.ReadAll(r.Body)

	   if err != nil {
	  json.NewEncoder(w).Encode(utils.NewError(err, false))
	    return
	  }

	   var user models.User
	  err = json.Unmarshal(body, &user)
	  if err != nil {
	   json.NewEncoder(w).Encode(utils.NewError(err, false))
	   return
	  }

	  user.Email = utils.NormalizeEmail(user.Email)

	if !govalidator.IsEmail(user.Email) {
		 json.NewEncoder(w).Encode(utils.NewError(err, false))
		 return
	}

	var dbUser models.User

	err = c.Database.DB("test").C("users").Find(bson.M{"email": user.Email}).One(&dbUser)

	if err == mgo.ErrNotFound {
		json.NewEncoder(w).Encode(utils.NewError(errors.New("user not found"), false))
		return 
	 }

	 success := utils.CheckPasswordHash(user.Password, dbUser.Password)

	 if success != true {
		json.NewEncoder(w).Encode(utils.NewError(errors.New("incorrect password"), false))
		return
	 }	

	 session, err := c.Session.Get(r, "auth")
	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(err, false))
		return
	}

	session.Values["authenticated"] = true
	session.Values["userId"] = dbUser.Id.Hex()

	session.Save(r, w)

	 json.NewEncoder(w).Encode(dbUser)

	 return
	}
	
	json.NewEncoder(w).Encode(utils.NewError(errors.New("Wrong http method."), false))
}

func (c Controller) Logout(w http.ResponseWriter, r *http.Request) {

	session, err := c.Session.Get(r, "auth")
	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(errors.New("Warning logout"), false))
		return
	}

	session.Values["authenticated"] = false
	session.Values["userId"] = ""

	session.Save(r, w)

	json.NewEncoder(w).Encode(utils.NewError(errors.New("Success logout"), false))
}

func (c Controller) Me(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		session, err := c.Session.Get(r, "auth")
	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(err, false))
		return
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
        json.NewEncoder(w).Encode(utils.NewError(errors.New("not authenticated"), false))
        return
	}

	userId, ok := session.Values["userId"].(string)
	if  !ok {
        json.NewEncoder(w).Encode(utils.NewError(errors.New("not authenticated"), false))
        return
	}

	var user models.User
	err = c.Database.DB("test").C("users").FindId(bson.ObjectIdHex(userId)).One(&user)

	if err != nil {
		json.NewEncoder(w).Encode(utils.NewError(errors.New("user not found"), false))
		return
	}
   json.NewEncoder(w).Encode(user)

   return
	}

	json.NewEncoder(w).Encode(utils.NewError(errors.New("Wrong http method."), false))
}