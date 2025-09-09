package harry_potter

import (
	"fmt"
	"net/http"
)

// RavenclawHandler serves a Ravenclaw-themed HTML page.
func RavenclawHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	htmlContent := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>ne</title>
	<style>
		body {
			background-color: #0e1a40; /* Ravenclaw Blue */
			color: #946b2d; /* Ravenclaw Bronze */
			font-family: 'Times New Roman', Times, serif;
			display: flex;
			justify-content: center;
			align-items: center;
			height: 100vh;
			margin: 0;
		}
		h1 {
			font-size: 15vw;
			font-weight: bold;
		}
	</style>
</head>
<body>
	<h1>3</h1>
	<h1>R</h1>
</body>
</html>`
	fmt.Fprint(w, htmlContent)
}
