package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	config "user/config"
	db "user/db"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TokenClaims struct {
	UserEmail string
	jwt.StandardClaims
}
type AuthResponse struct {
	SessionCookie  string `json:"sessionCookie"`
	ExpirationTime int64  `json:"expirationTime"`
	EmailID        string `json:"emailID"`
}
type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	EmailID      string             `json:"emailID,omitempty" bson:"emailId,omitempty"`
	Name         string             `json:"name,omitempty"`
	CreatedDate  time.Time          `json:"createdDate,omitempty"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	PasswordHash []byte             `json:"-"`
	Active       bool               `json:"active,omitempty" bson:"active"`
}

var user *User
var collection *mongo.Collection
var signedString string
var auth *AuthResponse

func init() {
	log.Info("Init starts")
	signedString = "secret"
	auth = new(AuthResponse)
	config.LoadConfig()
	client := db.Connect()
	collection = client.Database(viper.GetString("MongoDBName")).Collection("user")
	mod := mongo.IndexModel{
		Keys:    bson.M{"emailId": 1, "id": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.TODO(), mod)
	if err != nil {
		log.Printf("Error creating index on field:[emailID] : %v\n", err)
	}
	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	if count == 0 {
		ad := &User{}
		ad.ID = primitive.NewObjectID()
		ad.EmailID = "admin@gmail.com"
		ad.Password = ""
		Hash, _ := bcrypt.GenerateFromPassword([]byte("Test@1234"), 5)
		ad.PasswordHash = Hash
		ad.Name = "admin"
		ad.CreatedDate = time.Now()
		ad.Active = true
		_, err := collection.InsertOne(context.TODO(), ad)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	s.Handle("/login", http.HandlerFunc(Login)).Methods("POST")
	s.Handle("/add", http.HandlerFunc(AddUser)).Methods("POST")
	s.Handle("/list", http.HandlerFunc(Listuser)).Methods("GET")
	s.Use(middleware)
	log.Println("lisening on 5050...")
	err := http.ListenAndServe(":5050", s)
	if err != nil {
		log.Error("Http Error :", err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user *User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	var result *User
	err = collection.FindOne(context.TODO(), bson.M{"emailId": user.EmailID}).Decode(&result)
	if err != nil {
		log.Println(err)
	}

	err = bcrypt.CompareHashAndPassword(result.PasswordHash, []byte(user.Password))
	if err == nil {
		json.NewEncoder(w).Encode(&auth)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user *User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Print(err)
	}
	user.ID = primitive.NewObjectID()
	user.CreatedDate = time.Now()
	user.Active = true
	Hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	user.PasswordHash = Hash
	user.Password = ""
	_, err = collection.InsertOne(context.TODO(), &user)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("User Added Successfully")
}
func Listuser(w http.ResponseWriter, r *http.Request) {
	users := new([]User)
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
	}
	cursor.All(context.TODO(), users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.EscapedPath()
		log.Println("path:", path)
		body, _ := ioutil.ReadAll(r.Body)
		var user *User
		if err := json.Unmarshal(body, &user); err != nil {
			log.Println("unmarshaling error", err)
		}
		var token = new(TokenClaims)
		if path == "/api/login" {
			getAuthResponse(user, auth)
			log.Printf("auth %+v", auth)
			if auth != nil && auth.SessionCookie != "" {
				r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				next.ServeHTTP(w, r)
			}
		} else if r.Header["Token"] != nil {
			token = GetKeyClaims(r.Header["Token"][0])
			if token != nil {
				r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Invalid Token")
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid Token")
		}
	})
}

func getAuthResponse(existUser *User, reply *AuthResponse) {
	expireToken := time.Now().Add(time.Minute * 60).Unix()
	claims := TokenClaims{
		UserEmail: existUser.EmailID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "jwt-auth",
			Subject:   "User",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, error := token.SignedString([]byte(signedString))
	if error != nil {
		log.Println("forming token error", error)
	}
	reply.SessionCookie = tokenString
	reply.ExpirationTime = expireToken
	reply.EmailID = existUser.EmailID

}

func GetKeyClaims(auth string) *TokenClaims {

	token, _ := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(signedString), nil
	})

	if token != nil {
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			var userClaims *TokenClaims
			if err := mapstructure.Decode(claims, &userClaims); err != nil {
				return nil
			}
			return userClaims
		}
	}
	return nil
}
