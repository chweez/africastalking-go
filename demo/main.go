package main 

const (
	username = ""	//Your Africa's Talking Username
	apiKey = ""		//Production or Sandbox API Key
	option = ""		// Choose either Sandbox or Production
)

func main (){
	gateway, err := sms.NewGateway(username, apiKey, option)
	if err != nil {
		log.Fatal(err)
	}
}