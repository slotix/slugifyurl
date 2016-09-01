package slugifyurl

import (
	"fmt"
	"testing"
)

func TestSlugifyURL(t *testing.T) {

	options := Options{
		SlashChar:    "-",
		MaxLength:    50,
		SkipScheme:   true,
		SkipUserinfo: true,
		UnixOnly:     false,
	}

	fmt.Println(Slugify("http://www.example.com///cgi-bin///user.pl?user=admin///", options))
	fmt.Println(Slugify("https://admin:test@www.example.com/", options))
	fmt.Println(Slugify("https://www.example.com?q=<>:\"/\\|?*A", options))
	fmt.Println(Slugify("https://www.example.com/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/", options))

}
