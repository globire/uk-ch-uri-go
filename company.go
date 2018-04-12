package uri

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const url = "http://data.companieshouse.gov.uk/doc/company/%s.json"

// ChDate is type which supports unmarshalling from CH json response to a Go time type
type ChDate struct {
	time.Time
}

// UnmarshalJSON implements the unmarshalling functionality
func (cd *ChDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if len(s) == 0 {
		return
	}
	cd.Time, err = time.Parse("02/01/2006", s)
	return
}

type strint int

func (v *strint) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*v = strint(i)
	return nil
}

func (v *strint) Int() int {
	return int(*v)
}

func (v strint) String() string {
	return strconv.Itoa(v.Int())
}

type Company struct {
	Name             string `json:"CompanyName"`
	CompanyNumber    string `json:"CompanyNumber"`
	RegisteredOffice struct {
		CareOf       string `json:"Careof"`
		POBox        string `json:"POBox"`
		AddressLine1 string `json:"AddressLine1"`
		AddressLine2 string `json:"AddressLine2"`
		PostTown     string `json:"PostTown"`
		County       string `json:"County"`
		Country      string `json:"Country"`
		Postcode     string `json:"Postcode"`
	} `json:"RegAddress"`
	CompanyCategory   string `json:"CompanyCategory"`
	CompantStatus     string `json:"CompanyStatus"`
	CountryOfOrigin   string `json:"CountryofOrigin"`
	IncorporationDate ChDate `json:"IncorporationDate"`
	RegistrationDate  ChDate `json:"RegistrationDate"`
	DissolutionDate   ChDate `json:"DissolutionDate"`
	PreviousNames     []struct {
		CONDate     ChDate `json:"CONDate"`
		CompanyName string
	} `json:"PreviousName"`
	Accounts struct {
		AccountRefDay    strint `json:"AccountRefDay"`
		AccountRefMonth  strint `json:"AccountRefMonth"`
		NextDueDate      ChDate `json:"NextDueDate"`
		LastMadeUpDate   ChDate `json:"LastMadeUpDate"`
		AccountsCategory string `json:"AccountsCategory"`
	} `json:"Accounts"`
	Returns struct {
		NextDueDate    ChDate `json:"NextDueDate"`
		LastMadeUpDate ChDate `json:"LastMadeUpDate"`
	} `json:"Returns"`
	Mortgages struct {
		Charges       strint `json:"NumMortCharges"`
		Outstanding   strint `json:"NumMortOutstanding"`
		PartSatisfied strint `json:"NumMortPartSatisfied"`
		Satisfied     strint `json:"NumMortSatisfied"`
	} `json:"Mortgages"`
	SICCodes struct {
		Text []string `json:"SicText"`
	} `json:"SICCodes"`
	LimtitedPartnerships struct {
		GeneralPartners strint `json:"NumGenPartners"`
		NumLimPartners  strint `json:"NumLimPartners"`
	}
}

// HasTasks returns true if any statutory task is overdue
func (c Company) HasTasks() bool {
	now := time.Now()
	return now.Sub(c.Accounts.NextDueDate.Time) < 0 || now.Sub(c.Returns.NextDueDate.Time) < 0
}

type chResponse struct {
	PrimaryTopic Company `json:"primaryTopic"`
}

// GetCompany fetches data from Companies House and returns a Company
func GetCompany(companyNumber string) (*Company, error) {
	a, err := http.Get(fmt.Sprintf(url, companyNumber))
	if err != nil {
		return nil, err
	}
	defer a.Body.Close()
	res := chResponse{}
	if err := json.NewDecoder(a.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res.PrimaryTopic, nil
}
