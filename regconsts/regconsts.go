package regconsts

// Taken from https://bit.ly/2X3EFcS
//
//goland:noinspection ALL
const (
	Alpha                                  = "^[a-zA-Z]+$"                                                                                                                                                                                                                             // Alpha represents regular expression for alpha characters
	AlphaDash                              = "^[a-zA-Z0-9_\\-]+$"                                                                                                                                                                                                                      // AlphaDash represents regular expression for alpha characters with underscore and dash
	AlphaSpace                             = "^[-a-zA-Z0-9_ ]+$"                                                                                                                                                                                                                       // AlphaSpace represents regular expression for alpha characters with underscore, space and dash
	AlphaNumeric                           = "^[a-zA-Z0-9]+$"                                                                                                                                                                                                                          // AlphaNumeric represents regular expression for alpha numeric characters
	CreditCard                             = "^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"                                                                              // CreditCard represents regular expression for credit cards like (Visa, MasterCard, American Express, Diners Club, Discover, and JCB cards). Ref: https://stackoverflow.com/questions/9315647/regex-credit-card-number-tests
	Coordinate                             = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?),\\s*[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"                                                                                                                                      // Ref: https://stackoverflow.com/questions/3518504/regular-expression-for-matching-latitude-longitude-coordinates Coordinate represents latitude and longitude regular expression
	CSSColor                               = "^(#([\\da-f]{3}){1,2}|(rgb|hsl)a\\((\\d{1,3}%?,\\s?){3}(1|0?\\.\\d+)\\)|(rgb|hsl)\\(\\d{1,3}%?(,\\s?\\d{1,3}%?){2}\\))$"                                                                                                                 // CSSColor represents css valid color code with hex, rgb, rgba, hsl, hsla etc. Ref: http://www.regexpal.com/97509
	Date                                   = "^(((19|20)([2468][048]|[13579][26]|0[48])|2000)[/-]02[/-]29|((19|20)[0-9]{2}[/-](0[469]|11)[/-](0[1-9]|[12][0-9]|30)|(19|20)[0-9]{2}[/-](0[13578]|1[02])[/-](0[1-9]|[12][0-9]|3[01])|(19|20)[0-9]{2}[/-]02[/-](0[1-9]|1[0-9]|2[0-8])))$" // Date represents regular expression for valid date like: yyyy-mm-dd
	DateDDMMYY                             = "^(0?[1-9]|[12][0-9]|3[01])[\\/\\-](0?[1-9]|1[012])[\\/\\-]\\d{4}$"                                                                                                                                                                       // DateDDMMYY represents regular expression for valid date of format dd/mm/yyyy , dd-mm-yyyy etc.Ref: http://regexr.com/346hf
	MacAddress                             = "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$"                                                                                                                                                                                               // MacAddress represents regular expression for mac address
	Numeric                                = "^[-+]?[0-9]+$"                                                                                                                                                                                                                           // Numeric represents regular expression for numeric
	IntegerOnly                            = "^[0-9]+"
	SignedOrInteger                        = "^[-+]?[0-9]+"
	Path                                   = `[/\\]{0,2}[_-]*[Aa0-zZ9. ]+[_-]*[Aa0-zZ9. ]*(?:\!\@\#\$\%\&)?[/\\]?|^/$` // https://t.ly/Q58e {includes windows, unix and space}
	IpSimple                               = `(\d{1,3}\.){3}\d{1,3}[^\r\n]*`
	FloatingPointNumberOnly                = `^[-+]?[0-9]+[\.][0-9]+`
	SignOrIntegerOrFloatingPointNumber     = `^[-+]?[0-9]+[\.]?[0-9]+`
	AnyIdentifier                          = `[a-zA-Z][a-zA-Z_\d\-]*`
	StringDoubleQuote                      = `"(?:\\.|[^"])*"`
	StringSingleQuote                      = `'((?:\\.|[^'])*'|"(?:\\.|[^"])*")`
	StringSingleOrDoubleQuote              = `'((?:\\.|[^'])*'|"(?:\\.|[^"])*")`
	Url                                    = "^(?:http(s)?:\\/\\/)?[\\w.-]+(?:\\.[\\w\\.-]+)+[\\w\\-\\._~:/?#[\\]@!\\$&'\\(\\)\\*\\+,;=.]+$" // URL represents regular expression for url Ref: https://stackoverflow.com/questions/136505/searching-for-uuids-in-text-with-regex
	UUID                                   = "^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}$"                       // UUID represents regular expression for UUID
	UUID3                                  = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"                               // UUID3 represents regular expression for UUID version 3
	UUID4                                  = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"                         // UUID4 represents regular expression for UUID version 4
	UUID5                                  = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"                         // UUID5 represents regular expression for UUID version 5
	ISBN10                                 = "^(?:[0-9]{9}X|[0-9]{10})$"
	ISBN13                                 = "^(?:[0-9]{13})$"
	Alphanumeric                           = "^[a-zA-Z0-9]+$"
	Int                                    = `^[-+]?(?:[1-9][0-9]*)$`
	Hexadecimal                            = "^[0-9a-fA-F]+$"
	Hexcolor                               = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	RGBcolor                               = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	ASCII                                  = "^[\x00-\x7F]+$"
	Multibyte                              = "[^\x00-\x7F]"
	FullWidth                              = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	HalfWidth                              = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	Base64                                 = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	PrintableASCII                         = "^[\x20-\x7E]+$"
	DataUri                                = "^data:.+\\/(.+);base64$"
	MagnetUri                              = "^magnet:\\?xt=urn:[a-zA-Z0-9]+:[a-zA-Z0-9]{32,40}&dn=.+&tr=.+$"
	DNSName                                = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`
	UrlSchema                              = `((ftp|tcp|udp|wss?|https?):\/\/)`
	UrlUsername                            = `(\S+(:\S*)?@)`
	UrlPath                                = `((\/|\?|#)[^\s]*)`
	UrlPort                                = `(:(\d{1,5}))`
	UrlIP                                  = `([1-9]\d?|1\d\d|2[01]\d|22[0-3]|24\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-5]))`
	UrlSubdomain                           = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	SSN                                    = `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
	WinPathWithOrWithoutUnc                = `(\\\\\?\\)?[a-zA-Z]:(\\|/)[_-]*[Aa0-zZ9. ]+[_-]*[Aa0-zZ9. ]*(?:\!\@\#\$\%\&)?(\\|/)?[_-]*[Aa0-zZ9. ]+[_-]*[Aa0-zZ9. ]*` // https://regex101.com/r/a5CLhu/1
	UnixPathWithoutSpace                   = `/[Aa0-zZ9.]+[_-]*[Aa0-zZ9.]*|^/$`                                                                                       // https://regex101.com/r/DqsKNb/1
	Semver                                 = "^v?(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"
	HasLowerCase                           = ".*[[:lower:]]"
	HasUpperCase                           = ".*[[:upper:]]"
	HasWhitespace                          = ".*[[:space:]]"
	HasWhitespaceOnly                      = "^[[:space:]]+$"
	IMEI                                   = "^[0-9a-f]{14}$|^\\d{15}$|^\\d{18}$"
	Identifier                             = AnyIdentifier
	UnderscoreIdentifier                   = `_*` + Identifier
	ThirdBracketHeader                     = `\[[a-zA-Z][a-zA-Z_\d\-]*\]`
	IMSI                                   = "^\\d{14,15}$"
	HashComment                            = "#[^\\n]*"        // Finds all text after hash symbol
	SlashComment                           = `\\[^\n]*`        // Finds all text after slash symbol
	HashSemicolonComment                   = "[#;][^\\n]*"     // Finds all text after hash or semicolon symbol
	HashCommentWithSpaceOptional           = "\\s*#[^\\n]*"    // Finds anything starts with or ends with hash `#`, can have optional spaces before
	HashOrSemicolonComment                 = "\\s*[#;][^\\n]*" // Finds anything starts with or ends with hash `#` or semi-colon `;`, can have optional spaces before
	AllWhitespaces                         = "\\s+"
	AllWhitespacesOrPipe                   = "\\s+|\\|+"
	EachWordsWithDollarSymbolDefinition    = `(\$\{` + AnyIdentifier + `\}|\$` + AnyIdentifier + `)+` // Selects a full word as "$identifier" symbol or ${identifier}
	EachWordsWithinPercentSymbolDefinition = `(\%\{` + AnyIdentifier + `\}|\%` + AnyIdentifier + `)+` // Selects a full word as "%identifier" symbol or %{identifier}
	PrettyName                             = `^PRETTY_NAME=(.*)$`                                     // https://t.ly/VsNo
	ExactIdFieldMatching                   = `^ID=(.*)$`                                              // https://t.ly/VsNo
	ExactVersionIdFieldMatching            = `^VERSION_ID=(.*)$`                                      // https://t.ly/VsNo
	UbuntuNameChecker                      = `[\( ]([\d\.]+)`                                         // https://t.ly/VsNo
	CentOsNameChecker                      = `^CentOS( Linux)? release ([\d\.]+) `                    // https://t.ly/VsNo
	RedHatNameChecker                      = `[\( ]([\d\.]+)`                                         // https://t.ly/VsNo
	FirstNumberAnyWhere                    = `(\d+){1}`                                               // https://regex101.com/r/7euGv5/1
	WindowsVersionNumberChecker            = FirstNumberAnyWhere                                      // https://regex101.com/r/7euGv5/1
	IpEthernet                             = `(\d{1,3}\.){3}\d{1,3}`
	IpWithSubnet                           = `(\d{1,3}\.){3}\d{1,3}\/\d{1,2}` // Ref: https://regexr.com/5om44
	Ipv6WithSubnet                         = `(([0-9a-fA-F]{0,4}:){1,7}[0-9a-fA-F]{0,4})`
	IpWithPort                             = `(\d{1,3}\.){3}\d{1,3}:\d{1,6}`
	ValidMac                               = `^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$` // 00-0C-29-EA-31-5B
	SemiColonSeparatedEqualKeyValSplitter  = `^(.+=(.*);)+$`                             // ex : `email=abc@abc.com; phone=1223;` => []KeyVals => {email=>abc@abc.com}...
	CommaSeparatedValue                    = `^(.+)(,\s*.+)+$`                           // ex : name,number.....
)
