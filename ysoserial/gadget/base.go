package gadget

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strings"
)

func generateCmd(cmd string) string {
	finalCmd := bytes.Buffer{}
	finalCmd.WriteString("74")
	cmdLenByte := make([]byte, 2)
	binary.BigEndian.PutUint16(cmdLenByte, uint16(len(cmd)))
	cmdLenStr := strings.ToUpper(hex.EncodeToString(cmdLenByte))
	data := strings.ToUpper(hex.EncodeToString([]byte(cmd)))
	finalCmd.WriteString(cmdLenStr)
	finalCmd.WriteString(data)
	return finalCmd.String()
}
