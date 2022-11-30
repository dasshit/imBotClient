package imBotClient

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fastjson"
)

type Client struct {
	Token      string
	ApiBaseUrl string `default:"https://api.icq.net/bot/v1"`
	ParseMode  string `default:"HTML"`
	HttpClient *fasthttp.Client
	JsonParser fastjson.Parser
}

func NewClient(token string) Client {
	client := Client{}
	client.Token = token
	client.HttpClient = &fasthttp.Client{}
	return client
}

func (bot Client) GetReq(path string) {

}

func (bot Client) SelfGet() *fastjson.Value {

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(bot.ApiBaseUrl + "/self/get")

	query := req.URI().QueryArgs()
	query.Add("token", bot.Token)

	resp := fasthttp.AcquireResponse()
	err := bot.HttpClient.Do(req, resp)
	if err != nil {
		panic(err)
	}

	jsonBody, err := bot.JsonParser.ParseBytes(resp.Body())
	if err != nil {
		panic(err)
	}

	return jsonBody
}
