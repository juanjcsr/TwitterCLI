package twitter

// Entities ...
type Entities struct {
	Hashtags     []Hashtags
	Media        []MediaEntity
	Urls         []Urls
	UserMentions []UserMention `json:"user_mentions"`
}

// Hashtags ...
type Hashtags struct {
	Indices []int
	Text    string
}

// Urls ...
type Urls struct {
	Indices     []int
	URL         string
	DisplayURL  string `json:"display_url"`
	ExpandedURL string `json:"expanded_url"`
}

// UserMention ..
type UserMention struct {
	ID         int64
	IDStr      string `json:"id_str"`
	Indices    []int
	Name       string
	ScreenName string `json:"screen_name"`
}

// MediaEntity ...
type MediaEntity struct {
	ID            int64
	IDStr         string `json:"id_str"`
	MediaURL      string `json:"media_url"`
	MediaURLHTTPS string `json:"media_url_https"`
	URL           string
	DisplayURL    string `json:"display_url"`
	ExpandedURL   string `json:"expanded_url"`
	Sizes         MediaSizes
	Type          string
	Indices       []int
	// TODO add VideoInfo
}

// MediaSizes ...
type MediaSizes struct {
	Medium MediaSize
	Thumb  MediaSize
	Small  MediaSize
	Large  MediaSize
}

// MediaSize ..
type MediaSize struct {
	W      int
	H      int
	Resize string
}
