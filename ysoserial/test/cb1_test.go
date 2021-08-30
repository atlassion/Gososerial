package test

import (
	"encoding/hex"
	"fmt"
	gososerial "github.com/EmYiQing/Gososerial"
	"strings"
	"testing"
)

func TestCB1(t *testing.T) {
	ser := gososerial.GetCB1("calc.exe")
	serStr := strings.ToUpper(hex.EncodeToString(ser))
	stdStr := "ACED0005737200176A6176612E7574696C2E5072696F72697479517565756594DA30B4FB3F82B103000249000473697A654C000A636F6D70617261746F727400164C6A6176612F7574696C2F436F6D70617261746F723B7870000000027372002B6F72672E6170616368652E636F6D6D6F6E732E6265616E7574696C732E4265616E436F6D70617261746F72E3A188EA7322A4480200024C000A636F6D70617261746F7271007E00014C000870726F70657274797400124C6A6176612F6C616E672F537472696E673B78707372003F6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E732E636F6D70617261746F72732E436F6D70617261626C65436F6D70617261746F72FBF49925B86EB13702000078707400106F757470757450726F706572746965737704000000037372003A636F6D2E73756E2E6F72672E6170616368652E78616C616E2E696E7465726E616C2E78736C74632E747261782E54656D706C61746573496D706C09574FC16EACAB3303000649000D5F696E64656E744E756D62657249000E5F7472616E736C6574496E6465785B000A5F62797465636F6465737400035B5B425B00065F636C6173737400125B4C6A6176612F6C616E672F436C6173733B4C00055F6E616D6571007E00044C00115F6F757470757450726F706572746965737400164C6A6176612F7574696C2F50726F706572746965733B787000000000FFFFFFFF757200035B5B424BFD19156767DB37020000787000000002757200025B42ACF317F8060854E00200007870000006A0CAFEBABE0000003200390A0003002207003707002507002601001073657269616C56657273696F6E5549440100014A01000D436F6E7374616E7456616C756505AD2093F391DDEF3E0100063C696E69743E010003282956010004436F646501000F4C696E654E756D6265725461626C650100124C6F63616C5661726961626C655461626C6501000474686973010013537475625472616E736C65745061796C6F616401000C496E6E6572436C61737365730100354C79736F73657269616C2F7061796C6F6164732F7574696C2F4761646765747324537475625472616E736C65745061796C6F61643B0100097472616E73666F726D010072284C636F6D2F73756E2F6F72672F6170616368652F78616C616E2F696E7465726E616C2F78736C74632F444F4D3B5B4C636F6D2F73756E2F6F72672F6170616368652F786D6C2F696E7465726E616C2F73657269616C697A65722F53657269616C697A6174696F6E48616E646C65723B2956010008646F63756D656E7401002D4C636F6D2F73756E2F6F72672F6170616368652F78616C616E2F696E7465726E616C2F78736C74632F444F4D3B01000868616E646C6572730100425B4C636F6D2F73756E2F6F72672F6170616368652F786D6C2F696E7465726E616C2F73657269616C697A65722F53657269616C697A6174696F6E48616E646C65723B01000A457863657074696F6E730700270100A6284C636F6D2F73756E2F6F72672F6170616368652F78616C616E2F696E7465726E616C2F78736C74632F444F4D3B4C636F6D2F73756E2F6F72672F6170616368652F786D6C2F696E7465726E616C2F64746D2F44544D417869734974657261746F723B4C636F6D2F73756E2F6F72672F6170616368652F786D6C2F696E7465726E616C2F73657269616C697A65722F53657269616C697A6174696F6E48616E646C65723B29560100086974657261746F720100354C636F6D2F73756E2F6F72672F6170616368652F786D6C2F696E7465726E616C2F64746D2F44544D417869734974657261746F723B01000768616E646C65720100414C636F6D2F73756E2F6F72672F6170616368652F786D6C2F696E7465726E616C2F73657269616C697A65722F53657269616C697A6174696F6E48616E646C65723B01000A536F7572636546696C6501000C476164676574732E6A6176610C000A000B07002801003379736F73657269616C2F7061796C6F6164732F7574696C2F4761646765747324537475625472616E736C65745061796C6F6164010040636F6D2F73756E2F6F72672F6170616368652F78616C616E2F696E7465726E616C2F78736C74632F72756E74696D652F41627374726163745472616E736C65740100146A6176612F696F2F53657269616C697A61626C65010039636F6D2F73756E2F6F72672F6170616368652F78616C616E2F696E7465726E616C2F78736C74632F5472616E736C6574457863657074696F6E01001F79736F73657269616C2F7061796C6F6164732F7574696C2F476164676574730100083C636C696E69743E0100116A6176612F6C616E672F52756E74696D6507002A01000A67657452756E74696D6501001528294C6A6176612F6C616E672F52756E74696D653B0C002C002D0A002B002E01000863616C632E65786508003001000465786563010027284C6A6176612F6C616E672F537472696E673B294C6A6176612F6C616E672F50726F636573733B0C003200330A002B003401000D537461636B4D61705461626C6501001F79736F73657269616C2F50776E6572313133393534363630353630363030300100214C79736F73657269616C2F50776E6572313133393534363630353630363030303B002100020003000100040001001A000500060001000700000002000800040001000A000B0001000C0000002F00010001000000052AB70001B100000002000D0000000600010000002F000E0000000C000100000005000F003800000001001300140002000C0000003F0000000300000001B100000002000D00000006000100000034000E00000020000300000001000F0038000000000001001500160001000000010017001800020019000000040001001A00010013001B0002000C000000490000000400000001B100000002000D00000006000100000038000E0000002A000400000001000F003800000000000100150016000100000001001C001D000200000001001E001F00030019000000040001001A00080029000B0001000C00000024000300020000000FA70003014CB8002F1231B6003557B1000000010036000000030001030002002000000002002100110000000A000100020023001000097571007E0010000001D4CAFEBABE00000032001B0A0003001507001707001807001901001073657269616C56657273696F6E5549440100014A01000D436F6E7374616E7456616C75650571E669EE3C6D47180100063C696E69743E010003282956010004436F646501000F4C696E654E756D6265725461626C650100124C6F63616C5661726961626C655461626C6501000474686973010003466F6F01000C496E6E6572436C61737365730100254C79736F73657269616C2F7061796C6F6164732F7574696C2F4761646765747324466F6F3B01000A536F7572636546696C6501000C476164676574732E6A6176610C000A000B07001A01002379736F73657269616C2F7061796C6F6164732F7574696C2F4761646765747324466F6F0100106A6176612F6C616E672F4F626A6563740100146A6176612F696F2F53657269616C697A61626C6501001F79736F73657269616C2F7061796C6F6164732F7574696C2F47616467657473002100020003000100040001001A000500060001000700000002000800010001000A000B0001000C0000002F00010001000000052AB70001B100000002000D0000000600010000003C000E0000000C000100000005000F001200000002001300000002001400110000000A000100020016001000097074000450776E72707701007871007E000D78"
	fmt.Println(serStr)
	fmt.Println(stdStr)
	fmt.Println(serStr == stdStr)
}
