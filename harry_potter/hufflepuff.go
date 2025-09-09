package harry_potter

import (
	"fmt"
	"net/http"
)

// HufflepuffHandler serves an HTML page with a Hufflepuff theme.
func HufflepuffHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>ge</title>
	<style>
		body {
			background-color: #000000; /* Black */
			color: #FFD700;           /* Gold/Yellow */
			font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
			display: flex;
			flex-direction: column;
			justify-content: center;
			align-items: center;
			height: 100vh;
			margin: 0;
		}
		h1 {
			font-size: 15vw;
			margin: 0;
			line-height: 1;
		}
	</style>
</head>
<body>
	<h1>2</h1>
	<h1>O</h1>
</body>
</html>`
	fmt.Fprint(w, html)
}
