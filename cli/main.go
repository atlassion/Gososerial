package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	gososerial "github.com/EmYiQing/Gososerial"
	"github.com/EmYiQing/Gososerial/log"
	"github.com/EmYiQing/Gososerial/tool"
	"github.com/EmYiQing/Gososerial/ysoserial/gadget"
)

const (
	version = "0.1"
	author  = "4ra1n"
)

func main() {
	log.PrintLogo(version, author)
	var payload string
	var command string
	var list bool
	flag.StringVar(&payload, "payload", "", "use which payload")
	flag.StringVar(&command, "cmd", "", "command")
	flag.BoolVar(&list, "list", false, "show payload list")
	flag.Parse()
	if list {
		log.Info("payload list: ")
		all := gososerial.GetAllNames()
		for _, v := range all {
			fmt.Printf("\t%s\n", v)
		}
		return
	}
	if command == "" || payload == "" {
		log.Error("error input")
		return
	}
	fmt.Println(command)
	switch payload {
	case gadget.CB1:
		log.Info("get payload: %s", gadget.CB1)
		bytePayload := gososerial.GetCB1(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC1:
		log.Info("get payload: %s", gadget.CC1)
		bytePayload := gososerial.GetCC1(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC2:
		log.Info("get payload: %s", gadget.CC2)
		bytePayload := gososerial.GetCC2(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC3:
		log.Info("get payload: %s", gadget.CC3)
		bytePayload := gososerial.GetCC3(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC4:
		log.Info("get payload: %s", gadget.CC4)
		bytePayload := gososerial.GetCC4(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC5:
		log.Info("get payload: %s", gadget.CC5)
		bytePayload := gososerial.GetCC5(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC6:
		log.Info("get payload: %s", gadget.CC6)
		bytePayload := gososerial.GetCC6(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CC7:
		log.Info("get payload: %s", gadget.CC7)
		bytePayload := gososerial.GetCC7(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CCK1:
		log.Info("get payload: %s", gadget.CCK1)
		bytePayload := gososerial.GetCCK1(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CCK2:
		log.Info("get payload: %s", gadget.CCK2)
		bytePayload := gososerial.GetCCK2(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CCK3:
		log.Info("get payload: %s", gadget.CCK3)
		bytePayload := gososerial.GetCCK3(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	case gadget.CCK4:
		log.Info("get payload: %s", gadget.CCK4)
		bytePayload := gososerial.GetCCK4(command)
		output := tool.FormatPayload(hex.EncodeToString(bytePayload))
		log.Info("payload: %s", output)
	default:
		log.Error("error payload")
		return
	}
}
