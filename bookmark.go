package linkding

type Bookmark struct {
	Id                 int64    `json:"id"`
	Url                string   `json:"url"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	WebsiteTitle       string   `json:"website_title"`
	WebsiteDescription string   `json:"website_description"`
	IsArchived         bool     `json:"is_archived"`
	Unread             bool     `json:"unread"`
	Shared             bool     `json:"shared"`
	TagNames           []string `json:"tag_names"`
	DateAdded          string   `json:"date_added"`
	DateModified       string   `json:"date_modified"`
}
