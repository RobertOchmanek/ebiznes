package api

import (
	"github.com/RobertOchmanek/ebiznes_go/model"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"net/http"
)

func HandlePayment(c echo.Context) error {

	var paymentDetails model.PaymentDetails
	c.Bind(&paymentDetails)

	apiKey := ""
	stripe.Key = apiKey

	_, err := charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(paymentDetails.Amount),
		Currency:     stripe.String(string(stripe.CurrencyUSD)),
		Description:  stripe.String(paymentDetails.Description),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String(paymentDetails.PayerEmail)})

	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.JSON(http.StatusOK, "Zapłacone elegancko")

	//TODO: save payment to DB
}
