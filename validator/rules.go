package validator

import (
	"github.com/shopspring/decimal"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

var (
	loc, _ = time.LoadLocation("Asia/Shanghai")
)

// 判断字符是否为数字
func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// 判断字符是否为英文字符
func isAlpha(r rune) bool {

	if r >= 'A' && r <= 'Z' {
		return true
	} else if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func isPriv(s string) bool {

	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && r != '_' {
			return false
		}
	}

	return true
}

// 检测会员名
func CheckUName(str string, min, max int) bool {

	if !CtypeAlnum(str) || !First2IsAlpha(str) || !CheckStringLength(str, min, max) {
		return false
	}

	return true
}

// 检测会员密码
func CheckUPassword(str string, min, max int) bool {

	if !CtypeAlnum(str) || !CheckStringLength(str, min, max) {
		return false
	}

	return true
}

// 匹配值是否为空
func checkStr(str string) bool {

	n := len(str)
	if n <= 0 {
		return false
	}

	return true
}

// 判断是否为bool
func checkBool(str string) bool {

	_, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return true
}

// 判断是否为float
func CheckFloat(str string) bool {

	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}
	return true
}

// 判断长度
func checkLength(str string, min, max int) bool {

	if min == 0 && max == 0 {
		return true
	}

	n := len(str)
	if n < min || n > max {
		return false
	}

	return true
}

// 判断字符串长度
func CheckStringLength(val string, _min, _max int) bool {

	if _min == 0 && _max == 0 {
		return true
	}

	count := utf8.RuneCountInString(val)
	if count < _min || count > _max {

		return false
	}
	return true
}

// 判断数字范围
func CheckIntScope(s string, min, max int64) bool {

	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}

	if val < min || val > max {
		return false
	}

	return true
}

// 判断浮点范围
func CheckFloatScope(s, min, max string) (decimal.Decimal, bool) {

	fs, err := decimal.NewFromString(s)
	if err != nil {
		return fs, false
	}

	fMin, err := decimal.NewFromString(min)
	if err != nil {
		return fs, false
	}

	fMax, err := decimal.NewFromString(max)
	if err != nil {
		return fs, false
	}

	if fs.Cmp(fMin) == -1 || fMax.Cmp(fs) == -1 {
		return fs, false
	}

	return fs, true
}

// 判断是否全为数字
func CheckStringDigit(s string) bool {

	if s == "" {
		return false
	}
	for _, r := range s {
		if (r < '0' || r > '9') && r != '-' {
			return false
		}
	}
	return true
}

// 判断是否全为数字+逗号
func CheckStringCommaDigit(s string) bool {

	if s == "" {
		return false
	}
	for _, r := range s {
		if (r < '0' || r > '9') && r != ',' {
			return false
		}
	}
	return true
}

// 判断是不是中文
func CheckStringCHN(str string) bool {

	for _, r := range str {
		if !unicode.Is(unicode.Han, r) &&
			!isAlpha(r) && (r < '0' || r > '9') && r != '_' &&
			r != ' ' && r != '-' && r != '!' && r != '@' && r != ':' &&
			r != '?' && r != '+' && r != '.' && r != '/' && r != '\'' &&
			r != '(' && r != ')' && r != '·' && r != '&' {
			return false
		}
	}
	return true
}

// 判断是不是英文数字或者汉字
func CheckStringCHNAlnum(str string) bool {

	for _, r := range str {
		if !isDigit(r) && !isAlpha(r) &&
			r != ' ' && r != '-' && r != '!' && r != '_' &&
			r != '@' && r != '?' && r != '+' && r != ':' &&
			r != '.' && r != '/' && r != '(' && r != '\'' &&
			r != ')' && r != '·' && r != '&' && !unicode.Is(unicode.Han, r) {
			return false
		}
	}
	return true
}

// 判断是否module格式
func CheckStringModule(s string) bool {

	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && r != '/' {
			return false
		}
	}

	return true
}

