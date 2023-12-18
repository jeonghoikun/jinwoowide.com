package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonghoikun/gnfull.com/site"
	"github.com/jeonghoikun/gnfull.com/store"
)

type indexHandler struct{}

// GET /
func (*indexHandler) index(c *fiber.Ctx) error {
	m := fiber.Map{
		"Page": &PageConfig{
			Path: c.Path(),
			Author: &Author{
				Name:        site.Config.Author,
				ProfilePath: "/static/img/site/author/profile.png",
			},
			Title:         site.Config.Title,
			Description:   site.Config.Description,
			Keywords:      site.Config.Keywords.String(),
			PhoneNumber:   site.Config.PhoneNumber,
			DatePublished: site.Config.DatePublished,
			DateModified:  site.Config.DateModified,
			ThumbnailPath: "/static/img/site/thumbnail/thumb.png",
		},
		"Profile": map[string]string{
			"PhoneNumber": site.Config.PhoneNumber,
		},
	}
	return c.Status(http.StatusOK).Render("index", m, "layout/index")
}

// GET /robots.txt
func (*indexHandler) robots(c *fiber.Ctx) error {
	var ss []string
	ss = append(ss, "User-agent: *")
	ss = append(ss, "Allow: /")
	ss = append(ss, fmt.Sprintf("Sitemap: https://%s/sitemap.xml", site.Config.Domain))
	return c.Status(http.StatusOK).SendString(strings.Join(ss, "\n"))
}

// GET /sitemap.xml
func (*indexHandler) sitemap(c *fiber.Ctx) error {
	var ss []string
	ss = append(ss, `<?xml version="1.0" encoding="UTF-8"?>`)
	ss = append(ss, `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)

	host := fmt.Sprintf("https://%s", site.Config.Domain)
	dateModified := site.Config.DateModified.Format(time.RFC3339)

	// index
	ss = append(ss, `<url>`)
	ss = append(ss, fmt.Sprintf(`<loc>%s/</loc>`, host))
	ss = append(ss, fmt.Sprintf(`<lastmod>%s</lastmod>`, dateModified))
	ss = append(ss, `</url>`)

	// Custom: Categories by store type in Gangnam-gu, Seoul
	categories := []string{}
	for _, s := range store.ListAllStores() {
		do := url.QueryEscape(s.Location.Do)
		si := url.QueryEscape(s.Location.Si)
		storeType := url.QueryEscape(s.Type)
		var has bool
		for _, category := range categories {
			if category == fmt.Sprintf("%s:%s:%s", do, si, storeType) {
				has = true
				break
			}
		}
		if has {
			continue
		}
		ss = append(ss, `<url>`)
		ss = append(ss, fmt.Sprintf(`<loc>%s/category/%s/%s/%s</loc>`, host, do, si, storeType))
		dateModified = s.DateModified.Format(time.RFC3339)
		ss = append(ss, fmt.Sprintf(`<lastmod>%s</lastmod>`, dateModified))
		ss = append(ss, `</url>`)
		categories = append(categories, fmt.Sprintf("%s:%s:%s", do, si, storeType))
	}

	// stores
	for _, s := range store.ListAllStores() {
		ss = append(ss, `<url>`)
		do := url.QueryEscape(s.Location.Do)
		si := url.QueryEscape(s.Location.Si)
		dong := url.QueryEscape(s.Location.Dong)
		storeType := url.QueryEscape(s.Type)
		storeTitle := url.QueryEscape(s.Title)
		ss = append(ss, fmt.Sprintf(`<loc>%s/store/%s/%s/%s/%s/%s</loc>`, host, do, si, dong, storeType, storeTitle))
		dateModified = s.DateModified.Format(time.RFC3339)
		ss = append(ss, fmt.Sprintf(`<lastmod>%s</lastmod>`, dateModified))
		ss = append(ss, `</url>`)
	}

	ss = append(ss, `</urlset>`)
	c.Response().Header.Add("Content-Type", "application/xml")
	return c.Status(http.StatusOK).SendString(strings.Join(ss, "\n"))
}

// BaseURL = /
func handleIndex(r fiber.Router) {
	h := &indexHandler{}
	r.Get("/", h.index)
	r.Get("/robots.txt", h.robots)
	r.Get("/sitemap.xml", h.sitemap)
}
