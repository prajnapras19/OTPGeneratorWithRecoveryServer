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
	// TODO
	gc.Status(http.StatusNoContent)
}

func (h *handler) Get(gc *gin.Context) {
	// TODO
	gc.Status(http.StatusNoContent)
}
