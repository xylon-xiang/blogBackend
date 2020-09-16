package model

type Comment struct {
	CommentId         string `json:"comment_id"`
	CommentReceiverId string `json:"comment_receiver_id"`
	CommentSenderId   string `json:"comment_sender_id"`
	CommentContent    string `json:"comment_content"`
	CommentThumbupNum string `json:"comment_thumbup_num"`
}