// 判断是否全英文字母
func CheckStringAlpha(s string) bool {

	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && r != ' ' {
			return false
		}
	}

	return true
}

// 判断是否全英文字母+逗号
func CheckStringCommaAlpha(s string) bool {

	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && r != ',' {
			return false
		}
	}

	return true
}

// 判断是否全为英文字母和数字组合
func CheckStringAlnum(s string) bool {

	if s == "" {
		return false
	}
	for _, r := range s {
		if !isDigit(r) && !isAlpha(r) &&
			r != ' ' && r != '-' && r != '!' && r != '_' &&
			r != '@' && r != '?' && r != '+' && r != ':' &&
			r != '.' && r != '/' && r != '(' && r != '\'' &&
			r != ')' && r != '·' && r != '&' {
			return false
		}
	}
	return true
}

// 检查日期格式"YYYY-MM-DD"
func CheckDate(str string) bool {

	_, err := time.ParseInLocation("2006-01-02", str, loc)
	if err != nil {
		return false
	}
	return true
}

// 匹配时间 "HH:ii" or "HH:ii:ss"
func checkTime(str string) bool {

	_, err := time.ParseInLocation("15:04:05", str, loc)
	if err != nil {
		return false
	}
	return true
}

// 检查日期时间格式"YYYY-MM-DD HH:ii:ss"
func CheckDateTime(str string) bool {

	_, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if err != nil {
		return false
	}
	return true
}

func CheckMoney(money string) bool {

	// 金额小数验证
	_, err := strconv.Atoi(money)
	if err != nil {
		return false
	}
	_, err = strconv.ParseFloat(money, 64)
	if err != nil {
		return false
	}
	return true
}

// 判断字符串是不是数字
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

// 判断字符串是不是字母+数字
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

// 判断字符串是不是字母开头
func First2IsAlpha(s string) bool {

	if s == "" {
		return false
	}

	r := []rune(s)
	if len(r) < 2 {
		return false
	}

	if !isAlpha(r[0]) || !isAlpha(r[1]) {
		return false
	}

	return true
}

// 检查url
func CheckUrl(s string) bool {
	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func zip(a1, a2 []string) []string {

	r := make([]string, 2*len(a1))
	for i, e := range a1 {
		r[i*2] = e
		r[i*2+1] = a2[i]
	}

	return r
}

func FilterInjection(str string) string {

	array1 := []string{"<", ">", "&", `"`, " ", "?"}
	array2 := []string{"&lt;", "&gt;", "&amp;", "&quot;", "&nbsp;", "&iexcl;"}

	return strings.NewReplacer(zip(array1, array2)...).Replace(str)
}

func IsVietnamesePhone(phone string) bool {

	if !CtypeDigit(phone) {
		return false
	}

	r := []rune(phone)
	if len(r) == 10 {
		// 01 03 05 07 08 09开头
		if r[0] != '0' {
			return false
		}
		/*
			// 01 03 05 07 08 09开头
			if r[1] != '1' && r[1] != '3' && r[1] != '5' && r[1] != '7' && r[1] != '8' && r[1] != '9' {
				return false
			}

			switch r[1] {
			case '1':
			case '3': // 032、033、034、035、036、037、038、039
				switch r[2] {
				case '0', '1':
					return false
				}
			case '5': // 056、058、059
				switch r[2] {
				case '0', '1', '2', '3', '4', '5', '7':
					return false
				}
			case '7': // 070、079、077、076、078
				switch r[2] {
				case '1', '2', '3', '4', '5':
					return false
				}
			case '8': // 081 082 083 084 085
				switch r[2] {
				case '0', '6', '7', '8', '9':
					return false
				}
			case '9':
			default:
				return false
			}
		*/
		return true
	}

	//if len(r) == 9 {
	//	// 2 6 8 9开头
	//	if r[0] != '2' && r[0] != '6' && r[0] != '8' && r[0] != '9' {
	//		return false
	//	}
	//}

	return false
}
