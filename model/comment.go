package model

type Comment struct {
	CommentId         string `json:"comment_id"`
	CommentReceiverId string `json:"comment_receiver_id"`
	CommentSenderId   string `json:"comment_sender_id"`
	CommentContent    string `json:"comment_content"`
	CommentThumbupNum int `json:"comment_thumbup_num"`
}

type CommentPost struct {
	CommentSenderId string `json:"comment_sender_id"`
	CommentContent  string `json:"comment_content"`
}

type CommentPublishReturnModule struct {
	CommentId string `json:"comment_id"`
}