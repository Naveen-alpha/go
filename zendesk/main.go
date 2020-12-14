package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	config "./config"
)

type Zendesk struct {
	User   User   `json:"user,omitempty" bson:"user,omitempty"`
	Ticket Ticket `json:"ticket,omitempty" bson:"ticket,omitempty"`
}
type Response struct {
	Results      []Results   `json:"results,omitempty" bson:"results,omitempty"`
	Facets       interface{} `json:"facets,omitempty" bson:"facets,omitempty"`
	NextPage     interface{} `json:"next_page,omitempty" bson:"next_page,omitempty"`
	PreviousPage interface{} `json:"previous_page,omitempty" bson:"previous_page,omitempty"`
	Count        int         `json:"count,omitempty" bson:"count,omitempty"`
}
type Results struct {
	ID       int64  `json:"id,omitempty" bson:"id,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	MemberID string `json:"external_id,omitempty" bson:"external_id,omitempty"`
}
type User struct {
	ID       int64  `json:"id,omitempty" bson:"id,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Phone    string `json:"phone,omitempty" bson:"phone,omitempty"`
	MemberID string `json:"external_id,omitempty" bson:"external_id,omitempty"`
}

func main() {
	T := &Zendesk{}
	//GetTickets()
	T.GetUserId("account_54325")
	//T.CreateUser("bakkarit@gmail.com", "abu", "104", "8754553265")
	//T.CreateTicket("103", "400018010894", "1005")
}

type Ticket struct {
	ID                 int64              `json:"id,omitempty" bson:"id,omitempty"`
	URL                string             `json:"url,omitempty" bson:"url,omitempty"`
	ExternalID         string             `json:"external_id,omitempty" bson:"external_id,omitempty"`
	CreatedAt          string             `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt          string             `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Type               string             `json:"type,omitempty" bson:"type,omitempty"`
	Subject            string             `json:"subject,omitempty" bson:"subject,omitempty"`
	RawSubject         string             `json:"raw_subject,omitempty" bson:"raw_subject,omitempty"`
	Description        string             `json:"description,omitempty" bson:"description,omitempty"`
	Priority           string             `json:"priority,omitempty" bson:"priority,omitempty"`
	Status             string             `json:"status,omitempty" bson:"status,omitempty"`
	Recipient          string             `json:"recipient,omitempty" bson:"recipient,omitempty"`
	RequesterID        int                `json:"requester_id,omitempty" bson:"requester_id,omitempty"`
	SubmitterID        int64              `json:"submitter_id,omitempty" bson:"submitter_id,omitempty"`
	AssigneeID         int64              `json:"assignee_id,omitempty" bson:"assignee_id,omitempty"`
	OrganizationID     int64              `json:"organization_id,omitempty" bson:"organization_id,omitempty"`
	GroupID            int64              `json:"group_id,omitempty" bson:"group_id,omitempty"`
	CollaboratorID     []int64            `json:"collaborator_ids,omitempty" bson:"collaborator_ids,omitempty"`
	FollowerID         []int64            `json:"follower_ids,omitempty" bson:"follower_ids,omitempty"`
	ProblemID          int64              `json:"problem_id,omitempty" bson:"problem_id,omitempty"`
	HasIncidents       bool               `json:"has_incidents,omitempty" bson:"has_incidents,omitempty"`
	DueAt              time.Time          `json:"due_at,omitempty" bson:"due_at,omitempty"`
	Tags               []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Via                Via                `json:"via,omitempty" bson:"via,omitempty"`
	CustomFields       []CustomFields     `json:"custom_fields,omitempty" bson:"custom_fields,omitempty"`
	SatisfactionRating SatisfactionRating `json:"satisfaction_rating,omitempty" bson:"satisfaction_rating,omitempty"`
	SharingAgreementID []int64            `json:"sharing_agreement_ids,omitempty" bson:"sharing_agreement_ids,omitempty"`
	Comment            Comment            `json:"comment,omitempty" bson:"comment,omitempty"`
	ApplicationNumber  string             `json:"application_number,omitempty" bson:"application_number,omitempty"`
}
type Comment struct {
	Body string `json:"body,omitempty" bson:"body,omitempty"`
}
type DueAt struct {
	Some string `json:"some,omitempty" bson:"some,omitempty"`
}
type Via struct {
	Channel string `json:"channel,omitempty" bson:"channel,omitempty"`
}
type CustomFields struct {
	ID    int    `json:"id,omitempty" bson:"id,omitempty"`
	Value string `json:"value,omitempty" bson:"value,omitempty"`
}
type SatisfactionRating struct {
	ID      int64  `json:"id,omitempty" bson:"id,omitempty"`
	Score   string `json:"score,omitempty" bson:"score,omitempty"`
	Comment string `json:"comment,omitempty" bson:"comment,omitempty"`
}

