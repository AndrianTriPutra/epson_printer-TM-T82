package main

import (
	"net"
	"time"

	"github.com/hennedo/escpos"
)

func main() {
	socket, err := net.Dial("tcp", "xxx.xxx.x.xx:9100")
	if err != nil {
		println(err.Error())
	}
	defer socket.Close()

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	timestamp := now.Format(time.RFC3339)

	p := escpos.New(socket)
	p.Bold(true).Size(2, 2).Write("      Rumah Makan")
	p.LineFeed()
	p.Bold(true).Size(2, 2).Write("      Ujung Kulon")
	p.LineFeed()
	p.Bold(true).Size(2, 2).Write("")
	p.LineFeed()
	p.Bold(true).Size(1, 1).Write("          " + timestamp)
	p.LineFeed()

	p.Bold(true).Size(1, 1).Write("_____________________________________")
	p.LineFeed()
	p.Bold(true).Size(1, 1).Write("Qty|   Nama      |  Harga  | Jumlah")
	p.LineFeed()
	p.Bold(true).Size(1, 1).Write("-------------------------------------")
	p.LineFeed()
	p.Bold(true).Size(1, 1).Write("1  | Nasi        |   5000  | 5000")
	p.LineFeed()
	p.Bold(true).Size(1, 1).Write("-------------------------------------")
	p.LineFeed()
	p.Bold(true).Size(1, 1).Write("1  | Kakap Merah |  35000  | 35000")
	p.LineFeed()
	p.Bold(true).Size(1, 1).Write("-------------------------------------")
	p.LineFeed()
	p.Bold(true).Size(1, 1).Write("1  | Kelapa Muda |  10000  | 10000")
	p.LineFeed()
	p.Bold(true).Size(1, 1).Write("_____________________________________")
	p.LineFeed()
	p.Bold(true).Size(1, 1).Write("Total : 50000")
	p.LineFeed()
	p.Bold(true).Size(2, 2).Write("")
	p.LineFeed()

	p.QRCode("https://github.com/andriantriputra", true, 5, escpos.QRCodeErrorCorrectionLevelH)
	p.PrintAndCut()
}
