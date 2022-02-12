package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tarathep/tutorial-backend/db"
	"github.com/tarathep/tutorial-backend/model"
)

type TutorialHandler struct {
	DB db.TutorialRepository
}
type Resp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (h *TutorialHandler) CreateTutorial(c *gin.Context) {

	tutorial := model.Tutorial{}
	if err := c.ShouldBindJSON(&tutorial); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	if err := h.DB.Create(tutorial); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Resp{
		Code:    "200",
		Message: "Inserted a single document Success",
	})

}

func (h *TutorialHandler) ReadTutorials(c *gin.Context) {
	title := c.Query("title")
	tutorials, err := h.DB.FindAll(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, tutorials)
}

func (h *TutorialHandler) ReadTutorial(c *gin.Context) {
	id := c.Param("id")

	tutorials, err := h.DB.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, tutorials)
}

func (h *TutorialHandler) UpdateTutorial(c *gin.Context) {

	tutorial := model.Tutorial{}
	if err := c.ShouldBindJSON(&tutorial); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	if err := h.DB.Update(tutorial); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Resp{
		Code:    "200",
		Message: "Updated a single document Success",
	})
}

func (h *TutorialHandler) DeleteTutorial(c *gin.Context) {
	id := c.Param("id")

	if err := h.DB.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Resp{
		Code:    "200",
		Message: "Deleted id " + id,
	})
}

func (h *TutorialHandler) DeleteTutorials(c *gin.Context) {

	if err := h.DB.DeleteAll(); err != nil {
		c.JSON(http.StatusInternalServerError, Resp{
			Code:    "500",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Resp{
		Code:    "200",
		Message: "All deleted",
	})
}
