package twitter

import (
	"net/http"

	"github.com/dghubble/oauth1"
)

// Client that makes auth requests
type Client struct {
	httpClient *http.Client
	Timeline   *TimelineService
}

const twitterBaseURL = "https://api.twitter.com/1.1/"

//var requestSecret string
//var requestToken string
//var verifier string
//var oauthErrors error
// var accessToken, accessSecret string
// var twitterTokens *oauth1.Token
// var httpClient *http.Client
// var prefs *Preferences

// NewTwitterClient creates an auth twitter client
func NewTwitterClient(token *oauth1.Token, config oauth1.Config) *Client {
	///twitterClient := new(Client)
	// twitterClient.httpClient = config.Client(oauth1.NoContext, token)
	// return twitterClient
	authedClient := config.Client(oauth1.NoContext, token)
	return &Client{
		httpClient: authedClient,
		Timeline:   newTimeLineService(authedClient),
	}
}

// IsLoggedIn checks if the access tokens are working
func (c *Client) IsLoggedIn() bool {
	resp, _ := c.httpClient.Get(twitterBaseURL + "account/verify_credentials.json")
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true

}

// func handleAuth(w http.ResponseWriter, r *http.Request) {

// 	requestToken, requestSecret, oauthErrors = oauthConfig.RequestToken()
// 	authURL, err := oauthConfig.AuthorizationURL(requestToken)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	http.Redirect(w, r, authURL.String(), http.StatusTemporaryRedirect)
// }

// func handleOauth(w http.ResponseWriter, r *http.Request) {
// 	requestToken, verifier, oauthErrors := oauth1.ParseAuthorizationCallback(r)
// 	if oauthErrors != nil {
// 		fmt.Println("Error")
// 	}
// 	accessToken, accessSecret, oauthErrors := oauthConfig.AccessToken(requestToken, requestSecret, verifier)
// 	if oauthErrors != nil {
// 		fmt.Println(oauthErrors)
// 	}
// 	twitterTokens = oauth1.NewToken(accessToken, accessSecret)
// 	fmt.Printf("token: %s\nsecret: %s\n", twitterTokens.Token, twitterTokens.TokenSecret)
// 	prefs.Update("token", twitterTokens.Token)
// 	prefs.Update("token_secret", twitterTokens.TokenSecret)
// 	//httpClient := oauthConfig.Client(oauth1.NoContext, twitterTokens)

// }
