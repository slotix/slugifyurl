package slugifyurl

import (
	"fmt"
	"regexp"
)

//Options structure for slugification parameters initialization 
type Options struct {
	SlashChar    string //Used to replace slashes
	MaxLength    int    //Output string length limit
	SkipScheme   bool   //Omit scheme http(s)://
	SkipUserinfo bool   //Omit username and password information
	UnixOnly     bool   //if true it doesn't check invalid characters for Windows file names
}

// Slugify modifies URL to sanitized string which can be used, for instance, as a filename.
/*
1. Protocol and Userinfo could be omitted  (http:// or https:// and <user>:<password>@ )
2. Slashes are replaced with slashChar character, for example "-"
3. Resulted string is limited to maxLength options parameter

Examples
 - https://admin:test@www.example.com/ --> www.example.com
 - http://www.example.com///cgi-bin///user.pl?user=admin/// --> www.example.com-cgi-bin-user.pl-user=admin
 - https://www.example.com?q=<>:\"/\\|?*A --> www.example.com-q=-A
 - https://www.example.com/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/ -->
  www.example.com-extremely-long-url-extremely-long-

*/
func Slugify(rawurl string, options Options) string {
	sanitized := rawurl
    //if true remove scheme information from result
	if options.SkipScheme {
		re := regexp.MustCompile(`^[\w]+://`)
		sanitized = re.ReplaceAllString(sanitized, "")
	}
    //if true remove username and password information from result
	if options.SkipUserinfo {
		re := regexp.MustCompile(`[\w\-_\.]+:[^@]+?@`)
		sanitized = re.ReplaceAllString(sanitized, "")
	}
    //Replace all slashes with SlashChar
	re := regexp.MustCompile(`[/]`)
	sanitized = re.ReplaceAllString(sanitized, options.SlashChar)
	// replace invalid chars for Windows file names with slashChar
	if !options.UnixOnly {
		//according to MS specs at https://msdn.microsoft.com/en-us/library/aa365247(VS.85)
        //<>:\"/\\|?* These characters are not eligable for usage in Windows file names
		re := regexp.MustCompile(`[<>:\"/\\|?*]`)
		sanitized = re.ReplaceAllString(sanitized, options.SlashChar)
	}
	if len(options.SlashChar) > 0 {
        // replace multiple slashChars with just one
		re := regexp.MustCompile(fmt.Sprintf(`['\%s']{2,}`, options.SlashChar))
		sanitized = re.ReplaceAllString(sanitized, options.SlashChar)
        //remove trailing slashChars
		re = regexp.MustCompile(fmt.Sprintf(`['\%s']*$`, options.SlashChar))
		sanitized = re.ReplaceAllString(sanitized, "")
	}
	//truncate to maxLength
	if len(sanitized) > options.MaxLength {
		sanitized = sanitized[:options.MaxLength]
	}
	return sanitized
}
