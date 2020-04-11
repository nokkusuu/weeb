package weeb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var client *http.Client

type TophResp struct {
	Status   int      `json:"status"`
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	BaseType string   `json:"baseType"`
	NSFW     bool     `json:"nsfw"`
	Filetype string   `json:"fileType"`
	MIMEType string   `json:"mimeType"`
	Account  string   `json:"account"`
	Hidden   bool     `json:"hidden"`
	Tags     []string `json:"tags"`
	URL      string   `json:"url"`
}

func init() {
	client = &http.Client{}
}

func (cl *Client) TophRandom(Type string) (TophResp, error) {
	req, err := http.NewRequest("GET", "https://api.weeb.sh/images/random?type="+Type, nil)
	if err != nil {
		return TophResp{}, err
	}

	req.Header.Add("Authorization", cl.Token)
	req.Header.Add("User-Agent", cl.UserAgent)

	resp, err := client.Do(req)
	if err != nil {
		return TophResp{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return TophResp{}, err
	}

	var weebresp TophResp
	json.Unmarshal(body, &weebresp)

	return weebresp, nil
}

type Client struct {
	Token     string
	UserAgent string
}
