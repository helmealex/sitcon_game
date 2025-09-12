package end

import (
	"fmt"
	"net/http"
)

// EndHandler serves the final congratulations page for the adventurer.
func EndHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>The End</title>
	<style>
		@import url('https://fonts.googleapis.com/css2?family=MedievalSharp&display=swap');

		body {
			background-color: #111;
			color: #eee;
			font-family: 'MedievalSharp', cursive;
			display: flex;
			justify-content: center;
			align-items: center;
			height: 100vh;
			margin: 0;
			text-align: center;
		}

		.content {
			max-width: 600px;
			padding: 40px;
			border: 2px solid #888;
			border-radius: 10px;
			background: rgba(0, 0, 0, 0.3);
		}

		h1 {
			color: #ffd700; /* Gold color */
			font-size: 3rem;
			margin-bottom: 1rem;
		}

		p {
			font-size: 1.5rem;
			line-height: 1.6;
		}
	</style>
</head>
<body>
	<div class="content">
		<h1>Congratulations, Adventurer!</h1>
		<p>Your epic quest has reached its glorious conclusion.</p>
		<p>The realm is safe, thanks to your courage and might.</p>
	</div>
</body>
</html>
`
	fmt.Fprint(w, html)
}
