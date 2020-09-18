package main

import (
	"blogBackend/model"
	"blogBackend/service"
	"blogBackend/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func main() {

	e := echo.New()

	e.POST("/api/user/jwt", userLoginController)

	e.GET("/api/user/:userId", getSpecificUserController,
		middleware.JWTWithConfig(util.JwtConfig))

	e.GET("/api/articles", getAllArticleController,
		middleware.JWTWithConfig(util.JwtConfig))

	e.GET("/api/article/:articleId", getSpecificArticleController,
		middleware.JWTWithConfig(util.JwtConfig))

	e.POST("/api/article", publishArticleController,
		middleware.JWTWithConfig(util.JwtConfig))

	e.PUT("/api/article/:articleId", updateArticleController,
		middleware.JWTWithConfig(util.JwtConfig))

	e.DELETE("/api/article/:articleId", deleteArticleController,
		middleware.JWTWithConfig(util.JwtConfig))

	e.POST("/api/article/:articleId", publishCommentController)

	e.Logger.Fatal(e.Start(":1548"))

}

func userLoginController(context echo.Context) error {

	userPost := new(model.UserLogPost)
	if err := context.Bind(userPost); err != nil {
		return err
	}

	passwordRight, err := service.LogJudge(userPost.UserId, userPost.UserPassword)
	if err != nil {
		return err
	}

	if !passwordRight {
		return echo.ErrUnauthorized
	}

	logReturnModule, err := util.GenerateJwtToken(userPost.UserId)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, *logReturnModule)
}

func getSpecificUserController(context echo.Context) error {

	userId := context.Param("userId")

	userInfo, err := service.GetSpecificUser(userId)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, *userInfo)

}

func getAllArticleController(context echo.Context) error {

	tags := context.QueryParam("tag")

	results, err := service.GetArticle(true, tags, "")
	if err != nil {
		//switch err {
		//case mongo.ErrNoDocuments:
		//	return context.String(http.StatusNotFound, "no matched documents")
		//}
		return err
	}

	return context.JSON(http.StatusOK, results)
}

func getSpecificArticleController(context echo.Context) error {

	articleId := context.Param("articleId")

	results, err := service.GetArticle(false, "", articleId)
	if err != nil {
		//switch err {
		//case mongo.ErrNilDocument:
		//	return context.String(http.StatusNotFound, "no matched documents")
		//}
		return err
	}

	return context.JSON(http.StatusOK, results)
}

func publishArticleController(context echo.Context) error {

	articlePost := new(model.ArticleUpload)
	if err := context.Bind(articlePost); err != nil {
		return err
	}

	results, err := service.PublishArticle(articlePost)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusCreated, results)
}

func updateArticleController(context echo.Context) error {
	articlePut := new(model.ArticleUpload)
	if err := context.Bind(articlePut); err != nil {
		return err
	}

	err := service.UpdateArticle(articlePut)
	if err != nil {
		return err
	}

	return context.String(http.StatusOK, "resource updated successfully")
}

func deleteArticleController(context echo.Context) error {
	articleId := context.Param("articleId")

	err := service.DeleteArticleById(articleId)
	if err != nil {
		return err
	}

	return context.String(http.StatusOK, "resource deleted successfully")
}

func publishCommentController(context echo.Context) error {

	articleId := context.Param("articleId")

	commentPost := new(model.CommentPost)

	if err := context.Bind(commentPost); err != nil{
		return err
	}

	commentPubReturn, err := service.PublishComment(articleId, commentPost)
	if err != nil{
		return err
	}

	return context.JSON(http.StatusOK, *commentPubReturn)

}
