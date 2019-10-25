package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/micmonay/keybd_event"
	"go.bug.st/serial.v1"
)

func main() {

	kb, err := keybd_event.NewKeyBonding() //teclado virtual
	if err != nil {
		panic(err)
	}

	mode := &serial.Mode{BaudRate: 9600}           //taxa de trasmissão do monitor serial
	port, err := serial.Open("/dev/ttyUSB1", mode) //porta serial
	if err != nil {
		fmt.Println(err)
	}

	buff := make([]byte, 100)
	for {
		n, err := port.Read(buff) //valor lido na porta serial
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
		//fmt.Printf("%v", string(buff[:n])) // print do valor na porta serial
		//se o valor da porta serial for A vai entrar no if
		if strings.Contains(strings.Trim(string(buff[:n]), "\n"), "A") {
			kb.SetKeys(keybd_event.VK_A) //digita a letra A no teclado e o mesmo para cada letra posterior
		}
		if strings.Contains(strings.Trim(string(buff[:n]), "\n"), "B") {
			kb.SetKeys(keybd_event.VK_S) //no caso letra S
		}
		if strings.Contains(strings.Trim(string(buff[:n]), "\n"), "C") {
			kb.SetKeys(keybd_event.VK_D)
		}
		if strings.Contains(strings.Trim(string(buff[:n]), "\n"), "D") {
			kb.SetKeys(keybd_event.VK_F)
		}
		if strings.Contains(strings.Trim(string(buff[:n]), "\n"), "E") {
			kb.SetKeys(keybd_event.VK_G)
		}
		if strings.Contains(strings.Trim(string(buff[:n]), "\n"), "F") {
			kb.SetKeys(keybd_event.VK_H)
		}

		err = kb.Launching() //inicia o teclado
		if err != nil {
			panic(err)
		}
	}

	port.Close() //fecha a porta
}
