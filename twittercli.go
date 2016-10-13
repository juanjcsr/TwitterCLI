package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/juanjcsr/twittercli/twitter"
)

const tokenStorage = "token"
const tokenSecretStorage = "token_secret"

func main() {
	client := login()
	resp, err := client.HTTPClient.Get("https://api.twitter.com/1.1/statuses/user_timeline.json")
	if err != nil {
		fmt.Printf("Could not get, error: %s", err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("TIMELINE: %v", string(body))

}

func login() *twitter.Client {
	//const callbackURL = "http://localhost:8080/oauthcallback"
	const callbackURL = "oob"

	var oauthConfig = &oauth1.Config{
		// ClientID is the application's ID.
		ConsumerKey: os.Getenv("TWITTER_CONSUMER_KEY"),

		// ClientSecret is the application's secret.
		ConsumerSecret: os.Getenv("TWITTER_CONSUMER_SECRET"),
		// Endpoint contains the resource server's ,token endpoint
		// URLs. These are constants specific to each server and are
		// often available via site-specific packages, such as
		// google.Endpoint or github.Endpoint.
		Endpoint: oauth1.Endpoint{
			AuthorizeURL:    "https://api.twitter.com/oauth/authorize",
			RequestTokenURL: "https://api.twitter.com/oauth/request_token",
			AccessTokenURL:  "https://api.twitter.com/oauth/access_token",
		},

		// RedirectURL is the URL to redirect users going through
		// the OAuth flow, after the resource owner's URLs.
		CallbackURL: callbackURL,
	}
	prefs := new(twitter.Preferences)
	if _, err := prefs.Open(); err != nil {
		fmt.Printf("error %v", err)
	}
	defer prefs.Close()
	token := checkForTokensInDB(prefs)
	var twitterClient *twitter.Client
	if token == nil {
		reqToken, err := loginUserCLI(oauthConfig)
		if err != nil {
			log.Fatalf("Request token Phase: %s", err.Error())
		}
		token, err := getPinCLI(reqToken, oauthConfig)
		if err != nil {
			log.Fatalf("Access token phase: %s", err.Error())
		}
		prefs.Update(tokenStorage, token.Token)
		prefs.Update(tokenSecretStorage, token.TokenSecret)
		twitterClient = twitter.NewTwitterClient(token, *oauthConfig)
	} else {
		twitterClient = twitter.NewTwitterClient(token, *oauthConfig)
	}

	return twitterClient

	//http.HandleFunc("/", handleAuth)
	//http.HandleFunc("/oauthcallback", handleOauth)
	//http.ListenAndServe(":8080", nil)
}

func checkForTokensInDB(prefs *twitter.Preferences) *oauth1.Token {

	token := prefs.Read(tokenStorage)
	secret := prefs.Read(tokenSecretStorage)

	if token == "" && secret == "" {
		return nil
	}
	return oauth1.NewToken(token, secret)
}

func loginUserCLI(oauthConfig *oauth1.Config) (string, error) {
	requestToken, _, err := oauthConfig.RequestToken()
	authURL, err := oauthConfig.AuthorizationURL(requestToken)
	if err != nil {
		return "", err
	}
	fmt.Printf("Open this URL in your browser: \n%s\n", authURL.String())
	return requestToken, err
}

func getPinCLI(reqToken string, oauthConfig *oauth1.Config) (*oauth1.Token, error) {
	fmt.Printf("Paste the PIN here: ")
	var verifier string
	_, err := fmt.Scanf("%s", &verifier)
	if err != nil {
		return nil, err
	}
	accessToken, accessSecret, err := oauthConfig.AccessToken(reqToken, "secret does not matter", verifier)
	if err != nil {
		return nil, err
	}
	return oauth1.NewToken(accessToken, accessSecret), nil
}
