package main

import (
    "net/http"
    "time"
    "log"
    "net/smtp"
    tb "gopkg.in/tucnak/telebot.v2"
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type Config struct {
    Bar []string
}


var (
  count = 0
  http_address = ""
  instance_name = ""
  max_allowed = 10
  inaccessible = 0
  config Config
)


func main() {
  doEvery(2 * time.Second, pinger)
}

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func pinger(t time.Time) {
  filename := "instanceinfo.yml"
   source, err := ioutil.ReadFile(filename)
   if err != nil {
       panic(err)
   }
   err = yaml.Unmarshal(source, &config)
   if err != nil {
       panic(err)
   }
   http_address =  config.Bar[0] //+  "/images/newlayout/all_content.png" //server_url
   instance_name =  config.Bar[1] //server_name


  resp, err := http.Get(http_address)
  var number = 0
  if err != nil {
      print(err.Error())
      counter(number + 1)
  } else {
      print(string(resp.StatusCode) + resp.Status)
      counter(number)
  }
}

func counter(number int){
  if (number == 1) {
    count = count + number
    print("\n")
    print(count)
    print("\n")
    if (count >= max_allowed){
      print("\n")
      print("\n")
      print("\n")
      send("our Server is" + instance_name + " not available!!!!")
      sendtg()
      print("\n")
      print("\n")
      print("\n")
    }
  } else {
    count = 0
    max_allowed = 10
    print("\n")
    print(count)
    print("\n")
  }
}

func send(body string) {
	from := "test@gmail.com"
	pass := "testpassword"
	to := config.Bar[3]
  	cc := config.Bar[2]

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
    "CC: " + cc + "\n" +
		"Subject: Server NOT AVAILABLE\n\n" +
		body +
    "\n\ntried to ping: " + http_address

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	count = 0
  max_allowed = max_allowed + 10
}

func sendtg() {

  //var error_text = "Oh our server is not reachable"
  errormsg := "oohw our Server: " + instance_name + " is not reachable"
	b, err := tb.NewBot(tb.Settings{
		Token:  "URTELEGRAMTOKEN",
		URL: ("https://api.telegram.org/BOTID/sendMessage?chat_id=URCHARTID&text=" + errormsg),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
	})

}
