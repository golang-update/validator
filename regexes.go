package validator

import "regexp"

const (
	alphaRegexString                 = "^[a-zA-Z]+$"
	alphaNumericRegexString          = "^[a-zA-Z0-9]+$"
	alphaUnicodeRegexString          = "^[\\p{L}]+$"
	alphaUnicodeNumericRegexString   = "^[\\p{L}\\p{N}]+$"
	numericRegexString               = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	numberRegexString                = "^[0-9]+$"
	hexadecimalRegexString           = "^(0[xX])?[0-9a-fA-F]+$"
	hexColorRegexString              = "^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{4}|[0-9a-fA-F]{6}|[0-9a-fA-F]{8})$"
	rgbRegexString                   = "^rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)$"
	rgbaRegexString                  = "^rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	hslRegexString                   = "^hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)$"
	hslaRegexString                  = "^hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)$"
	emailRegexString                 = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	e164RegexString                  = "^\\+[1-9]?[0-9]{7,14}$"
	base64RegexString                = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	base64URLRegexString             = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$"
	base64RawURLRegexString          = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2,4})$"
	iSBN10RegexString                = "^(?:[0-9]{9}X|[0-9]{10})$"
	iSBN13RegexString                = "^(?:(?:97(?:8|9))[0-9]{10})$"
	iSSNRegexString                  = "^(?:[0-9]{4}-[0-9]{3}[0-9X])$"
	uUID3RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uUID4RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUID5RegexString                 = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUIDRegexString                  = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	uUID3RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	uUID4RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUID5RFC4122RegexString          = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	uUIDRFC4122RegexString           = "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	uLIDRegexString                  = "^[A-HJKMNP-TV-Z0-9]{26}$"
	md4RegexString                   = "^[0-9a-f]{32}$"
	md5RegexString                   = "^[0-9a-f]{32}$"
	sha256RegexString                = "^[0-9a-f]{64}$"
	sha384RegexString                = "^[0-9a-f]{96}$"
	sha512RegexString                = "^[0-9a-f]{128}$"
	ripemd128RegexString             = "^[0-9a-f]{32}$"
	ripemd160RegexString             = "^[0-9a-f]{40}$"
	tiger128RegexString              = "^[0-9a-f]{32}$"
	tiger160RegexString              = "^[0-9a-f]{40}$"
	tiger192RegexString              = "^[0-9a-f]{48}$"
	aSCIIRegexString                 = "^[\x00-\x7F]*$"
	printableASCIIRegexString        = "^[\x20-\x7E]*$"
	multibyteRegexString             = "[^\x00-\x7F]"
	dataURIRegexString               = `^data:((?:\w+\/(?:([^;]|;[^;]).)+)?)`
	latitudeRegexString              = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	longitudeRegexString             = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	sSNRegexString                   = `^[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])$`
	hostnameRegexStringRFC952        = `^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`                                                                   // https://tools.ietf.org/html/rfc952
	hostnameRegexStringRFC1123       = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62}){1}(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?$`                                 // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	fqdnRegexStringRFC1123           = `^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$` // same as hostnameRegexStringRFC1123 but must contain a non numerical TLD (possibly ending with '.')
	btcAddressRegexString            = `^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$`                                                                             // bitcoin address
	btcAddressUpperRegexStringBech32 = `^BC1[02-9AC-HJ-NP-Z]{7,76}$`                                                                                   // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	btcAddressLowerRegexStringBech32 = `^bc1[02-9ac-hj-np-z]{7,76}$`                                                                                   // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	ethAddressRegexString            = `^0x[0-9a-fA-F]{40}$`
	ethAddressUpperRegexString       = `^0x[0-9A-F]{40}$`
	ethAddressLowerRegexString       = `^0x[0-9a-f]{40}$`
	uRLEncodedRegexString            = `^(?:[^%]|%[0-9A-Fa-f]{2})*$`
	hTMLEncodedRegexString           = `&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`
	hTMLRegexString                  = `<[/]?([a-zA-Z]+).*?>`
	jWTRegexString                   = "^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]*$"
	splitParamsRegexString           = `'[^']*'|\S+`
	bicRegexString                   = `^[A-Za-z]{6}[A-Za-z0-9]{2}([A-Za-z0-9]{3})?$`
	semverRegexString                = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$` // numbered capture groups https://semver.org/
	dnsRegexStringRFC1035Label       = "^[a-z]([-a-z0-9]*[a-z0-9]){0,62}$"
	cveRegexString                   = `^CVE-(1999|2\d{3})-(0[^0]\d{2}|0\d[^0]\d{1}|0\d{2}[^0]|[1-9]{1}\d{3,})$` // CVE Format Id https://cve.mitre.org/cve/identifiers/syntaxchange.html
	mongodbIdRegexString             = "^[a-f\\d]{24}$"
	mongodbConnStringRegexString     = "^mongodb(\\+srv)?:\\/\\/(([a-zA-Z\\d]+):([a-zA-Z\\d$:\\/?#\\[\\]@]+)@)?(([a-z\\d.-]+)(:[\\d]+)?)((,(([a-z\\d.-]+)(:(\\d+))?))*)?(\\/[a-zA-Z-_]{1,64})?(\\?(([a-zA-Z]+)=([a-zA-Z\\d]+))(&(([a-zA-Z\\d]+)=([a-zA-Z\\d]+))?)*)?$"
	cronRegexString                  = `(@(annually|yearly|monthly|weekly|daily|hourly|reboot))|(@every (\d+(ns|us|µs|ms|s|m|h))+)|((((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7})`
	spicedbIDRegexString             = `^(([a-zA-Z0-9/_|\-=+]{1,})|\*)$`
	spicedbPermissionRegexString     = "^([a-z][a-z0-9_]{1,62}[a-z0-9])?$"
	spicedbTypeRegexString           = "^([a-z][a-z0-9_]{1,61}[a-z0-9]/)?[a-z][a-z0-9_]{1,62}[a-z0-9]$"
)

var (
	alphaRegex                 = regexp.MustCompile(alphaRegexString)
	alphaNumericRegex          = regexp.MustCompile(alphaNumericRegexString)
	alphaUnicodeRegex          = regexp.MustCompile(alphaUnicodeRegexString)
	alphaUnicodeNumericRegex   = regexp.MustCompile(alphaUnicodeNumericRegexString)
	numericRegex               = regexp.MustCompile(numericRegexString)
	numberRegex                = regexp.MustCompile(numberRegexString)
	hexadecimalRegex           = regexp.MustCompile(hexadecimalRegexString)
	hexColorRegex              = regexp.MustCompile(hexColorRegexString)
	rgbRegex                   = regexp.MustCompile(rgbRegexString)
	rgbaRegex                  = regexp.MustCompile(rgbaRegexString)
	hslRegex                   = regexp.MustCompile(hslRegexString)
	hslaRegex                  = regexp.MustCompile(hslaRegexString)
	e164Regex                  = regexp.MustCompile(e164RegexString)
	emailRegex                 = regexp.MustCompile(emailRegexString)
	base64Regex                = regexp.MustCompile(base64RegexString)
	base64URLRegex             = regexp.MustCompile(base64URLRegexString)
	base64RawURLRegex          = regexp.MustCompile(base64RawURLRegexString)
	iSBN10Regex                = regexp.MustCompile(iSBN10RegexString)
	iSBN13Regex                = regexp.MustCompile(iSBN13RegexString)
	iSSNRegex                  = regexp.MustCompile(iSSNRegexString)
	uUID3Regex                 = regexp.MustCompile(uUID3RegexString)
	uUID4Regex                 = regexp.MustCompile(uUID4RegexString)
	uUID5Regex                 = regexp.MustCompile(uUID5RegexString)
	uUIDRegex                  = regexp.MustCompile(uUIDRegexString)
	uUID3RFC4122Regex          = regexp.MustCompile(uUID3RFC4122RegexString)
	uUID4RFC4122Regex          = regexp.MustCompile(uUID4RFC4122RegexString)
	uUID5RFC4122Regex          = regexp.MustCompile(uUID5RFC4122RegexString)
	uUIDRFC4122Regex           = regexp.MustCompile(uUIDRFC4122RegexString)
	uLIDRegex                  = regexp.MustCompile(uLIDRegexString)
	md4Regex                   = regexp.MustCompile(md4RegexString)
	md5Regex                   = regexp.MustCompile(md5RegexString)
	sha256Regex                = regexp.MustCompile(sha256RegexString)
	sha384Regex                = regexp.MustCompile(sha384RegexString)
	sha512Regex                = regexp.MustCompile(sha512RegexString)
	ripemd128Regex             = regexp.MustCompile(ripemd128RegexString)
	ripemd160Regex             = regexp.MustCompile(ripemd160RegexString)
	tiger128Regex              = regexp.MustCompile(tiger128RegexString)
	tiger160Regex              = regexp.MustCompile(tiger160RegexString)
	tiger192Regex              = regexp.MustCompile(tiger192RegexString)
	aSCIIRegex                 = regexp.MustCompile(aSCIIRegexString)
	printableASCIIRegex        = regexp.MustCompile(printableASCIIRegexString)
	multibyteRegex             = regexp.MustCompile(multibyteRegexString)
	dataURIRegex               = regexp.MustCompile(dataURIRegexString)
	latitudeRegex              = regexp.MustCompile(latitudeRegexString)
	longitudeRegex             = regexp.MustCompile(longitudeRegexString)
	sSNRegex                   = regexp.MustCompile(sSNRegexString)
	hostnameRegexRFC952        = regexp.MustCompile(hostnameRegexStringRFC952)
	hostnameRegexRFC1123       = regexp.MustCompile(hostnameRegexStringRFC1123)
	fqdnRegexRFC1123           = regexp.MustCompile(fqdnRegexStringRFC1123)
	btcAddressRegex            = regexp.MustCompile(btcAddressRegexString)
	btcUpperAddressRegexBech32 = regexp.MustCompile(btcAddressUpperRegexStringBech32)
	btcLowerAddressRegexBech32 = regexp.MustCompile(btcAddressLowerRegexStringBech32)
	ethAddressRegex            = regexp.MustCompile(ethAddressRegexString)
	uRLEncodedRegex            = regexp.MustCompile(uRLEncodedRegexString)
	hTMLEncodedRegex           = regexp.MustCompile(hTMLEncodedRegexString)
	hTMLRegex                  = regexp.MustCompile(hTMLRegexString)
	jWTRegex                   = regexp.MustCompile(jWTRegexString)
	splitParamsRegex           = regexp.MustCompile(splitParamsRegexString)
	bicRegex                   = regexp.MustCompile(bicRegexString)
	semverRegex                = regexp.MustCompile(semverRegexString)
	dnsRegexRFC1035Label       = regexp.MustCompile(dnsRegexStringRFC1035Label)
	cveRegex                   = regexp.MustCompile(cveRegexString)
	mongodbIdRegex             = regexp.MustCompile(mongodbIdRegexString)
	mongodbConnectionRegex     = regexp.MustCompile(mongodbConnStringRegexString)
	cronRegex                  = regexp.MustCompile(cronRegexString)
	spicedbIDRegex             = regexp.MustCompile(spicedbIDRegexString)
	spicedbPermissionRegex     = regexp.MustCompile(spicedbPermissionRegexString)
	spicedbTypeRegex           = regexp.MustCompile(spicedbTypeRegexString)
)
