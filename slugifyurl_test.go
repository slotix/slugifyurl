package slugifyurl

import (
	"log"
	"testing"
)


type testCase struct {
	input, expect string
}

func testSlugify(t *testing.T, input, expect string) {
	options := Options{
		SlashChar:    "-",
		MaxLength:    50,
		SkipScheme:   true,
		SkipUserinfo: true,
		UnixOnly:     false,
	}

	ret := Slugify(input, options)
	check(t, ret, expect)
}

func check(t *testing.T, ret, expect string) {
	if ret != expect {
		t.Errorf("Expected %v, got %v", expect, ret)
	}
}

func TestSlugifyURL(t *testing.T) {

	
	cases := []testCase {
		{"http://www.example.com///cgi-bin///user.pl?user=admin///","www.example.com-cgi-bin-user.pl-user=admin"},
		{"https://admin:test@www.example.com/", "www.example.com"},
		{"https://www.example.com?q=<>:\"/\\|?*A", "www.example.com-q=-A"},
		{"https://www.example.com/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/", "www.example.com-extremely-long-url-extremely-long-"},
	}
	for _, c := range cases {
		log.Printf("input %v, expect %v", c.input, c.expect)
		testSlugify(t, c.input, c.expect)
	}
	//fmt.Println(Slugify("http://www.example.com///cgi-bin///user.pl?user=admin///", options))
	//fmt.Println(Slugify("https://admin:test@www.example.com/", options))
	//fmt.Println(Slugify("https://www.example.com?q=<>:\"/\\|?*A", options))
	//fmt.Println(Slugify("https://www.example.com/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/", options))

}
