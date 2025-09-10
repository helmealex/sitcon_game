package matrix

import (
	"fmt"
	"net/http"
)

const htmlContent = `
<!DOCTYPE html>
<html>
<head>
<title>Matrix</title>
<style>
	body {
		margin: 0;
		overflow: hidden;
		background-color: black;
	}
	canvas {
		display: block;
	}
</style>
</head>
<body>
<canvas id="matrix-canvas"></canvas>
<script>
	const canvas = document.getElementById('matrix-canvas');
	const ctx = canvas.getContext('2d');

	canvas.width = window.innerWidth;
	canvas.height = window.innerHeight;

	const katakana = 'アァカサタナハマヤャラワガザダバパイィキシチニヒミリヰギジヂビピウゥクスツヌフムユュルグズブヅプエェケセテネヘメレヱゲゼデベペオォコソトノホモヨョロヲゴゾドボポヴッン';
	const latin = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
	const nums = '0123456789';
	const alphabet = katakana + latin + nums;

	const fontSize = 16;
	const columns = Math.floor(canvas.width / fontSize);

	const rainDrops = [];
	for (let x = 0; x < columns; x++) {
		rainDrops[x] = 1;
	}

	const draw = () => {
		ctx.fillStyle = 'rgba(0, 0, 0, 0.05)';
		ctx.fillRect(0, 0, canvas.width, canvas.height);

		ctx.fillStyle = '#0F0'; // Green text
		ctx.font = fontSize + 'px monospace';

		for (let i = 0; i < rainDrops.length; i++) {
			const text = alphabet.charAt(Math.floor(Math.random() * alphabet.length));
			ctx.fillText(text, i * fontSize, rainDrops[i] * fontSize);

			if (rainDrops[i] * fontSize > canvas.height && Math.random() > 0.975) {
				rainDrops[i] = 0;
			}
			rainDrops[i]++;
		}
	};

	setInterval(draw, 33);

	window.addEventListener('resize', () => {
		canvas.width = window.innerWidth;
		canvas.height = window.innerHeight;
		// Recalculate columns and reset drops if needed
	});
</script>

<img src="http://localhost:8080/matrix/flag.png" display="none" alt="hidden flag">
</body>
</html>
`

// MatrixHandler serves the Matrix-themed HTML page.
func MatrixHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, htmlContent)
}
