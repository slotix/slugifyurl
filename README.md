# SlugifyURL
Slugify modifies URL to a sanitized string which can be used, for instance, as a filename.

# Description
1. Protocol and Userinfo could be omitted  (http:// or https:// and <user>:<password>@ )
2. Slashes are replaced with slashChar character, for example "-"
3. Resulted string is limited to maxLength options parameter

# Examples
 - https://admin:test@www.example.com/ --> www.example.com
 - http://www.example.com///cgi-bin///user.pl?user=admin/// --> www.example.com-cgi-bin-user.pl-user=admin
 - https://www.example.com?q=<>:\"/\\|?*A --> www.example.com-q=-A
 - https://www.example.com/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/extremely-long-url/ -->
  www.example.com-extremely-long-url-extremely-long-
# Installation 
'''go get github.com/slotix/slugifyurl'''