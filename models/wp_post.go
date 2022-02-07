package models

type WpPost struct {
	Id        int    `json:id`
	PostName  string `json:post_name`
	PostTitle string `json:post_title`
}
