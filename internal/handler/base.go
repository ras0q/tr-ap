package handler

import (
	"fmt"
	"net/http"
	"strings"

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

// GET /.well-known/host-meta
func (h *Handler) GetHostMeta(c echo.Context) error {
	hostMeta := fmt.Sprintf(`<?xml version="1.0"?>
<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
	<Link rel="lrdd" type="application/xrd+xml" template="%s/.well-known/webfinger?resource={uri}" />
</XRD>`,
		h.baseURL,
	)

	return c.XML(http.StatusOK, hostMeta)
}


// GET /.well-known/webfinger?resource={resource}
func (h *Handler) GetWebFinger(c echo.Context) error {
	resource := c.QueryParam("resource")
	if resource == "" {
		return c.String(http.StatusBadRequest, "resource query parameter is required")
	}

	userID := strings.Split(strings.TrimPrefix(resource, "acct:"), "@")[0]
	webFinger := map[string]interface{}{
		"subject": resource,
		"links": []map[string]interface{}{
			{
				"rel":  "self",
				"type": "application/activity+json",
				"href": h.baseURL.AddPath("u", userID),
			},
		},
	}

	return c.JSON(http.StatusOK, webFinger)
}
