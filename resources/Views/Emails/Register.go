package emails

import (
	"github.com/labstack/echo/v4"
	"github.com/matcornic/hermes/v2"
)

func Register(name string, c echo.Context) string {
	h := hermes.Hermes{
		Product: hermes.Product{
			Copyright:   "Setup",
			TroubleText: "Setup",
			Name:        "Setup",
			Link:        "https://github.com/akifkadioglu/go-echo-gorm-mysql-setup",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{

			Greeting: "Greeting",
			Name:     "name",
			Intros: []string{
				"description 1",
				"description 2",
			},
			Actions: []hermes.Action{
				{
					Instructions: "description 3",
					Button: hermes.Button{
						Color: "#357EC7",
						Text:  "Confirm button",
						Link:  "https://github.com/akifkadioglu/go-echo-gorm-mysql-setup",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
			Signature: "Akif Kadıoğlu",
		},
	}
	emailBody, _ := h.GenerateHTML(email)
	return emailBody
}
