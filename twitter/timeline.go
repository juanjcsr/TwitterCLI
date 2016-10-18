package twitter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TimelineService ...
type TimelineService struct {
	client *http.Client
}

const timelineURL = "https://api.twitter.com/1.1/statuses/home_timeline.json"

// NewTimeLineService ...
func newTimeLineService(authedClient *http.Client) *TimelineService {
	var theTimelineService = new(TimelineService)
	theTimelineService.client = authedClient
	return theTimelineService
}

// GetTimeline ...
func (t *TimelineService) GetTimeline() ([]Tweet, error) {
	resp, err := t.client.Get(timelineURL)
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
