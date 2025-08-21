package main

import (
	"golang_til/ch18/ex18.2/fedex"
	"golang_til/ch18/ex18.2/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}
func main() {
	sender := &koreaPost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
