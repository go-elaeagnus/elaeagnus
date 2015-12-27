package ela

var URIMapping map[string]interface{}

func init() {
	URIMapping = make(map[string]interface{})
}

// put uri-func mapping into map
func Router(uri string, f func(Context)) {
	URIMapping[uri] = f
}
