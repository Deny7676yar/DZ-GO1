package configreader

import "net/url"



func IsUrl(filename string) bool {

	u, err := url.ParseRequestURI(filename)
	return err == nil && u.Scheme != "" && u.Host != ""
}