func GetTickets() {
	log.Println("Get Tickets Called.....")
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.CreateTicketURL, nil)
	authString := "Basic " + base64.StdEncoding.EncodeToString([]byte(config.UserName+":"+config.Password))
	//var authString = "Basic Yi5uYXZlZW5rdW1hckBiYW5rY2J3Lm9yZzpDYnduYXZlZW4x"

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authString)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	var t *[]Ticket
	json.Unmarshal(bodyText, &t)
	fmt.Printf("Tickets %+v :", &t)
	log.Println(s)

}
func (t *Zendesk) CreateUser(email string, name string, memberId string, phone string) string {
	log.Println(" CreateUser Called.....")

	client := &http.Client{}
	//user := &Zendesk{name, email, phone, memberId}

	user := &Zendesk{
		User: User{
			Name:     name,
			Email:    email,
			MemberID: memberId,
			Phone:    phone,
		},
	}
	log.Println("user details ", user)
	userbt, er := json.Marshal(user)
	if er != nil {
		log.Printf("Cannot Marshal %s", er)
	}

	//req, err := http.NewRequest("POST", config.UserUrl, strings.NewReader(v.Encode()))
	req, err := http.NewRequest("POST", config.CreateUserURL, bytes.NewBuffer(userbt))

	//authString := "Basic " + base64.StdEncoding.EncodeToString([]byte(config.UserName+":"+config.Password))
	var authString = "Basic Yi5uYXZlZW5rdW1hckBiYW5rY2J3Lm9yZzpDYnduYXZlZW4x"

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authString)
	//req.SetBasicAuth(config.UserName, config.Password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	rt, _ := ioutil.ReadAll(resp.Body)
	var u Zendesk
	errr := json.Unmarshal(rt, &u)
	if errr != nil {
		log.Println("Error on unmarshaling", errr)
	}
	log.Println("User Id (ZenId):", u.User.ID)
	resp.Body.Close()
	log.Println(string(rt))
	return string(rt)
}

func (t *Zendesk) GetUserId(memberId string) (userId string) {
	log.Println("Get UserId Called.....")
	client := &http.Client{}

	req, err := http.NewRequest("GET", config.GetExIDURL+memberId, nil)
	authString := "Basic " + base64.StdEncoding.EncodeToString([]byte(config.UserName+":"+config.Password))
	//var authString = "Basic Yi5uYXZlZW5rdW1hckBiYW5rY2J3Lm9yZzpDYnduYXZlZW4x"
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authString)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	rt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error on read", err)
	}
	log.Println("RESPONSE :", string(rt))
	var test Response
	errr := json.Unmarshal(rt, &test)
	if errr != nil {
		log.Println("Error on unmarshaling", errr)
	}
	defer resp.Body.Close()

	userID := test.Results[0].ID
	log.Println("\n User Id (ZenId):", userID)
	return string(userID)
}

