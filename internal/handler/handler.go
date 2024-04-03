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
	e.GET("/", h.GetService)
	e.GET("/.well-known/webfinger", h.GetWebFinger)

	e.GET("/ping", h.Ping)

	e.GET("/u", h.GetUsers)
	e.GET("/u/:userID", h.GetUser)
	e.GET("/u/:userID/inbox", h.GetUserInbox)
	e.GET("/u/:userID/outbox", h.GetUserOutbox)
	e.GET("/u/:userID/following", h.GetUserFollowing)
	e.GET("/u/:userID/followers", h.GetUserFollowers)
	e.GET("/u/:userID/liked", h.GetUserLiked)
}
