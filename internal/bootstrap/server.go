package bootstrap

import (
	"Stat4Market/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	cnt        controller.Controller
}

// NewServer объединили  пакет http , контроллер и мидлвар
func NewServer(cnt controller.Controller) Server {
	return Server{
		httpServer: &http.Server{
			Addr:           ":8080",
			MaxHeaderBytes: 1 << 20,          //1MB
			ReadTimeout:    10 * time.Second, //10 сек
			WriteTimeout:   10 * time.Second,
		},
		cnt: cnt,
	}
}

// InitRoutes инициализируем все наши эндпоинты
func (s Server) InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/api/event", s.cnt.CreateEvent)

	return router

}
