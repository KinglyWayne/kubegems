package types

import "time"

/*
post:										<--- post id

[post content]

comments:

- [user a] at 00:12 : nice job!				<--- comment(id)
- [user b] at 00:15 : where is the cat?		<--- comment's content
- [user c] at 00:35 : > where is the cat?	<--- reply to a comment(id)
					  in the dark!
*/
// nolint: tagliatelle
type Comment struct {
	ID           string    `json:"id,omitempty" bson:"_id,omitempty"` // id
	Username     string    `json:"username,omitempty"`                // user's id (username)
	PostID       string    `json:"postID,omitempty"`                  // post id
	Content      string    `json:"content,omitempty"`                 // comment's content
	Replies      []Comment `json:"replies,omitempty"`                 // comment's replies
	ReplyID      string    `json:"replyID,omitempty"`                 // comment's id reply to
	Rating       int       `json:"rating,omitempty"`                  // rating value (1-5)
	CreationTime time.Time `json:"creationTime,omitempty"`            // comment's create time
	UpdationTime time.Time `json:"updationTime,omitempty"`            // comment's update time
	Deleted      bool      `json:"deleted,omitempty"`                 // comment's deleted
}
