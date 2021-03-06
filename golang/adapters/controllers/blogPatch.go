package controllers

import (
	"net/http"
	"strconv"

	"clean_architecture/golang/adapters/presenters"
	"clean_architecture/golang/adapters/presenters/json"
	uc "clean_architecture/golang/usecases"

	"github.com/gin-gonic/gin"
)

func (rH RouterHandler) blogPatch(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	req := &BlogRequest{}
	if err := c.BindJSON(req); err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log(err)
		c.Status(http.StatusBadRequest)
		return
	}

	useCase := uc.EditBlogUseCase{
		OutputPort: json.NewPresenter(presenters.New(c)),
		InputPort: uc.EditBlogParams{
			Id:    id,
			Title: *req.Blog.Title,
			Body:  *req.Blog.Body,
		},
	}
	rH.ucHandler.BlogEdit(useCase)

}
