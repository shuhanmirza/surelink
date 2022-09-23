package external

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2"
	auth "golang.org/x/oauth2/google"
)

// this is a test code. will implement later
//https://pkg.go.dev/google.golang.org/api/idtoken
//https://stackoverflow.com/questions/72275338/get-access-token-for-a-google-cloud-service-account-in-golang
func gcp_webrisk() {

	jsonFile, err := os.Open("cred.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonData, _ := ioutil.ReadAll(jsonFile)

	var token *oauth2.Token
	ctx := context.Background()
	scopes := []string{
		"https://www.googleapis.com/auth/cloud-platform",
	}
	credentials, err := auth.CredentialsFromJSON(ctx, jsonData, scopes...)
	if err == nil {
		log.Printf("found default credentials. %v", credentials)
		token, err = credentials.TokenSource.Token()
		log.Printf("token: %v, err: %v", token, err)
		if err != nil {
			log.Print(err)
		}

		baseUrl := "https://webrisk.googleapis.com/v1/uris:search?"
		params := url.Values{}
		params.Add("threatTypes", "MALWARE")
		params.Add("threatTypes", "SOCIAL_ENGINEERING")
		params.Add("threatTypes", "UNWANTED_SOFTWARE")
		params.Add("uri", "http://testsafebrowsing.appspot.com/s/malware.html")
		finalUrl := baseUrl + params.Encode()

		req, err := http.NewRequest("GET", finalUrl, nil)
		if err != nil {
			fmt.Println(err)
		}

		token.SetAuthHeader(req)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(resp.Body)

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Println(bodyString)
		} else {
			fmt.Println(resp.StatusCode)
		}
	}

}
