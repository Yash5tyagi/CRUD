package handler

import "github.com/gin-gonic/gin"

type Serv struct {
	Srv *gin.Engine
}

func New() (s *Serv) {
	return &Serv{
		Srv: gin.Default(),
	}
}
