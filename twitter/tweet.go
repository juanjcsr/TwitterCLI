package twitter

import (
	"html/template"
	"io"
	"log"
)

// Tweet ...
type Tweet struct {
	Entities            Entities
	Coordinates         *Coordinates           `json:"coordinates"`
	CreatedAt           string                 `json:"created_at"`
	FavoriteCount       int                    `json:"favorite_count"`
	Favorited           bool                   `json:"favorited"`
	FilterLevel         string                 `json:"filter_level"`
	ID                  int64                  `json:"id"`
	IDSting             string                 `json:"id_str"`
	InReplyToName       string                 `json:"in_reply_to_screen_name"`
	InReplyToID         int64                  `json:"in_reply_to_status_id"`
	InReplyToUser       string                 `json:"in_reply_to_user_id_str"`
	QuotedStatusID      int64                  `json:"quoted_status_id"`
	QuotedStatusIDStr   string                 `json:"quoted_status_id_str"`
	QuotedStatus        *Tweet                 `json:"quoted_status"`
	PossiblySensitive   bool                   `json:"possibly_sensitive"`
	RetweetCount        int                    `json:"retweet_count"`
	Retweeted           bool                   `json:"retweeted"`
	RetweetedStatus     *Tweet                 `json:"retweeted_status"`
	Source              string                 `json:"source"`
	Scopes              map[string]interface{} `json:"scopes"`
	Text                string                 `json:"text"`
	Truncated           bool                   `json:"truncated"`
	User                User                   `json:"user"`
	WithheldCopyright   bool                   `json:"withheld_copyright"`
	WithheldInCountries []string               `json:"withheld_in_countries"`
	WithheldScope       string                 `json:"withheld_scope"`
}

// TweetSet ...
type TweetSet []Tweet

// Coordinates ...
type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"`
	Type        string     `json:"type"`
}

const twitTemplate = `{{define "singletwit"}}
	By:  @{{.User.ScreenName}}
	---------
	{{.Text}}
	----------
	In reply to: {{.InReplyToName}}
	Created: {{.CreatedAt}}
	_____________________________________ 
	
	{{end}}`

const twitSetTemplate = `{{define "twitset"}}
		TUITS
		{{range .}}
			{{template "singletwit"}}
		{{end}}
	{{end}}`

// Print the tweet
func (t *Tweet) Print(writter io.Writer) {
	report, err := template.New("singletwit").Parse(twitTemplate)
	if err != nil {
		log.Fatal(err)
	}
	report.Execute(writter, t)
}

// PrintSet the tweet
func PrintSet(t []Tweet, writter io.Writer) {
	report := template.Must(template.New("twitset").Parse(twitSetTemplate))

	report.Execute(writter, &t)
}
