package server

import (
	"BIGGO/internal/pkg/storage"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	host    string
	storage *storage.Storage
}

func New(host string, st storage.Storage) *Server {
	s := &Server{
		host:    host,
		storage: &st,
	}

	return s
}

type Entry struct {
	Value string `json:"value"`
}

func (s *Server) newAPI() *gin.Engine {
	engine := gin.New()

	engine.GET("/health", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	engine.GET("/scalar/get/:key", s.handlerGet)

	engine.PUT("/scalar/put/:key", s.handlerSet)

	return engine
}

func (s *Server) handlerSet(ctx *gin.Context) {
	key := ctx.Param("key")

	var v Entry

	if err := json.NewDecoder(ctx.Request.Body).Decode(&v); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}

	s.storage.Set(key, v.Value)
	ctx.AbortWithStatus(http.StatusOK)

}

func (s *Server) handlerGet(ctx *gin.Context) {
	key := ctx.Param("key")

	v := s.storage.Get(key)
	if v == nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, Entry{Value: *v})

}

func (s *Server) Start() {
	s.newAPI().Run(s.host)
}
