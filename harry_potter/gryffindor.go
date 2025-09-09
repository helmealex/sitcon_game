package harry_potter

import (
	"fmt"
	"net/http"
)

// gryffindorHandler serves an HTML page with a Gryffindor theme.
func GryffindorHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>vi</title>
	<style>
		body {
			background-color: #740001; /* Scarlet */
			color: #D3A625;           /* Gold */
			display: flex;
			justify-content: center;
			align-items: center;
			height: 100vh;
			margin: 0;
			font-family: 'Times New Roman', Times, serif;
			font-size: 20rem;
		}
	</style>
</head>
<body>
	<h1>1</h1>
	<h1>L</h1>
</body>
</html>
`
	fmt.Fprint(w, html)
}
