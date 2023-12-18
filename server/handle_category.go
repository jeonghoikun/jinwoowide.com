package server

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jeonghoikun/hamjayoung.com/site"
	"github.com/jeonghoikun/hamjayoung.com/store"
)

type categoryHandler struct{}

// GET /category/:do/:si/:storeType
func (*categoryHandler) listPage(c *fiber.Ctx) error {
	do, err := url.QueryUnescape(c.Params("do"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	si, err := url.QueryUnescape(c.Params("si"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	storeType, err := url.QueryUnescape(c.Params("storeType"))
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}
	listStores := store.ListStoresByDoSiAndStoreType(do, si, storeType)
	if len(listStores) == 0 {
		return c.Status(http.StatusNotFound).SendString("카테고리가 존재하지 않습니다")
	}
	sort.Slice(listStores, func(i, j int) bool {
		return listStores[i].DatePublished.UnixNano() > listStores[j].DatePublished.UnixNano()
	})
	var storeNames []string
	for _, s := range listStores {
		storeNames = append(storeNames, s.Title)
	}
	si = strings.Replace(si, "구", "", -1)
	m := fiber.Map{}
	m["Page"] = &PageConfig{
		Path: c.Path(),
		Author: &Author{
			Name:        site.Config.Author,
			ProfilePath: "/static/img/site/author/profile.png",
		},
		Title: fmt.Sprintf("[%s > %s > %s] 업소 목록", do, si, storeType),
		Description: fmt.Sprintf("%s %s 지역에 %d개의 %s 업소가 있습니다: %s",
			do, si, len(listStores), storeType, strings.Join(storeNames, ", ")),
		Keywords: strings.Join(
			[]string{fmt.Sprintf("%s %s %s 업소 목록", do, si, storeType)},
			",",
		),
		PhoneNumber:   site.Config.PhoneNumber,
		DatePublished: site.Config.DatePublished,
		DateModified:  site.Config.DateModified,
		ThumbnailPath: "/static/img/site/thumbnail/thumb.png",
	}
	m["Profile"] = map[string]string{"PhoneNumber": listStores[0].PhoneNumber}
	m["Breadcrumbs"] = map[string]string{"StoreType": listStores[0].Type}
	m["Stores"] = listStores
	return c.Status(http.StatusOK).Render("category/index", m, "layout/category")
}

// BaseURL = /category
func handleCategory(r fiber.Router) {
	h := &categoryHandler{}
	r.Get("/:do/:si/:storeType", h.listPage)
}
