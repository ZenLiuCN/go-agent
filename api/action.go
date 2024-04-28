package api

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/ZenLiuCN/fn"
	"github.com/ZenLiuCN/go-agent/model"
	"github.com/ZenLiuCN/go-agent/units"
	"github.com/dop251/goja"
	"io"
	"net/http"
)

type Action struct {
	User  *model.User
	Cache *units.Cached[string, string]
	http.ResponseWriter
	*http.Request
	read bool
	dump *bytes.Buffer
}

func (a *Action) SearchQuery(key string) string {
	return a.URL.Query().Get(key)
}
func (a *Action) JsonBody() any {
	if a.read {
		return goja.Undefined()
	}
	if a.Body == nil {
		return goja.Undefined()
	}
	defer fn.IgnoreClose(a.Body)
	defer func() { a.read = true }()
	b := fn.Panic1(io.ReadAll(a.Body))
	if b[0] == '{' {
		var v map[string]any
		fn.Panic(json.Unmarshal(b, &v))
		return v
	} else if b[0] == '[' {
		var v []any
		fn.Panic(json.Unmarshal(b, &v))
		return v
	} else {
		var v any
		fn.Panic(json.Unmarshal(b, &v))
		return v
	}
}
func (a *Action) BinaryBody() any {
	if a.read {
		return goja.Undefined()
	}
	if a.Body == nil {
		return goja.Undefined()
	}
	defer fn.IgnoreClose(a.Body)
	defer func() { a.read = true }()
	return fn.Panic1(io.ReadAll(a.Body))
}
func (a *Action) TextBody() any {
	if a.read {
		return goja.Undefined()
	}
	if a.Body == nil {
		return goja.Undefined()
	}
	defer fn.IgnoreClose(a.Body)
	return string(fn.Panic1(io.ReadAll(a.Body)))
}
func (a *Action) Ok() {
	a.WriteHeader(http.StatusOK)
}
func (a *Action) Status(code int) {
	a.WriteHeader(code)
}
func (a *Action) Json(value goja.Value) {
	a.ResponseWriter.Header().Set("Content-Type", "application/json;charset=utf-8")
	switch v := value.Export().(type) {
	case []byte:
		_, _ = a.ResponseWriter.Write(v)
	case string:
		_, _ = a.ResponseWriter.Write([]byte(v))
	case goja.ArrayBuffer:
		_, _ = a.ResponseWriter.Write(v.Bytes())
	case map[string]any, []any, []map[string]any:
		b := fn.Panic1(json.Marshal(v))
		_, _ = a.ResponseWriter.Write(b)
	case nil:
		_, _ = a.ResponseWriter.Write([]byte("null"))
	default:
		b := fn.Panic1(json.Marshal(v))
		_, _ = a.ResponseWriter.Write(b)
	}
}
func (a *Action) Redirect(url string, status int) {
	http.Redirect(a.ResponseWriter, a.Request, url, status)
}
func (a *Action) Binary(contentType string, value goja.Value) {
	a.ResponseWriter.Header().Set("Content-Type", contentType)
	switch v := value.Export().(type) {
	case []byte:
		_, _ = a.ResponseWriter.Write(v)
	case goja.ArrayBuffer:
		_, _ = a.ResponseWriter.Write(v.Bytes())
	default:
		panic(fmt.Errorf("bad type: %T", value))
	}
}

func (a *Action) Text(text string) {
	a.ResponseWriter.Header().Set("Content-Type", "text/plain;charset=utf-8")
	_, _ = a.ResponseWriter.Write([]byte(text))
}
func (a *Action) Dump() string {
	if a.dump == nil || a.dump.Len() == 0 {
		if a.dump == nil {
			a.dump = new(bytes.Buffer)
		}
		g := gabs.New()
		req := gabs.New()
		_, _ = req.Set(a.Request.Header, "header")
		_, _ = req.Set(a.Request.RequestURI, "uri")
		_, _ = req.Set(a.Request.Method, "method")
		_, _ = req.Set(a.Request.RemoteAddr, "remoteAddr")
		_, _ = req.Set(a.Request.Referer(), "referer")
		_, _ = req.Set(a.Request.URL.Query(), "query")
		if !a.read && a.Request.Body != nil {
			var snap = make([]byte, 64)
			n, _ := a.Request.Body.Read(snap)
			a.read = true
			_, _ = req.Set(hex.EncodeToString(snap[:n]), "body")
		}
		_, _ = g.Set(a.User, "user")
		_, _ = g.Set(req, "request")
		_, _ = g.Set(a.Request.Header, "response", "header")
		a.dump.WriteString(g.String())
	}
	return a.dump.String()
}
