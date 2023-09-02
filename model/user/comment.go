package user

type Comment struct {
	Id            int       `json:"id"`
	GoodCommentId int       `json:"good_comment_id"`
	CommentId     int       `json:"comment_id"`
	UserId        string    `json:"user_id"`
	GoodId        string    `json:"good_id"`
	Content       string    `json:"content"`
	Replay        []*Replay `json:"replay" gorm:"many2many:comment_replay;foreignKey:Id;joinForeignKey:commentId;joinReferences:replayId"`
}

func (c *Comment) TableName() string {
	return "comment"
}
