package api

import (
	"github.com/Ventilateur/crew-interview/domain/talent"
	"github.com/Ventilateur/crew-interview/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPHandler interface {
	ListTalents(ctx *gin.Context)
	AddTalent(ctx *gin.Context)
	Health(ctx *gin.Context)
}

type httpHandler struct {
	talentApp talent.Application
}

func NewHTTPHAndler(talentApp talent.Application) HTTPHandler {
	return &httpHandler{
		talentApp: talentApp,
	}
}

func (h *httpHandler) Health(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		statusResponse{
			Status: statusOK,
		},
	)
}

func (h *httpHandler) ListTalents(ctx *gin.Context) {
	talents, err := h.talentApp.ListTalents(ctx, "", 0)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	resp := make([]getTalentResponse, len(talents))
	for i, t := range talents {
		resp[i] = getTalentResponse{
			Id:        utils.StringPtr(t.Id),
			FirstName: utils.StringPtr(t.FirstName),
			LastName:  utils.StringPtr(t.LastName),
			Picture:   utils.StringPtr(t.Picture),
			Job:       utils.StringPtr(t.Job),
			Location:  utils.StringPtr(t.Location),
			LinkedIn:  utils.StringPtr(t.LinkedIn),
			Github:    utils.StringPtr(t.Github),
			Twitter:   utils.StringPtr(t.Twitter),
			Tags:      make([]string, len(t.Tags)),
			Stage:     utils.StringPtr(t.Stage),
		}
		copy(resp[i].Tags, t.Tags)
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *httpHandler) AddTalent(ctx *gin.Context) {
	var req addTalentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	entity := req.Entity()
	if err := h.talentApp.AddTalent(ctx, entity); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, addTalentResponse{
		Id: entity.Id,
	})
}
