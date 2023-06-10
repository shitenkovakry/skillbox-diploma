package alldata

import (
	"encoding/json"
	"log"
	"net/http"
	"skillbox-diploma/models"
)

type ResultT struct {
	Status bool        `json:"status"` // true, если все этапы сбора данных прошли успешно, false во всех остальных случаях
	Data   *ResultSetT `json:"data"`   // заполнен, если все этапы сбора данных прошли успешно, nil во всех остальных случаях
	Error  string      `json:"error"`  // пустая строка если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки (детали ниже)
}

type ResultSetT struct {
	SMS       []models.SMSData     `json:"sms"`
	MMS       []models.MMSData     `json:"mms"`
	VoiceCall models.VoiceCallData `json:"voice_call"`
	Email     []models.EmailData   `json:"email"`
	Billing   *models.BillingDatum `json:"billing"`
	Support   []int                `json:"support"`
	Incidents models.IncidentData  `json:"incident"`
}

type SourceSMS interface {
	Read() models.SMSData
}

type SourceMMS interface {
	Read() models.MMSData
}

type SourceVoiceCall interface {
	Read() models.VoiceCallData
}

type SourceEmail interface {
	Read() models.EmailData
}

type SourceBilling interface {
	Read() *models.BillingDatum
}

type SourceIncidents interface {
	Read() models.IncidentData
}

type SourceSupport interface {
	Read() models.SupportData
}

type Handler struct {
	sms       SourceSMS
	mms       SourceMMS
	voiceCall SourceVoiceCall
	email     SourceEmail
	billing   SourceBilling
	incidents SourceIncidents
	support   SourceSupport
}

func (handler *Handler) sendResponse(writer http.ResponseWriter, result *ResultT) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Print("data, err := json.Marshal(result)")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := writer.Write(data); err != nil {
		log.Print("_, err := writer.Write(data)")
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	result := &ResultT{
		Status: false,
		Data:   &ResultSetT{},
		Error:  "Error on collect data",
	}

	result.Data.SMS = handler.obtainSMSData()
	result.Data.MMS = handler.obtainMMSData()
	result.Data.VoiceCall = handler.voiceCall.Read()
	result.Data.Email = handler.obtainEmailData()
	result.Data.Billing = handler.billing.Read()
	result.Data.Incidents = handler.obtainIncidentData()
	result.Data.Support = handler.obtainSupportData()

	if result.Data.SMS != nil &&
		result.Data.MMS != nil &&
		result.Data.VoiceCall != nil &&
		result.Data.Email != nil &&
		result.Data.Billing != nil &&
		result.Data.Incidents != nil &&
		result.Data.Support != nil {
		result.Status = true
		result.Error = ""
	} else {
		result.Data = nil
	}

	handler.sendResponse(writer, result)
}

func New(sms SourceSMS, mms SourceMMS, voiceCall SourceVoiceCall, email SourceEmail, billing SourceBilling, incidents SourceIncidents, support SourceSupport) *Handler {
	return &Handler{
		sms:       sms,
		mms:       mms,
		voiceCall: voiceCall,
		email:     email,
		billing:   billing,
		incidents: incidents,
		support:   support,
	}
}