func (t *Zendesk) CreateTicket(memberId string, userId string, applicationNumber string) string {
	client := &http.Client{}
	requesterID, err := strconv.Atoi(userId)
	if err != nil {
		log.Println("Error converting userId to int ", err)
	}
	ticket := Zendesk{
		Ticket: Ticket{
			Subject:     "Customer " + memberId + " Remittance" + applicationNumber,
			RequesterID: requesterID,
			ExternalID:  memberId,
			Tags:        []string{"CustomerVerification"},
			CustomFields: []CustomFields{
				CustomFields{
					ID:    config.MemberID,
					Value: memberId,
				},
				CustomFields{
					ID:    config.ApplicationID,
					Value: applicationNumber,
				},
			},
			Comment: Comment{
				Body: "Review remittance " + applicationNumber},
		},
	}
	tktbt, err := json.Marshal(&ticket)
	req, err := http.NewRequest("POST", config.CreateTicketURL, bytes.NewBuffer(tktbt))
	authString := "Basic " + base64.StdEncoding.EncodeToString([]byte(config.UserName+":"+config.Password))

	req.Header.Add("Authorization", authString)
	req.Header.Add("Content-Type", "application/json")
	log.Println("req: ", req)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	log.Println("resp: ", resp)
	rt, _ := ioutil.ReadAll(resp.Body)
	// var u Zendesk
	// errr := json.Unmarshal(rt, &u)
	// if errr != nil {
	// 	log.Println("Error on unmarshaling", errr)
	// }
	//log.Println("User Id (ZenId):", u.User.ID)
	resp.Body.Close()
	log.Println(string(rt))
	return string(rt)

}

/*
{
	"id":               35436,
	"url":              "https://company.zendesk.com/api/v2/tickets/35436.json",
	"external_id":      "ahg35h3jh",
	"created_at":       "2009-07-20T22:55:29Z",
	"updated_at":       "2011-05-05T10:38:52Z",
	"type":             "incident",
	"subject":          "Help, my printer is on fire!",
	"raw_subject":      "{{dc.printer_on_fire}}",
	"description":      "The fire is very colorful.",
	"priority":         "high",
	"status":           "open",
	"recipient":        "support@company.com",
	"requester_id":     20978392,
	"submitter_id":     76872,
	"assignee_id":      235323,
	"organization_id":  509974,
	"group_id":         93998617821148738,
	"collaborator_ids": [35334, 234],
	"follower_ids":     [35334, 234], // This functionally is the same as collaborators for now.
	"problem_id":       9873764,
	"has_incidents":    false,
	"due_at":           null,
	"tags":             ["enterprise", "other_tag"],
	"via": {
	  "channel": "web"
	},
	"custom_fields": [
	  {
		"id":    27642,
		"value": "745"
	  },
	  {
		"id":    27648,
		"value": "yes"
	  }
	],
	"satisfaction_rating": {
	  "id": 1234,
	  "score": "good",
	  "comment": "Great support!"
	},
	"sharing_agreement_ids": [84432]
  }*/
/*static def createTicketFromTransaction(RemittanceNewCustomer remittanceNewCustomer) {
def memberId = remittanceNewCustomer.memberId
def applicationNumber = remittanceNewCustomer.applicationNumber
def csrTicket = new CsrTicket()
csrTicket.requester_id = remittanceNewCustomer.csrUserId
csrTicket.subject = "Customer ${memberId} Remittance ${applicationNumber}"
csrTicket.external_id = remittanceNewCustomer.id
csrTicket.tags = ["customerverification"]
csrTicket.custom_fields = [
new CsrCustomField(id:RemoteServerConfig.csrFieldMemberId, value:memberId),
new CsrCustomField(id:RemoteServerConfig.csrFieldApplicationNumber, value:applicationNumber)
]
csrTicket.comment = new CsrComment(body:"Review remittance ${applicationNumber} originated by new customer")
def responseObject = CsrApiService.createTicket(csrTicket)
}*/
