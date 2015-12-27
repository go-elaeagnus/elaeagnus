package ela

import (
	"fmt"
	"github.com/elago/ela/debug"
	"github.com/gogather/com/log"
	"net/http"
)

func Http(port int) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), &ElaHandler{})
	if err != nil {
		fmt.Printf("HTTP Server Start Failed Port [%d]\n%s", port, err)
	} else {
		log.Blueln("port [%d]", port)
	}
}

// Http handler
type ElaHandler struct{}

func (*ElaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// package ctx
	ctx := Context{}
	ctx.w = w
	ctx.r = r
	ctx.Data = make(map[string]interface{})
	// parse router and call action
	path := r.URL.String()
	debug.Log(path)
	f := URIMapping[path]
	if f != nil {
		function := f.(func(Context))
		function(ctx)
	} else {
		// show uri
		if StaticExist(path) {
			StaticServ(path, w)
		} else {
			http.Error(w, "404, File Not Exist", 404)
		}
	}
}
