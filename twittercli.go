package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/oauth1"
	"github.com/juanjcsr/twittercli/twitter"
)

const accessToken = "access_token"
const accessSecret = "access_secret"
const tokenStorage = "token"
const tokenSecretStorage = "token_secret"

func main() {
	client, _ := Login()

	resp, err := client.Timeline.GetTimeline()
	if err != nil {
		fmt.Printf("Could not get, error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("TIMELINE: %v", resp)

}

func saveTwitterKeys(prefs *twitter.Preferences) (string, string) {
	fmt.Println("Getting ACCESS Keys")
	var consumer, secret string
	fmt.Printf("Paste your TWITTER CONSUMER KEY here: ")
	if _, err := fmt.Scanf("%s", &consumer); err != nil {
		fmt.Println("You need to provide your consumer key")
		os.Exit(1)
	}

	fmt.Printf("Paste your TWITTER SECRET KEY here: ")
	if _, err := fmt.Scanf("%s", &secret); err != nil {
		fmt.Println("You need to provide your secret key")
		os.Exit(1)
	}
	return consumer, secret
}

func getTwitterKeys(prefs *twitter.Preferences) (string, string, error) {
	token := prefs.Read(accessToken)
	secret := prefs.Read(accessSecret)

	if token == "" || secret == "" {
		return token, secret, fmt.Errorf("twitter keys: no keys")
	}

	return token, secret, nil
}

// Login logs the user and requests all the needed tokens
func Login() (*twitter.Client, error) {
	//const callbackURL = "http://localhost:8080/oauthcallback"
	const callbackURL = "oob"
	var twitterClient *twitter.Client

	prefs := new(twitter.Preferences)
	if _, err := prefs.Open(); err != nil {
		fmt.Printf("error %v", err)
	}
	defer prefs.Close()

	consumer, secret, noTokens := getTwitterKeys(prefs)
	if noTokens != nil {
		consumer, secret = saveTwitterKeys(prefs)
	}

	var oauthConfig = &oauth1.Config{
		ConsumerKey:    consumer,
		ConsumerSecret: secret,
		Endpoint: oauth1.Endpoint{
			AuthorizeURL:    "https://api.twitter.com/oauth/authorize",
			RequestTokenURL: "https://api.twitter.com/oauth/request_token",
			AccessTokenURL:  "https://api.twitter.com/oauth/access_token",
		},
		CallbackURL: callbackURL,
	}

	var token *oauth1.Token

	if noTokens != nil {
		reqToken, err := loginUserCLI(oauthConfig)
		if err != nil {
			log.Fatalf("Request token Phase: %s", err.Error())
		}
		token, err = getPinCLI(reqToken, oauthConfig)
		if err != nil {
			log.Fatalf("Access token phase: %s", err.Error())
		}

		twitterClient = twitter.NewTwitterClient(token, *oauthConfig)
	} else {
		token = getTokensInDB(prefs)
		twitterClient = twitter.NewTwitterClient(token, *oauthConfig)
	}

	if twitterClient.IsLoggedIn() {
		prefs.Update(accessToken, consumer)
		prefs.Update(accessSecret, secret)
		prefs.Update(tokenStorage, token.Token)
		prefs.Update(tokenSecretStorage, token.TokenSecret)

		return twitterClient, nil
	}
	return nil, fmt.Errorf("login: could not auth user")

	//http.HandleFunc("/", handleAuth)
	//http.HandleFunc("/oauthcallback", handleOauth)
	//http.ListenAndServe(":8080", nil)
}

func getTokensInDB(prefs *twitter.Preferences) *oauth1.Token {

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
