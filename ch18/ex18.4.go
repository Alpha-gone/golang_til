package main

import (
	"golang_til/ch18/ex18.2/fedex"
	"golang_til/ch18/ex18.2/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}
func main() {
	korePostSender := &koreaPost.PostSender{}
	SendBook("어린 왕자", korePostSender)
	SendBook("그리스인 조르바", korePostSender)

	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)
}
