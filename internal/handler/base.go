package handler

import (
	"net/http"

	ap "github.com/go-ap/activitypub"
	"github.com/labstack/echo/v4"
)

// GET /
func (h *Handler) GetService(c echo.Context) error {
	service := ap.Actor{
		Context:      ap.ActivityBaseURI,
		ID:           h.baseURL,
		Type:         ap.ServiceType,
		Name:         ap.DefaultNaturalLanguageValue("Example Service"),
		Summary:      ap.DefaultNaturalLanguageValue("This is an example service."),
		AttributedTo: ap.IRI("https://github.com/ras0q"),
		Audience:     ap.ItemCollection{ap.PublicNS},
		URL:          h.baseURL,
		Inbox:        h.baseURL.AddPath("inbox"),
		Outbox:       h.baseURL.AddPath("outbox"),
		Streams: ap.ItemCollection{
			h.baseURL.AddPath("actors"),
			h.baseURL.AddPath("activities"),
			h.baseURL.AddPath("objects"),
		},
	}

	return c.JSON(http.StatusOK, service)
}

// GET /.well-known/webfinger?resource={resource}
func (h *Handler) GetWebFinger(c echo.Context) error {
	resource := c.QueryParam("resource")
	if resource == "" {
		return c.String(http.StatusBadRequest, "resource query parameter is required")
	}

	webFinger := map[string]interface{}{
		"subject": resource,
		"links": []map[string]interface{}{
			{
				"rel":  "self",
				"type": "application/activity+json",
				"href": h.baseURL.AddPath("actors", resource),
			},
		},
	}

	return c.JSON(http.StatusOK, webFinger)
}
