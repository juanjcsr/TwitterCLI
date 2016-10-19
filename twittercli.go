package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/dghubble/oauth1"
	"github.com/juanjcsr/twittercli/twitter"
)

const accessToken = "access_token"
const accessSecret = "access_secret"
const tokenStorage = "token"
const tokenSecretStorage = "token_secret"

func main() {

	l, _ := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	defer l.Close()

	log.SetOutput(l.Stderr())

	client, _ := Login()

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		log.Printf("LINE: %s", line)
		switch {
		case strings.HasPrefix(line, "user"):
			//line := strings.TrimSpace(line[4:])
			username := strings.TrimPrefix(line, "user")
			log.Printf("username: %s", username)
			if username != "" {
				log.Printf("LINE in user: %s", line)
				resp, err := client.Timeline.GetUserTimeline(username)
				if err != nil {
					log.Printf("Could not get, error: %s", err.Error())
					break
				}

				for _, tuit := range resp {
					tuit.Print(os.Stdout)
				}
			}

			// arg, err := l.Readline()
			// if arg == "" {
			// 	log.Printf("error: %s", err)
			// 	break
			// }
			// log.Printf("HOLAAA, %s", arg)
		case strings.HasPrefix(line, "home"):
			resp, err := client.Timeline.GetHomeTimeline()
			if err != nil {
				log.Printf("Could not get, error: %s", err.Error())
				break
			}
			for _, tuit := range resp {
				tuit.Print(os.Stdout)
			}
		}
	}

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
