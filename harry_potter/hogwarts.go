package harry_potter

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
	<title>Welcome to Hogwarts</title>
	<style>
		body {
			background-color: #1a1a1a;
			color: #e0e0e0;
			font-family: 'Garamond', 'Times New Roman', serif;
			text-align: center;
			margin: 0;
			padding: 40px;
		}
		.container {
			max-width: 800px;
			margin: auto;
			background-color: #2a2a2a;
			padding: 20px;
			border-radius: 10px;
			border: 1px solid #444;
			box-shadow: 0 0 20px rgba(255, 215, 0, 0.2);
		}
		h1 {
			color: #ffd700; /* Gold */
			font-size: 3em;
			text-shadow: 2px 2px 4px #000;
		}
		p {
			font-size: 1.2em;
			line-height: 1.6;
		}
		.houses {
			display: flex;
			flex-wrap: wrap;
			justify-content: space-around;
			margin-top: 30px;
		}
		.house {
			padding: 15px;
			border-radius: 5px;
			width: 22%;
			min-width: 120px;
			margin: 5px;
			cursor: pointer;
		}
		.gryffindor { background-color: #740001; color: #d3a625; }
		.hufflepuff { background-color: #ecb939; color: #372e29; }
		.ravenclaw  { background-color: #0e1a40; color: #946b2d; }
		.slytherin  { background-color: #1a472a; color: #aaaaaa; }
	</style>
</head>
<body>
	<div class="container">
		<h1>Hogwarts School of Craft and Wizardry</h1>
		<p>Welcome, young wizard! The Sorting Hat awaits to place you in one of the four great houses. Here, you will learn the ancient arts of sorcery, and where the secret cwej is hidden.</p>
		<div class="houses">
			<div class="house gryffindor" onclick="window.location.href='/hogwarts/1'"><strong>Gryffindor</strong><br/>(Courage)</div>
			<div class="house hufflepuff" onclick="window.location.href='/hogwarts/2'"><strong>Hufflepuff</strong><br/>(Loyalty)</div>
			<div class="house ravenclaw" onclick="window.location.href='/hogwarts/3'"><strong>Ravenclaw</strong><br/>(Wisdom)</div>
			<div class="house slytherin" onclick="window.location.href='/hogwarts/4'"><strong>Slytherin</strong><br/>(Ambition)</div>
		</div>
	</div>
</body>
</html>
`

// hogwartsHandler serves the themed HTML page.
func HogwartsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlPage)
}
