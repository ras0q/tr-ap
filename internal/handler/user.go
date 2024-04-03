package handler

import (
	"net/http"

	ap "github.com/go-ap/activitypub"
	"github.com/labstack/echo/v4"
)

// GET /u
func (h *Handler) GetUsers(_ echo.Context) error {
	return nil
}

// GET /u/:userID
func (h *Handler) GetUser(c echo.Context) error {
	userID := c.Param("userID")
	userIRI := h.baseURL.AddPath("u", userID)
	user := ap.Actor{
		ID:        userIRI,
		Type:      ap.PersonType,
		Inbox:     userIRI.AddPath("inbox"),
		Outbox:    userIRI.AddPath("outbox"),
		Following: userIRI.AddPath("following"),
		Followers: userIRI.AddPath("followers"),
		Liked:     userIRI.AddPath("liked"),
	}

	return c.JSON(http.StatusOK, user)
}

// GET /u/:userID/inbox
func (h *Handler) GetUserInbox(c echo.Context) error {
	userID := c.Param("userID")
	userIRI := h.baseURL.AddPath("u", userID)
	inbox := ap.OrderedCollection{
		ID:   userIRI.AddPath("inbox"),
		Type: ap.OrderedCollectionType,
	}

	return c.JSON(http.StatusOK, inbox)
}

// GET /u/:userID/outbox
func (h *Handler) GetUserOutbox(c echo.Context) error {
	userID := c.Param("userID")
	userIRI := h.baseURL.AddPath("u", userID)
	outbox := ap.OrderedCollection{
		ID:   userIRI.AddPath("outbox"),
		Type: ap.OrderedCollectionType,
	}

	return c.JSON(http.StatusOK, outbox)
}

// GET /u/:userID/following
func (h *Handler) GetUserFollowing(c echo.Context) error {
	userID := c.Param("userID")
	userIRI := h.baseURL.AddPath("u", userID)
	following := ap.Collection{
		ID:   userIRI.AddPath("following"),
		Type: ap.CollectionType,
	}

	return c.JSON(http.StatusOK, following)
}

// GET /u/:userID/followers
func (h *Handler) GetUserFollowers(c echo.Context) error {
	userID := c.Param("userID")
	userIRI := h.baseURL.AddPath("u", userID)
	followers := ap.Collection{
		ID:   userIRI.AddPath("followers"),
		Type: ap.CollectionType,
	}

	return c.JSON(http.StatusOK, followers)
}

// GET /u/:userID/liked
func (h *Handler) GetUserLiked(c echo.Context) error {
	userID := c.Param("userID")
	userIRI := h.baseURL.AddPath("u", userID)
	liked := ap.Collection{
		ID:   userIRI.AddPath("liked"),
		Type: ap.CollectionType,
	}

	return c.JSON(http.StatusOK, liked)
}
