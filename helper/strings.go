package helper

import (
	"encoding/hex"
	"fmt"
	"github.com/minio/md5-simd"
	"github.com/shopspring/decimal"
	"github.com/valyala/fasthttp"
	"github.com/wI2L/jettison"
	"lukechampine.com/frand"
	"strconv"
	"strings"
	"time"
)

var (
	o = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
	n = []string{
		"&#65;", "&#66;", "&#67;", "&#68;", "&#69;", "&#70;", "&#71;", "&#72;", "&#73;", "&#74;", "&#75;", "&#76;", "&#77;",
		"&#78;", "&#79;", "&#80;", "&#81;", "&#82;", "&#83;", "&#84;", "&#85;", "&#86;", "&#87;", "&#88;", "&#89;", "&#90;",
		"&#97;", "&#98;", "&#99;", "&#100;", "&#101;", "&#102;", "&#103;", "&#104;", "&#105;", "&#106;", "&#107;", "&#108;", "&#109;",
		"&#110;", "&#111;", "&#112;", "&#113;", "&#114;", "&#115;", "&#116;", "&#117;", "&#118;", "&#119;", "&#120;", "&#121;", "&#122;",
	}
	so = []string{
		"<", ">", "&", " ", "(", ")",
	}
	sn = []string{
		"&lt;", "&gt;", "&amp;", "&quot;", "&#40;", "&#40;",
	}
)

func init() {
	genWords("insert")
	genWords("update")
	genWords("delete")
	genWords("alter")
	genWords("table")
	genWords("create")
	genWords("database")
	genWords("user")
	genWords("drop")
	genWords("truncate")
	genWords("set")
	genWords("values")
}

const (
	TYpeNone = iota
	TypePhone
	TypeBankCardNumber
	TypeVirtualCurrencyAddress
	TypeRealName
	TypeEmail
)

type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func Print(ctx *fasthttp.RequestCtx, state bool, data interface{}) {

	ctx.SetStatusCode(200)
	ctx.SetContentType("application/json")

	res := Response{
		Status: state,
		Data:   data,
	}

	bytes, err := jettison.Marshal(res)
	if err != nil {
		ctx.SetBody([]byte(err.Error()))
		return
	}

	ctx.SetBody(bytes)
}

func PrintJson(ctx *fasthttp.RequestCtx, state bool, data string) {

	ctx.SetStatusCode(200)
	ctx.SetContentType("application/json")

	builder := strings.Builder{}

	builder.WriteString(`{"status":`)
	builder.WriteString(strconv.FormatBool(state))
	builder.WriteString(`,"data":`)
	builder.WriteString(data)
	builder.WriteString("}")

	ctx.SetBody([]byte(builder.String()))
}

func GenId() string {
	return fmt.Sprintf("%d%d", Cputicks(), frand.Uint64n(20))
}

//判断字符是否为数字
func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

