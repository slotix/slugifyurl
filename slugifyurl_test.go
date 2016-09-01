package slugifyurl

import (
	"fmt"
	slug "github.com/slotix/slugifyurl"
	"testing"
)

func TestSlugifyURL(t *testing.T) {
	
	options := slug.Options{
		SlashChar:    "-",
		MaxLength:    50,
		SkipScheme:   true,
		SkipUserinfo: true,
		UnixOnly:     false,
	}

	fmt.Println(slug.Slugify("http://www.example.com///cgi-bin///user.pl?user=admin///", options))
	fmt.Println(slug.Slugify("https://admin:test@www.example.com/", options))
	fmt.Println(slug.Slugify("https://www.example.com?q=<>:\"/\\|?*A", options))
    fmt.Println(slug.Slugify("https://www.example.com/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/", options))
    
}
