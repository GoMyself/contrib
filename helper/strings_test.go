package helper

import "testing"

func TestHide(t *testing.T) {
	t.Log(Hide("1234567890",TypePhone))
	t.Log(Hide("438953453948537",TypeBankCardNumber))
	t.Log(Hide("1PyMi4EYzGZKoxK7DozMMuoQ91EdrMMkBP",TypeVirtualCurrencyAddress))
	t.Log(Hide("邓陈芳",TypeRealName))
	t.Log(Hide("148213318@qq.com",TypeEmail))
}
