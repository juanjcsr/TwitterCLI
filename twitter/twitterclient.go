package twitter

import (
	"net/http"

	"github.com/dghubble/oauth1"
)

// Client that makes auth requests
type Client struct {
	HTTPClient *http.Client
}

//var requestSecret string
//var requestToken string
//var verifier string
//var oauthErrors error
var accessToken, accessSecret string
var twitterTokens *oauth1.Token
var httpClient *http.Client
var prefs *Preferences

func main() {

}

// NewTwitterClient creates an auth twitter client
func NewTwitterClient(token *oauth1.Token, config oauth1.Config) *Client {
	twitterClient := new(Client)
	twitterClient.HTTPClient = config.Client(oauth1.NoContext, token)
	return twitterClient
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
