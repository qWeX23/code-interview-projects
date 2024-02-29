package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Patient struct {
	ID                    int                     `json:"id,omitempty"`
	FirstName             string                  `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName              string                  `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Gender                string                  `json:"gender,omitempty" bson:"gender,omitempty"`
	PhoneNumber           string                  `json:"phone_number,omitempty" bson:"phone_number,omitempty"`
	Email                 string                  `json:"email,omitempty" bson:"email,omitempty"`
	Address               string                  `json:"address,omitempty" bson:"address,omitempty"`
	VisitDate             string                  `json:"visit_date,omitempty" bson:"visit_date,omitempty"`
	Diagnosis             string                  `json:"diagnosis,omitempty" bson:"diagnosis,omitempty"`
	DrugCode              string                  `json:"drug_code,omitempty" bson:"drug_code,omitempty"`
	AdditionalInformation []AdditionalInformation `json:"additional_information,omitempty" bson:"additional_information,omitempty"`
}
type AdditionalInformation struct {
	Notes      string `json:"notes,omitempty"`
	NewPatient bool   `json:"new_patient,omitempty"`
	Race       string `json:"race,omitempty"`
	Ssn        string `json:"ssn,omitempty"`
}

type PatientService interface {
	GetPatientHandler(c *gin.Context)
	PostPatientHandler(c *gin.Context)
	PutPatientHandler(c *gin.Context)
	DeletePatientHandler(c *gin.Context)
	SearchPatientHandler(c *gin.Context)
}

type patientService struct {
	db Database
}

func NewPatientService(db Database) PatientService {
	return &patientService{
		db: db,
	}
}

func (p *patientService) GetPatientHandler(c *gin.Context) {
	//TODO add this to middleware
	patientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
	}
	patient, err := p.db.ReadPatient(c.Request.Context(), patientId)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding patient"})
	}
	c.JSON(http.StatusOK, patient)
}
func (p *patientService) PostPatientHandler(c *gin.Context) {
	var request = Patient{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	patient, err := p.db.CreatePatient(c.Request.Context(), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving patient"})
	}
	c.JSON(http.StatusCreated, patient)
}
func (p *patientService) PutPatientHandler(c *gin.Context) {
	//TODO add this to middleware
	patientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
	}
	var request = Patient{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	patient, err := p.db.UpdatePatient(c.Request.Context(), patientId, request)
	if err != nil {
		fmt.Printf("\n%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating patient"})
	}
	c.JSON(http.StatusOK, patient)

}
func (p *patientService) DeletePatientHandler(c *gin.Context) {
	//TODO add this to middleware
	patientId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be int"})
	}
	if err := p.db.DeletePatient(c.Request.Context(), patientId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting patient"})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (p *patientService) SearchPatientHandler(c *gin.Context) {
	//TODO figure out some validation rules here.
	ids := []int{}

	for k, v := range c.Request.URL.Query() {
		fmt.Printf("\n%v %v\n", k, v)
		i, err := p.db.Search(c.Request.Context(), k, v[0])
		if err != nil {
			fmt.Printf("%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching"})
		}
		ids = i
		break
	}

	c.JSON(http.StatusOK, ids)
}
