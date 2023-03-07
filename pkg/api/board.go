package api

import (
	"net/http"
	"strconv"

	"github.com/choisangh/board-crud-backend/pkg/model"
	"github.com/gin-gonic/gin"
)

func (apis *APIs) CreateBoard(c *gin.Context) {
	req := &model.Board{}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	res, err := apis.db.CreateBoard(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) GetBoardList(c *gin.Context) {
	res, err := apis.db.GetBoardList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{Res: "Server error"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) GetBoardByID(c *gin.Context) {
	id := c.Param("id")

	res, err := apis.db.GetBoardByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) UpdateBoard(c *gin.Context) {
	idS := c.Param("id")
	req := &model.Board{}

	id, err := strconv.ParseUint(idS, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	res, err := apis.db.UpdateBoard(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) DeleteBoardByID(c *gin.Context) {
	id := c.Param("id")

	err := apis.db.DeleteBoardByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})

		return
	}

	c.JSON(http.StatusOK, &Response{Res: "Success"})
}
