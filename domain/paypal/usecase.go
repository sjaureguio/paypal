package paypal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/google/uuid"
	"github.com/sjaureguio/paypal/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	ExpectedVerification = "SUCCESS"
	ExpectedStatus       = "completed"
)

const (
	EventTypeProduct      = "PAYMENT.CAPTURE.COMPLETED"
	EventTypeSubscription = "PAYMENT.SALE.COMPLETED"
)

type UseCase struct {
	useCaseOrder        Order
	useCaseSubscription Subscription
	useCaseInvoice      Invoice
}

func New(o Order, s Subscription, i Invoice) UseCase {
	return UseCase{
		useCaseOrder:        o,
		useCaseSubscription: s,
		useCaseInvoice:      i,
	}
}

func (uc UseCase) ProcessRequest(headers http.Header, body []byte) error {
	payPalRequest := parsePayPalRequest(headers, body)

	err := validate(&payPalRequest)
	if err != nil {
		return err
	}

	eventType, err := jsonparser.GetString(body, "event_type")
	if err != nil {
		return err
	}

	return uc.processPayment(eventType, &payPalRequest, body)
}

func (uc UseCase) processPayment(eventType string, req *models.PayPalRequest, body []byte) error {

	switch eventType {
	case EventTypeProduct:
		return uc.saleProduct(req, body)
	case EventTypeSubscription:
		return uc.saleSubscription(req, body)
	}

	log.Printf("the event type %q is not processed", eventType)

	return nil
}

func (uc UseCase) saleProduct(req *models.PayPalRequest, body []byte) error {
	var err error

	req.ID, err = jsonparser.GetString(body, "id")
	if err != nil {
		return err
	}

	req.ResourceID, err = jsonparser.GetString(body, "resource", "id")
	if err != nil {
		return err
	}

	req.Status, err = jsonparser.GetString(body, "resource", "status")
	if err != nil {
		return err
	}

	req.Custom, err = jsonparser.GetString(body, "resource", "custom")
	if err != nil {
		return err
	}

	req.Price, err = jsonparser.GetString(body, "resource", "amount", "total")
	if err != nil {
		return err
	}

	order, err := uc.useCaseOrder.FindByID(uuid.MustParse(req.Custom))
	if err != nil {
		return err
	}

	value, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		return err
	}

	if order.Price != value {
		return fmt.Errorf(
			"el valor recibido: %0.2f, es diferente al esperado %0.2f",
			value,
			order.Price,
		)
	}

	if !strings.EqualFold(req.Status, ExpectedStatus) {
		return fmt.Errorf(
			"el estado de la transaccion: %q no es el esperado: %q",
			req.Status,
			ExpectedStatus,
		)
	}

	subscription := models.Subscription{
		CustomerEmail: order.CustomerEmail,
		TypeSubs:      order.TypeSubs,
	}

	err = uc.useCaseSubscription.Create(&subscription)
	if err != nil {
		return err
	}

	return uc.useCaseInvoice.Create(&order, subscription.ID)
}

func (uc UseCase) saleSubscription(req *models.PayPalRequest, body []byte) error {
	var err error

	req.ID, err = jsonparser.GetString(body, "id")
	if err != nil {
		return err
	}

	req.ResourceID, err = jsonparser.GetString(body, "resource", "id")
	if err != nil {
		return err
	}

	req.Status, err = jsonparser.GetString(body, "resource", "status")
	if err != nil {
		return err
	}

	req.Custom, err = jsonparser.GetString(body, "resource", "custom_id")
	if err != nil {
		return err
	}

	req.Price, err = jsonparser.GetString(body, "resource", "amount", "value")
	if err != nil {
		return err
	}

	order, err := uc.useCaseOrder.FindByID(uuid.MustParse(req.Custom))
	if err != nil {
		return err
	}

	value, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		return err
	}

	if order.Price != value {
		return fmt.Errorf(
			"el valor recibido: %0.2f, es diferente al esperado %0.2f",
			value,
			order.Price,
		)
	}

	if !strings.EqualFold(req.Status, ExpectedStatus) {
		return fmt.Errorf(
			"el estado de la transaccion: %q no es el esperado: %q",
			req.Status,
			ExpectedStatus,
		)
	}

	return uc.useCaseInvoice.Create(&order, uuid.Nil)
}

func parsePayPalRequest(headers http.Header, body []byte) models.PayPalRequest {
	return models.PayPalRequest{
		AuthAlgo:         headers.Get("Paypal-Auth-Algo"),
		CertURL:          headers.Get("Paypal-Cert-Url"),
		TransmissionID:   headers.Get("Paypal-Transmission-Id"),
		TransmissionSig:  headers.Get("Paypal-Transmission-Sig"),
		TransmissionTime: headers.Get("Paypal-Transmission-Time"),
		WebhookEvent:     body,
		WebhookID:        os.Getenv("PAYPAL_WEBHOOK_ID"),
	}
}

func validate(p *models.PayPalRequest) error {
	data, err := json.Marshal(p)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("PAYPAL_VALIDATION_URL"),
		bytes.NewReader(data),
	)

	if err != nil {
		return err
	}

	request.Header.Set(
		"Content-Type",
		"application/json",
	)
	request.SetBasicAuth(
		os.Getenv("PAYPAL_CLIENT_ID"),
		os.Getenv("PAYPAL_SECRET_KEY"),
	)

	client := http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var body []byte

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"PayPal response with status code %d, body: %s",
			resp.StatusCode,
			string(body),
		)
	}

	bodyMap := make(map[string]string)

	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		return err
	}

	if bodyMap["verification_status"] != ExpectedVerification {
		return fmt.Errorf(
			"verification status is %s",
			bodyMap["verification_status"],
		)
	}

	return nil
}
