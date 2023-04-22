package sharedsecret

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler interface {
	Insert(gc *gin.Context)
	Get(gc *gin.Context)
}

type handler struct {
	sharedSecretRepository Repository
}

func NewHandler(sharedSecretRepository Repository) Handler {
	return &handler{
		sharedSecretRepository: sharedSecretRepository,
	}
}

func (h *handler) Insert(gc *gin.Context) {
	var req []SharedSecret

	if err := gc.ShouldBind(&req); err != nil {
		gc.Status(http.StatusBadRequest)
		return
	}

	if err := h.sharedSecretRepository.Insert(req); err != nil {
		gc.Status(http.StatusInternalServerError)
		return
	}

	gc.Status(http.StatusNoContent)
}

func (h *handler) Get(gc *gin.Context) {
	clientID := gc.Param("client_id")

	res, err := h.sharedSecretRepository.GetAndDelete(clientID)
	if err != nil {
		gc.Status(http.StatusInternalServerError)
		return
	}

	gc.JSON(http.StatusOK, res)
}
