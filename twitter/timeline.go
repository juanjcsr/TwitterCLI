package twitter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// TimelineService ...
type TimelineService struct {
	client *http.Client
}

const homeTimelineURL = "https://api.twitter.com/1.1/statuses/home_timeline.json"
const userTimelineURL = "https://api.twitter.com/1.1/statuses/user_timeline.json"

// NewTimeLineService ...
func newTimeLineService(authedClient *http.Client) *TimelineService {

	var theTimelineService = new(TimelineService)
	theTimelineService.client = authedClient
	return theTimelineService
}

// GetHomeTimeline ...
func (t *TimelineService) GetHomeTimeline() ([]Tweet, error) {
	tlURL, _ := url.Parse(homeTimelineURL)
	resp, err := t.client.Get(tlURL.String())
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("timeline: could not return Tweets: %s", err.Error())
	}
	tweets := new([]Tweet)
	if err := json.NewDecoder(resp.Body).Decode(tweets); err != nil {
		return nil, fmt.Errorf("timeline: could not decode Tweets: %s", err.Error())
	}

	return *tweets, nil
}

// GetUserTimeline ..
func (t *TimelineService) GetUserTimeline(user string) ([]Tweet, error) {
	userURL, _ := url.Parse(userTimelineURL)
	if user != "" {
		params := url.Values{}
		params.Set("screen_name", user)
		userURL.RawQuery = params.Encode()
	}
	resp, _ := t.client.Get(userURL.String())
	defer resp.Body.Close()
	var tweets = new([]Tweet)
	if err := json.NewDecoder(resp.Body).Decode(tweets); err != nil {
		return nil, fmt.Errorf("timeline: could not decode User Timeline Tweets: %s", err.Error())
	}
	return *tweets, nil
}
