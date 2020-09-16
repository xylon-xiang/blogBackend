package model

type Article struct {
	ArticleId          string `json:"article_id" bson:"article_id"`
	ArticleName        string `json:"article_name" bson:"article_name"`
	ArticleAuthorId    string `json:"article_author_id" bson:"article_author_id"`
	ArticleSummary     string `json:"article_summary" bson:"article_summary"`
	ArticlePublishTime int64  `json:"article_publish_time" bson:"article_publish_time"`
	ArticleContent     string `json:"article_content" bson:"article_content"`
	ArticleTag         string `json:"article_tag" bson:"article_tag"`
	ArticleThumbupNum  int    `json:"article_thumbup_num" bson:"article_thumbup_num"`
	ArticleUpdateTime  int64  `json:"article_update_time" bson:"article_update_time"`
}

type ArticleInfoReturnModule struct {
	Articles []Article `json:"articles"`
	Comments []Comment `json:"comments"`
}

type ArticleUpload struct {
	ArticleId       string `json:"article_id"`
	ArticleName     string `json:"article_name"`
	ArticleAuthorId string `json:"article_author_id"`
	ArticleContent  string `json:"article_content"`
	ArticleTag      string `json:"article_tag"`
}

type ArticlePublishReturnModule struct {
	ArticleId      string `json:"article_id"`
	ArticleSummary string `json:"article_summary"`
}
