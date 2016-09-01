SlugifyURL
----------
[![Build Status](https://travis-ci.org/slotix/slugifyurl.svg?branch=master)](https://travis-ci.org/slotix/slugifyurl)
[![Go Report Card](https://goreportcard.com/badge/github.com/slotix/slugifyurl)](https://goreportcard.com/report/github.com/slotix/slugifyurl)
Slugify modifies URL to a sanitized string which can be used, for instance, as a filename.

Description
-----------
1. Protocol and Userinfo could be omitted  (http:// or https:// and user:password@ )
2. Slashes are replaced with slashChar character, for example "-"
3. Resulted string is limited to maxLength options parameter

Examples
--------
 - https://admin:test@www.example.com/ --> www.example.com
 - http://www.example.com///cgi-bin///user.pl?user=admin/// --> www.example.com-cgi-bin-user.pl-user=admin
 - https://www.example.com?q=<>:\"/\\|?*A --> www.example.com-q=-A
 - https://www.example.com/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/ -->
  www.example.com-extremely-long-url-extremely-long-

Installation 
------------

```
go get -u github.com/slotix/slugifyurl
```

Usage
-----

```go
package main
import (
    "fmt"
    slugify "github.com/slotix/slugifyurl"
)

func main() {
    options := slug.Options{
		SlashChar:    "-",      //Used to replace slashes
		MaxLength:    50,       //Output string length limit 
		SkipScheme:   true,     //Omit scheme http(s)://
		SkipUserinfo: true,     //Omit username and password information
		UnixOnly:     false,    //Vlidate file names for Windows OS.    
	}
    url := "http://admin:test@www.example.com///cgi-bin///user.pl?user=admin///"
    
    fmt.Println(slugify(url))
    //Output: www.example.com-cgi-bin-user.pl-user=admin
}
