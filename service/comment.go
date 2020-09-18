package service

import (
	"blogBackend/model"
	"blogBackend/util"
)

func PublishComment(articleId string, commentPost *model.CommentPost) (*model.CommentPublishReturnModule, error) {

	var (
		comment model.Comment
		commentReturn model.CommentPublishReturnModule
	)

	comment.CommentId = util.GenerateId()
	comment.CommentSenderId = commentPost.CommentSenderId
	comment.CommentContent = commentPost.CommentContent
	comment.CommentReceiverId = articleId
	comment.CommentThumbupNum = 0


	err := model.Save("comment", &comment)
	if err != nil{
		return nil, err
	}

	commentReturn.CommentId = comment.CommentId

	return &commentReturn, nil
}