package service

import (
	"blogBackend/model"
	"blogBackend/util"
	"github.com/deckarep/golang-set"
	"regexp"
	"time"
)

func GetArticle(isAll bool, tags string, articleId string) (*model.ArticleInfoReturnModule, error) {

	articleInfo := mapset.NewSet()

	articleReturn := new(model.ArticleInfoReturnModule)

	if isAll {
		if tags == "" {
			articles, err := model.FindAll("article", "", "", "")
			if err != nil {
				return nil, err
			}

			articleReturn.Articles = articles.([]model.Article)

		} else {
			//([\u4400-\u9fa5, \w]*)[,，]{1}
			tagSlice := regexp.MustCompile("(\\w*)[,，]{1}").FindAll([]byte(tags), -1)
			for _, tag := range tagSlice {
				articles, err := model.FindAll("article", "", "", string(tag))
				if err != nil {
					return nil, err
				}

				for _, article := range articles.([]model.Article) {
					articleInfo.Add(article)
				}
			}

			articleInfo.Each(func(i interface{}) bool {
				articleReturn.Articles = append(articleReturn.Articles, i.(model.Article))
				return i == nil
			})
		}

		// get the specific articles by id and relative comments
	} else {
		article, err := model.FindById("article", articleId)
		if err != nil {
			return nil, err
		}

		articleReturn.Articles = append(articleReturn.Articles, article.(model.Article))

		comments, err := model.FindAll("comment", "comment_receiver_id", articleId, "")
		if err != nil {
			return nil, err
		}

		articleReturn.Comments = comments.([]model.Comment)

	}

	return articleReturn, nil
}

func PublishArticle(postArticle *model.ArticleUpload) (*model.ArticlePublishReturnModule, error) {

	article := new(model.Article)

	article.ArticleId = util.GenerateId()
	article.ArticleAuthorId = postArticle.ArticleAuthorId
	article.ArticleContent = postArticle.ArticleContent
	article.ArticleName = postArticle.ArticleName
	article.ArticleTag = postArticle.ArticleTag
	article.ArticleSummary = util.GenerateSummary(postArticle.ArticleContent)
	article.ArticlePublishTime = time.Now().Unix()
	article.ArticleUpdateTime = time.Now().Unix()
	article.ArticleThumbupNum = 0

	err := model.Save("article", article)
	if err != nil {
		return nil, err
	}

	articlePubReturn := model.ArticlePublishReturnModule{
		ArticleId:      article.ArticleId,
		ArticleSummary: article.ArticleSummary,
	}

	return &articlePubReturn, nil
}

func UpdateArticle(postArticle *model.ArticleUpload) error {

	result, err := model.FindById("article", postArticle.ArticleId)
	if err != nil {
		return err
	}
	article := result.(model.Article)

	article.ArticleName = postArticle.ArticleName
	article.ArticleAuthorId = postArticle.ArticleAuthorId
	article.ArticleContent = postArticle.ArticleContent
	article.ArticleTag = postArticle.ArticleTag
	article.ArticleUpdateTime = time.Now().Unix()

	err = model.UpdateOne("article", article)
	if err != nil {
		return err
	}

	return nil
}

func DeleteArticleById(articleId string) error {

	err := model.DeleteById("article", articleId)
	if err != nil {
		return err
	}

	return nil
}
