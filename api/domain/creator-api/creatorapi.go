package creatorapi

import (
	"net/http"

	creatorapp "github.com/bentenison/microservice/app/domain/creator-app"
	"github.com/bentenison/microservice/foundation/logger"
	"github.com/gin-gonic/gin"
)

type api struct {
	creatorapp *creatorapp.App
	log        *logger.CustomLogger
	// proto.UnimplementedAuthServiceServer
}

func newAPI(log *logger.CustomLogger, creatorapp *creatorapp.App) *api {
	return &api{
		creatorapp: creatorapp,
		log:        log,
	}
}

func (a *api) createNewQuestions(c *gin.Context) {
	var questions []creatorapp.Question
	if err := c.Bind(&questions); err != nil {
		a.log.Errorc(c.Request.Context(), "error binding data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}

	res, err := a.creatorapp.AddNewQuestions(c.Request.Context(), questions)
	if err != nil {
		a.log.Errorc(c.Request.Context(), "error adding data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func (a *api) getAllQuestions(c *gin.Context) {
	// var questions []creatorapp.Question
	// if err := c.Bind(&questions); err != nil {
	// 	a.log.Errorc(c.Request.Context(), "error binding data", map[string]interface{}{
	// 		"error": err.Error(),
	// 	})
	// 	c.JSON(http.StatusExpectationFailed, err.Error())
	// 	return
	// }

	res, err := a.creatorapp.GetAllQuestionsDAO(c.Request.Context())
	if err != nil {
		a.log.Errorc(c.Request.Context(), "error getting data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func (a *api) getQuestionByTag(c *gin.Context) {
	// var questions []creatorapp.Question
	tag, ok := c.Params.Get("tag")
	if !ok {
		a.log.Errorc(c.Request.Context(), "tag is mandatory", map[string]interface{}{
			"error": "tag is mandatory parameter",
		})
		c.JSON(http.StatusExpectationFailed, "tag is mandatory parameter")
		return
	}

	res, err := a.creatorapp.GetQuestionsByTag(c.Request.Context(), tag)
	if err != nil {
		a.log.Errorc(c.Request.Context(), "error adding data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func (a *api) getQuestionsBylang(c *gin.Context) {
	// var questions []creatorapp.Question
	lang, ok := c.Params.Get("lang")
	if !ok {
		a.log.Errorc(c.Request.Context(), "error: language is mandatory.", map[string]interface{}{
			"error": "error: language is mandatory.",
		})
		c.JSON(http.StatusExpectationFailed, "error: language is mandatory.")
		return
	}

	res, err := a.creatorapp.GetQuestionsByLang(c.Request.Context(), lang)
	if err != nil {
		a.log.Errorc(c.Request.Context(), "error adding data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func (a *api) qcQuestion(c *gin.Context) {
	// var questions []creatorapp.Question
	var questions creatorapp.Question
	if err := c.Bind(&questions); err != nil {
		a.log.Errorc(c.Request.Context(), "error binding data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}

	// res, err := a.creatorapp.GetQuestionsByLang(c.Request.Context(), lang)
	// if err != nil {
	// 	a.log.Errorc(c.Request.Context(), "error adding data", map[string]interface{}{
	// 		"error": err.Error(),
	// 	})
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, res)
}
func (a *api) getSingleQuestion(c *gin.Context) {
	// var questions []creatorapp.Question
	id, ok := c.Params.Get("id")
	if !ok {
		a.log.Errorc(c.Request.Context(), "error: id is mandatory.", map[string]interface{}{
			"error": "error: id is mandatory.",
		})
		c.JSON(http.StatusExpectationFailed, "error: id is mandatory.")
		return
	}

	res, err := a.creatorapp.GetSingleQuestion(c.Request.Context(), id)
	if err != nil {
		a.log.Errorc(c.Request.Context(), "error getting data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
func (a *api) deleteSelectedQuestion(c *gin.Context) {
	// var questions []creatorapp.Question
	var ids []string
	if err := c.Bind(&ids); err != nil {
		a.log.Errorc(c.Request.Context(), "error binding data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusExpectationFailed, err.Error())
		return
	}

	res, err := a.creatorapp.DeleteSelectedQuestions(c.Request.Context(), ids)
	if err != nil {
		a.log.Errorc(c.Request.Context(), "error adding data", map[string]interface{}{
			"error": err.Error(),
		})
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
