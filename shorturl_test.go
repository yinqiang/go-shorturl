package shorturl

import (
	"testing"
)

func Test_Short(t *testing.T) {
	t.Log("cn.bing.com", Short("cn.bing.com"))
	t.Log("cn.bing.com", Short("cn.bing.com"))
	t.Log("cn1.bing.com", Short("cn1.bing.com"))
	t.Log("cn2.bing.com", Short("cn2.bing.com"))
	t.Log("cn.bing.com/v1", Short("cn.bing.com/v1"))
	t.Log("cn.bing.com/v2", Short("cn.bing.com/v2"))
	t.Log("en.bing.com", Short("en.bing.com"))
	t.Log("www.bing.com", Short("www.bing.com"))
	t.Log("bing.com", Short("bing.com"))
}
