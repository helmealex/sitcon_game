package lord_of_the_rings

import (
	"fmt"
	"net/http"
)

const htmlPage = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Welcome to Middle-earth</title>
	<style>
		@import url('https://fonts.googleapis.com/css2?family=Cinzel:wght@700&family=Lora:ital@0;1&display=swap');

		body {
			font-family: 'Lora', serif;
			background-color: #f0e6d2; /* Parchment paper */
			color: #4a2c2a; /* Dark brown */
			text-align: center;
			margin: 0;
			padding: 2rem;
		}

		.container {
			max-width: 800px;
			margin: 0 auto;
		}

		h1 {
			font-family: 'Cinzel', serif;
			font-size: 3.5rem;
			color: #3e2723; /* Darker brown */
			text-shadow: 2px 2px 4px #aaa;
		}

		p {
			font-size: 1.2rem;
			line-height: 1.6;
		}

		.quote {
			font-style: italic;
			font-size: 1.4rem;
			margin: 2rem 0;
			padding: 1rem;
			border-left: 4px solid #8d6e63; /* Muted brown border */
			background-color: #e6d9c1; /* Slightly darker parchment */
		}
	</style>
</head>
<body>
	<div class="container">
		<h1>Lord of the Rings</h1>
		<p>Welcome, traveler, to a land of myth and legend.</p>
		
		<div class="quote">
			<p>"One Ring to rule them all, One Ring to find them,<br>
			One Ring to bring them all and in the darkness bind them."</p>
		</div>

		<p>
			It's a dangerous business, Frodo, going out your door. You step onto the road, and if you don't keep your feet, there's no knowing where you might be swept off to.
		</p>
	</div>
</body>
</html>
`

// LotrHandler serves a Lord of the Rings themed HTML page.
func LotrHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlPage)
}
