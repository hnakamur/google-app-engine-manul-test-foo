package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

const resp = `<!DOCTYPE html>
<html>
<head>
<meta name="go-import" content="manul-test-foo.appspot.com/a/b git https://github.com/kovetskiy/manul-test-foo">

<meta name="go-source" content="manul-test-foo.appspot.com/a/b https://github.com/kovetskiy/manul-test-foo https://github.com/kovetskiy/manul-test-foo/tree/master{/dir} https://github.com/kovetskiy/manul-test-foo/tree/master{/dir}/{file}#L{line}">

<meta http-equiv="refresh" content="0; url=https://github.com/kovetskiy/manul-test-foo">
</head>
<body>
Nothing to see here. Please <a href="https://github.com/kovetskiy/manul-test-foo">move along</a>.
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, resp)
}
