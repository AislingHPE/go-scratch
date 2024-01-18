package gin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hpe-hcss/monitoring/pkg/monitoring"
	"github.com/hpe-hcss/tracing/pkg/trace"
)

func Main() {
	gr := gin.New()
	gr.HandleMethodNotAllowed = true

	gr.Use(gin.Recovery())
	gr.Use(monitoring.Middleware)
	gr.Use(trace.Middleware())
	group := gr.Group("group")
	a := group.Group("a")
	a.Use(Sandal)
	gr.Any("group/a/b", Handle)
	a.Any("b", Handle)
	gr.Any("lalala/b", Handle)
	fmt.Print(a)
}

func Handle(c *gin.Context) {}
func Sandal(c *gin.Context) {}
