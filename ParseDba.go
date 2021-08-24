package ParseDBA

import (
	//"fmt"
	//    	"os"
	//	"bytes"
	//	"encoding/binary"
	"bytes"
	"io/ioutil"
	"strings"
	//	"encoding/binary"
	"encoding/json"
	//	"unicode/utf8"

	//	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

)

//var SQLKey []byte = []byte("19465912879oiuxc ensdfaiuo3i73798kjl")

type ConnectInfo struct {
	Server string
	DB     string
	User   string `json:"UID"`
	Psw    string `json:"PWD"`
	Crc    string `json:"Checksum"`
}

func ConvToUTF8(b []byte, e encoding.Encoding) (result []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(b), unicode.BOMOverride(e.NewDecoder()))
	return ioutil.ReadAll(reader)
}

func XOR(ParA byte, ParB byte) byte {
	var Res byte
	var BitA byte
	var BitB byte
	var BitC byte
	var Koef int = 1
	Res = byte(0)

	for BitNumber := 1; BitNumber <= 8; BitNumber++ {
		BitA = byte(int(ParA) % 2)
		BitB = byte(int(ParB) % 2)
		ParA = byte(int(ParA) / 2)
		ParB = byte(int(ParB) / 2)
		BitC = byte((int(BitA) + int(BitB)) % 2)
		Res = byte(int(Res) + int(BitC)*Koef)
		Koef = Koef * 2
	}

	return Res
}

func (cn *ConnectInfo) ParseConnect(data_ string) (string, error) {

	data = []byte(data_)

	var SQLKey []byte = []byte("19465912879oiuxc ensdfaiuo3i73798kjl")
	nlen := len(data)
	Connect := ""
	for i := 0; i < nlen; i++ {
		Connect = Connect + string(XOR(data[i], SQLKey[i%36]))
	}

	out := strings.Replace(Connect, "\",\"", "\":\"", -1)
	out = strings.Replace(out, "\"},{\"", "\",\"", -1)
	out = strings.Replace(out, "{{", "{", -1)
	out = strings.Replace(out, "}}", "}", -1)
	data = []byte(out)
	err = json.Unmarshal(data, cn)

	return Connect, err
}

func (cn *ConnectInfo) ParseDBA(fileName string) error {

	data, err := ioutil.ReadFile(fileName)
	// Если во время считывания файла произошла ошибка
	// ничего не делаем
	if err != nil {
		//fmt.Println(err)
		return err
	}

	var SQLKey []byte = []byte("19465912879oiuxc ensdfaiuo3i73798kjl")
	nlen := len(data)
	Connect := ""
	for i := 0; i < nlen; i++ {
		Connect = Connect + string(XOR(data[i], SQLKey[i%36]))
	}

	out := strings.Replace(Connect, "\",\"", "\":\"", -1)
	out = strings.Replace(out, "\"},{\"", "\",\"", -1)
	out = strings.Replace(out, "{{", "{", -1)
	out = strings.Replace(out, "}}", "}", -1)
	data = []byte(out)
	err = json.Unmarshal(data, cn)

	return err
}

/*
func main() {

	data, err := ioutil.ReadFile("1Cv7.DBA")
	// Если во время считывания файла произошла ошибка
	// выводим ее
	if err != nil {
		fmt.Println(err)
	}
	SQLKey := []byte("19465912879oiuxc ensdfaiuo3i73798kjl")
	/*
		b := []byte("19465912879oiuxc ensdfaiuo3i73798kjl")
		for len(b) > 0 {
			//r, size := utf8.DecodeLastRune(b)
			_, size := utf8.DecodeLastRune(b)
			//fmt.Printf("%d %c %d %c %v\n", b[len(b)-1], b[len(b)-1], r, r, size)

			b = b[:len(b)-size]
		}
*/ /*
	//fmt.Print(string(data))
	fmt.Printf("Длина файла %d\n", len(data))
	nlen := len(data)
	Connect := ""
	for i := 0; i < nlen; i++ {
		Connect = Connect + string(XOR(data[i], SQLKey[i%36]))
	}
	fmt.Println(Connect)

	out := strings.Replace(Connect, "\",\"", "\":\"", -1)
	out = strings.Replace(out, "\"},{\"", "\",\"", -1)
	out = strings.Replace(out, "{{", "{", -1)
	out = strings.Replace(out, "}}", "}", -1)
	fmt.Println(out)

	var cn = ConnectInfo{"192.168.101.2", "Analiz_EN", "sa", "1551", "1234"}
	data, err = json.Marshal(cn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", data)
	data = []byte(out)
	err = json.Unmarshal(data, &cn)

	cnv := new(ConnectInfo)
	cnv.ParseDBA("1Cv7.DBA")
}
*/
