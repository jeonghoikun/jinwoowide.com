package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonghoikun/hamjayoung.com/site"
	"github.com/jeonghoikun/hamjayoung.com/store"
)

type storeHandler struct{}

// GET /store/:do/:si/:dong/:type/:title
func (*storeHandler) page(c *fiber.Ctx) error {
	do, err := url.QueryUnescape(c.Params("do"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	si, err := url.QueryUnescape(c.Params("si"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	dong, err := url.QueryUnescape(c.Params("dong"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	storeType, err := url.QueryUnescape(c.Params("type"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	storeTitle, err := url.QueryUnescape(c.Params("title"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	store, has := store.Get(do, si, dong, storeType, storeTitle)
	if !has {
		return c.Status(http.StatusNotFound).SendString("Store not found")
	}
	si = strings.Replace(si, "구", "", -1)
	title := fmt.Sprintf("%s %s %s", si, store.Title, store.Type)
	if store.Active.IsPermanentClosed {
		title += fmt.Sprintf(" (폐업: %s)", store.Active.Reason)
	} else {
		title += " (영업중)"
	}
	m := fiber.Map{
		"Page": &PageConfig{
			Path: c.Path(),
			Author: &Author{
				Name:        site.Config.Author,
				ProfilePath: "/static/img/site/author/profile.png",
			},
			Title:         title,
			Description:   store.Description,
			Keywords:      store.Keywords.String(),
			PhoneNumber:   store.PhoneNumber,
			DatePublished: store.DatePublished,
			DateModified:  store.DateModified,
			ThumbnailPath: fmt.Sprintf("/static/img/store/%s/%s/%s/%s/%s/thumbnail.png",
				store.Location.Do, store.Location.Si, store.Location.Dong, store.Type, store.Title),
		},
		"Profile": map[string]string{
			"PhoneNumber": store.PhoneNumber,
		},
		"Store":  store,
		"SiMini": si,
	}
	embedFilePath := fmt.Sprintf("store/%s/%s/%s/%s/%s",
		store.Location.Do, store.Location.Si, store.Location.Dong, store.Type, store.Title)
	return c.Status(http.StatusOK).Render(embedFilePath, m, "layout/store")
}

// BaseURL = /store
func handleStore(r fiber.Router) {
	h := &storeHandler{}
	r.Get("/:do/:si/:dong/:type/:title", h.page)
}
