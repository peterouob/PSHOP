package user

type Replay struct {
	Id        int        `json:"id"`
	ReplayId  int        `json:"replay_id"`
	UserId    string     `json:"user_id"`
	CommentId int        `json:"comment_id"`
	Content   string     `json:"content"`
	Comment   []*Comment `json:"comment" gorm:"many2many:comment_replay;foreignKey:Id;joinForeignKey:replayId;joinReferences:commentId"`
}

func (r *Replay) TableName() string {
	return "replay"
}
