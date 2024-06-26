package v1_summary

import (
	"net/http"
	"too-lazy-to-watch-api/routes"
	custom_error "too-lazy-to-watch-api/src/error"
	"too-lazy-to-watch-api/src/summary"

	"github.com/labstack/echo/v4"
)

type handler struct {
	summaryService summary.ISummaryService
}

func NewHandler(g *echo.Group, summaryService summary.ISummaryService) {
	h := &handler{
		summaryService: summaryService,
	}

	g.POST("/create-from-youtube", h.CreateFromYoutube)
}

func (h *handler) CreateFromYoutube(c echo.Context) error {
	payload := new(CreateFromYoutubeDTO)
	if err := routes.ParseAndValidatePayload(payload, c); err != nil {
		return routes.HandleError(c, custom_error.NewBadRequestError(err.Error()))
	}

	claim := routes.GetUserClaim(c)
	userId := claim["sub"].(string)

	res, err := h.summaryService.CreateFromYoutubeVideo(userId, payload.Url)
	if err != nil {
		return routes.HandleError(c, custom_error.NewBadRequestError(err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}
