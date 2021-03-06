package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/cms-orbits/cms-sao/app"
	"github.com/cms-orbits/cms-sao/storage"
)

// DraftSubmitTrxController implements the DraftSubmitTrx resource.
type DraftSubmitTrxController struct {
	*goa.Controller
	repository storage.EntrySubmitTrxRepository
}

// NewDraftSubmitTrxController creates a DraftSubmitTrx controller.
func NewDraftSubmitTrxController(service *goa.Service, store storage.EntrySubmitTrxRepository) *DraftSubmitTrxController {
	return &DraftSubmitTrxController{
		Controller: service.NewController("DraftSubmitTrxController"),
		repository: store,
	}
}

// Get runs the get action.
func (c *DraftSubmitTrxController) Get(ctx *app.GetDraftSubmitTrxContext) error {
	entryTrx, err := c.repository.FindByID(ctx, ctx.TrxID)

	if err != nil {
		return ctx.NotFound()
	}

	res := &app.ComJossemargtSaoDraftSubmitTransactionFull{
		ID:        entryTrx.ID,
		Status:    entryTrx.Status,
		UpdatedAt: &entryTrx.UpdatedAt,
		CreatedAt: &entryTrx.CreatedAt,
		Href:      fmt.Sprintf("%s%s", app.DraftSubmitTrxHref(), entryTrx.ID),
		Links:     new(app.ComJossemargtSaoDraftSubmitTransactionLinks),
	}

	if entryTrx.EntryID != 0 {
		res.Links.Draft = &app.ComJossemargtSaoDraftLink{
			ID:   entryTrx.EntryID,
			Href: fmt.Sprintf("%s%d", app.DraftHref(), entryTrx.EntryID),
		}
	} else {
		res.Links = nil
	}

	return ctx.OKFull(res)
}

// Show runs the show action.
func (c *DraftSubmitTrxController) Show(ctx *app.ShowDraftSubmitTrxContext) error {
	return ctx.NotImplemented()
}
