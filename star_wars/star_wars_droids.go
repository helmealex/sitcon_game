package star_wars

import (
	"fmt"
	"net/http"
)

func DroidsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)

	html := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Not Found</title>
	<style>
		@import url('https://fonts.googleapis.com/css?family=News+Cycle:700');
		body {
			background-color: #000;
			color: #FFE81F; /* Star Wars Yellow */
			font-family: 'News Cycle', 'Arial Narrow Bold', sans-serif;
			display: flex;
			justify-content: center;
			align-items: center;
			height: 100vh;
			margin: 0;
			overflow: hidden;
			perspective: 400px;
		}
		.stars {
			position: absolute;
			top: 0;
			left: 0;
			right: 0;
			bottom: 0;
			width: 100%;
			height: 100%;
			display: block;
			background: transparent;
		}
		.stars:before, .stars:after {
			content: "";
			position: absolute;
			top: 0;
			left: 0;
			width: 1px;
			height: 1px;
			background: white;
			box-shadow: 
				50vw 80vh #FFF, 10vw 20vh #FFF, 85vw 30vh #FFF, 60vw 70vh #FFF,
				30vw 90vh #FFF, 90vw 10vh #FFF, 40vw 50vh #FFF, 70vw 60vh #FFF,
				20vw 40vh #FFF, 80vw 5vh #FFF, 5vw 95vh #FFF, 95vw 85vh #FFF,
				25vw 75vh #FFF, 75vw 25vh #FFF, 15vw 65vh #FFF, 65vw 15vh #FFF,
				12vw 58vh #FFF, 88vw 42vh #FFF, 37vw 19vh #FFF, 63vw 81vh #FFF,
				49vw 3vh #FFF, 51vw 97vh #FFF, 2vw 23vh #FFF, 98vw 77vh #FFF;
			border-radius: 50%;
		}
		.stars:after {
			box-shadow: 
				5vw 15vh #FFF, 95vw 25vh #FFF, 45vw 55vh #FFF, 55vw 45vh #FFF,
				35vw 85vh #FFF, 85vw 35vh #FFF, 65vw 95vh #FFF, 95vw 65vh #FFF,
				22vw 33vh #FFF, 33vw 22vh #FFF, 77vw 88vh #FFF, 88vw 77vh #FFF,
				18vw 48vh #FFF, 48vw 18vh #FFF, 82vw 62vh #FFF, 62vw 82vh #FFF,
				3vw 8vh #FFF, 97vw 92vh #FFF, 28vw 68vh #FFF, 72vw 32vh #FFF,
				58vw 28vh #FFF, 42vw 72vh #FFF, 13vw 43vh #FFF, 87vw 57vh #FFF;
		}
		.content {
			text-align: center;
			font-size: 2.5em;
			font-weight: bold;
			transform: rotateX(20deg);
			max-width: 80%;
			position: relative;
			z-index: 1;
		}
	</style>
</head>
<body>
	<div class="stars"></div>
	<div class="content">
		<p>the droids you are looking for are not here</p>
	</div>
</body>
</html>
`
	fmt.Fprint(w, html)
}
