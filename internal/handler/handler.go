package handler

import (
	"tr-ap/internal/repository"

	ap "github.com/go-ap/activitypub"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	baseURL ap.IRI
	repo    *repository.Repository
}

func New(baseURL ap.IRI, repo *repository.Repository) *Handler {
	return &Handler{
		baseURL: baseURL,
		repo:    repo,
	}
}

func (h *Handler) SetupRoutes(e *echo.Echo) {
	activityJSONMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, "application/activity+json")

			return next(c)
		}
	}

	e.GET("/", h.GetService, activityJSONMiddleware)
	e.GET("/.well-known/host-meta", h.GetHostMeta)
	e.GET("/.well-known/webfinger", h.GetWebFinger)

	e.GET("/ping", h.Ping)

	e.GET("/u", h.GetUsers, activityJSONMiddleware)
	e.GET("/u/:userID", h.GetUser, activityJSONMiddleware)
	e.GET("/u/:userID/inbox", h.GetUserInbox, activityJSONMiddleware)
	e.GET("/u/:userID/outbox", h.GetUserOutbox, activityJSONMiddleware)
	e.GET("/u/:userID/following", h.GetUserFollowing, activityJSONMiddleware)
	e.GET("/u/:userID/followers", h.GetUserFollowers, activityJSONMiddleware)
	e.GET("/u/:userID/liked", h.GetUserLiked, activityJSONMiddleware)
}
