package models

type Comment struct {
	Id         int
	Content    string
	GameId     string
	CreateTime string
	WriterId   string
}

func GetCommentList(game_id int) *[]Comment {
	var comments []Comment
	qqr := o.QueryTable("comment").Filter("game_id", game_id).OrderBy("-create_time").Limit(10)
	qqr.All(&comments)
	return &comments

}
