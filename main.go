package imBotClient

func main() {
	client := NewClient("001.2177161926.0690735952:1000001281")
	client.ApiBaseUrl = "https://api.internal.myteam.mail.ru/bot/v1"
	client.selfGet()
}
