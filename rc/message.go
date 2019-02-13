package rc

import "time"

type Message struct {
	Alias       string       `json:"alias,omitempty"`
	Avatar      string       `json:"avatar,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Emoji       string       `json:"emoji,omitempty"`
	RoomID      string       `json:"room_id,omitempty"`
	Text        string       `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type Attachment struct {
	AudioURL          string    `json:"audio_url,omitempty"`
	AuthorName        string    `json:"author_name,omitempty"`
	AuthorLink        string    `json:"author_link,omitempty"`
	AuthorIcon        string    `json:"author_icon,omitempty"`
	Collapsed         bool      `json:"collapsed,omitempty"`
	Color             string    `json:"color,omitempty"`
	Fields            []Field   `json:"fields,omitempty"`
	ImageURL          string    `json:"image_url,omitempty"`
	MessageLink       string    `json:"message_link,omitempty"`
	Text              string    `json:"text,omitempty"`
	ThumbURL          string    `json:"thumb_url,omitempty"`
	Title             string    `json:"title,omitempty"`
	TitleLink         string    `json:"title_link,omitempty"`
	TitleLinkDownload bool      `json:"title_link_download,omitempty"`
	Timestamp         time.Time `json:"ts,omitempty"`
	VideoURL          string    `json:"video_url,omitempty"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}
