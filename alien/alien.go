package alien

import (
	"fmt"
	"net/http"
)

// AlienHandler serves an Alien vs. Predator-themed HTML page.
func AlienHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	htmlContent := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Alien vs. Predator</title>
	<style>
		body {
			background-color: #000000;
			color: #cccccc;
			font-family: 'Arial Black', Gadget, sans-serif;
			text-align: center;
			padding: 50px;
			overflow: hidden;
		}
		h1 {
			color: #b22222; /* Firebrick red */
			text-shadow: 0 0 15px #ff0000;
			font-size: 2.5em;
			letter-spacing: 3px;
		}
		.battleground {
			display: flex;
			justify-content: center;
			align-items: center;
			gap: 40px;
			margin-top: 50px;
		}
		.combatant {
			font-size: 100px;
			line-height: 1;
			animation: float 4s ease-in-out infinite;
		}
		.predator {
			animation-delay: -2s;
		}
		.vs {
			font-size: 50px;
			color: #ffc107;
			animation: pulse 2s ease-in-out infinite;
		}
		p {
			margin-top: 50px;
			font-style: italic;
			color: #00ff00;
		}
		@keyframes float {
			0% { transform: translateY(0px); }
			50% { transform: translateY(-25px); }
			100% { transform: translateY(0px); }
		}
		@keyframes pulse {
			0% { transform: scale(1); opacity: 0.7; }
			50% { transform: scale(1.2); opacity: 1; }
			100% { transform: scale(1); opacity: 0.7; }
		}
	</style>
</head>
<body>
	<h1>WHOEVER WINS... WE LOSE</h1>
	<div class="battleground">
		<div class="combatant alien">👽</div>
		<div class="vs">VS</div>
		<div class="combatant predator">💀</div>
	</div>
	<p>The hunt has begun.</p>
</body>
</html>
`
	fmt.Fprint(w, htmlContent)
}