//判断字符是否为英文字符
func isAlpha(r rune) bool {

	if r >= 'A' && r <= 'Z' {
		return true
	} else if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

//判断字符串是不是数字
func CtypeDigit(s string) bool {

	if s == "" {
		return false
	}
	for _, r := range s {
		if !isDigit(r) {
			return false
		}
	}
	return true
}

//判断字符串是不是字母+数字
func CtypeAlnum(s string) bool {

	if s == "" {
		return false
	}
	for _, r := range s {
		if !isDigit(r) && !isAlpha(r) {
			return false
		}
	}
	return true
}

// 简单隐藏信息
func Hide(s string, flag int) string {

	switch flag {
	case TypePhone: //电话号码
		if len(s)  <= 6{
			return s
		}

		runes := []rune(s)
		l := len(runes)
		for i := 3; i < l-3; i++ {
			runes[i] = '*'
		}

		return string(runes)
	case TypeBankCardNumber,TypeVirtualCurrencyAddress: //银行卡 虚拟币地址
		if len(s)  <= 8{
			return s
		}

		runes := []rune(s)
		l := len(runes)
		for i := 4; i < l-4; i++ {
			runes[i] = '*'
		}

		return string(runes)
	case TypeRealName: //真实姓名
	    // 越南语姓名
	    n := strings.IndexAny(s, " ")
		if n != -1{
			runes := []rune(s)
			l := len(runes)
			for i := n + 1; i < l; i++ {
				runes[i] = '*'
			}

			return string(runes)
		}

		runes := []rune(s)
		l := len(runes)
		for i := 1; i < l; i++ {
			runes[i] = '*'
		}

		return string(runes)
	case TypeEmail: 	// 邮箱
		if !strings.Contains(s, "@") {
			return s
		}

		r := strings.Split(s, "@")
		runes := []rune(r[0])
		l := len(runes)
		if l <= 2 {
			return string(runes[:1]) + "*" + "@" + r[1]
		}

		for i := 2; i < l; i++ {
			runes[i] = '*'
		}

		return string(runes) + "@" + r[1]
	}

	return s
}

//获取source的子串,如果start小于0或者end大于source长度则返回""
//start:开始index，从0开始，包括0
//end:结束index，以end结束，但不包括end
func Substring(source string, start int, end int) string {

	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}

//字符串特殊字符转译
func Addslashes(str string) string {

	tmpRune := []rune{}
	strRune := []rune(str)
	for _, ch := range strRune {
		switch ch {
		case []rune{'\\'}[0], []rune{'"'}[0], []rune{'\''}[0]:
			tmpRune = append(tmpRune, []rune{'\\'}[0])
			tmpRune = append(tmpRune, ch)
		default:
			tmpRune = append(tmpRune, ch)
		}
	}
	return string(tmpRune)
}

func GetMD5Hash(text string) string {

	server := md5simd.NewServer()
	md5Hash := server.NewHash()

	md5Hash.Write([]byte(text))

	digest := md5Hash.Sum([]byte{})
	encrypted := hex.EncodeToString(digest)

	server.Close()
	md5Hash.Close()

	return encrypted
}

func StrToTime(value string, loc *time.Location) time.Time {

	if value == "" {
		return time.Time{}
	}
	layouts := []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	var t time.Time
	var err error
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	for _, layout := range layouts {
		t, err = time.ParseInLocation(layout, value, loc)
		if err == nil {
			return t
		}
	}
	return t
}

func TrimStr(val decimal.Decimal) string {

	s := "0.000"
	sDigit := strings.Split(val.String(), ".")
	if len(sDigit) != 2 {
		if len(sDigit) == 1 && CtypeDigit(sDigit[0]) {
			return sDigit[0] + ".000"
		}
		return s
	}

	// 浮点位数校验
	if len(sDigit[1]) <= 3 {
		return val.String()
	}

	return sDigit[0] + "." + sDigit[1][:3]
}

func zip(old, new []string) []string {

	r := make([]string, 2*len(old))
	for i, e := range old {
		r[i*2] = e
		r[i*2+1] = new[i]
	}

	return r
}

func genWords(s string) {

	fmt.Println("gen worlds: " + s)
	so = append(so, s)
	sn = append(sn, replace(s))
	r := []rune(s)
	l := len(r)
	for i := 0; i < l; i++ {
		up(s, i)
	}
}

func replace(s string) string {
	return strings.NewReplacer(zip(o, n)...).Replace(s)
}

func up(s string, i int) {

	r := []rune(s)
	l := len(r)
	if i > l {
		return
	}

	for j := i; j < l; j++ {
		if r[j] >= 'a' && r[j] <= 'z' {
			r[j] -= 32
			so = append(so, string(r))
			sn = append(sn, replace(string(r)))
		}
	}
}

func FilterInjection(s string) string {

	return strings.NewReplacer(zip(so, sn)...).Replace(s)
}