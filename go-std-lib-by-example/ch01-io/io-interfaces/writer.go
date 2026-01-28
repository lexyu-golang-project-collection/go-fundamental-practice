package main

import (
	"bytes"
	"fmt"
	"os"
)

func ByteWriterExample() {
TOP:
	for {
		fmt.Println("請輸入要透過`WriteByte`寫入的一個`ASCII`字元(b: Return to Top; q: exit) :")
		var ch byte
		fmt.Scanf("%c\n", &ch)
		switch ch {
		case 'b':
			fmt.Println("Return Top")
			break TOP
		case 'q':
			fmt.Println("exit !")
			os.Exit(0)
		default:
			buf := new(bytes.Buffer)
			err := buf.WriteByte(ch)
			if err == nil {
				fmt.Println("寫入一個位元組成功! 準備讀取該位元組...")
				newCh, _ := buf.ReadByte()
				fmt.Println("讀取的位元組: ", newCh)
			} else {
				fmt.Println("寫入錯誤")
			}
		}
	}
}
