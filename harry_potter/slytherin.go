package harry_potter

import (
	"fmt"
	"net/http"
)

// SlytherinHandler serves an HTML page with a Slytherin theme.
func SlytherinHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>re</title>
	<style>
		body {
			background-color: #1A472A; /* Slytherin Green */
			color: #AAAAAA;           /* Slytherin Silver */
			font-family: serif;
			text-align: center;
			padding-top: 50px;
		}
		h1 {
			font-size: 4em;
			text-shadow: 2px 2px 4px #000000;
		}
	</style>
</head>
<body>
	<h1>4</h1>
	<h1>D</h1>
</body>
</html>
`
	fmt.Fprint(w, html)
}
