package imBotClient

import "github.com/valyala/fasthttp"

type Client struct {
	Token      string
	ApiBaseUrl string `default:"https://api.icq.net/bot/v1"`
	ParseMode  string `default:"HTML"`
	HttpClient *fasthttp.Client
}

func NewClient(token string) Client {
	client := Client{}
	client.Token = token
	client.HttpClient = &fasthttp.Client{}
	return client
}

func (bot Client) selfGet() string {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(bot.ApiBaseUrl + "/self/get")

	resp := fasthttp.AcquireResponse()
	err := bot.HttpClient.Do(req, resp)
	if err != nil {
		panic(err)
	}

	bodyBytes := resp.Body()
	return string(bodyBytes)
}
