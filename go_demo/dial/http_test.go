package dial

import "testing"

func TestHttpGet(t *testing.T) {
	url := "https://www.imooc.com/"
	s := NewSpider()
	resp, err := s.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
