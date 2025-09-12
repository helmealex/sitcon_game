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
<script>
function _0x138a(_0x3d4b56,_0x4855ee){const _0xa53ad9=_0xa53a();return _0x138a=function(_0x138acb,_0x2ccc54){_0x138acb=_0x138acb-0x10b;let _0x22606e=_0xa53ad9[_0x138acb];return _0x22606e;},_0x138a(_0x3d4b56,_0x4855ee);}const _0x1c218e=_0x138a;(function(_0x385e7f,_0x4dcc3e){const _0x48bcdf=_0x138a,_0x1d6e4a=_0x385e7f();while(!![]){try{const _0x57336a=-parseInt(_0x48bcdf(0x116))/0x1*(-parseInt(_0x48bcdf(0x110))/0x2)+-parseInt(_0x48bcdf(0x113))/0x3*(-parseInt(_0x48bcdf(0x11c))/0x4)+parseInt(_0x48bcdf(0x120))/0x5*(parseInt(_0x48bcdf(0x10f))/0x6)+-parseInt(_0x48bcdf(0x111))/0x7+-parseInt(_0x48bcdf(0x112))/0x8+parseInt(_0x48bcdf(0x10c))/0x9*(parseInt(_0x48bcdf(0x12f))/0xa)+-parseInt(_0x48bcdf(0x11e))/0xb*(-parseInt(_0x48bcdf(0x12d))/0xc);if(_0x57336a===_0x4dcc3e)break;else _0x1d6e4a['push'](_0x1d6e4a['shift']());}catch(_0x3e0add){_0x1d6e4a['push'](_0x1d6e4a['shift']());}}}(_0xa53a,0x832f8));const _m={'A':'.-','B':'-...','C':_0x1c218e(0x11d),'D':_0x1c218e(0x123),'E':'.','F':'..-.','G':_0x1c218e(0x10d),'H':_0x1c218e(0x131),'I':'..','J':_0x1c218e(0x11f),'K':'-.-','L':_0x1c218e(0x126),'M':'--','N':'-.','O':'---','P':'.--.','Q':'--.-','R':'.-.','S':_0x1c218e(0x118),'T':'-','U':_0x1c218e(0x121),'V':'...-','W':_0x1c218e(0x132),'X':_0x1c218e(0x127),'Y':_0x1c218e(0x117),'Z':_0x1c218e(0x115),'1':_0x1c218e(0x12a),'2':_0x1c218e(0x114),'3':_0x1c218e(0x125),'4':_0x1c218e(0x11a),'5':_0x1c218e(0x119),'6':_0x1c218e(0x12c),'7':'--...','8':_0x1c218e(0x130),'9':_0x1c218e(0x11b),'0':'-----','\x20':'/'},_d=0xc8,_a=_d*0x3,_s=_d,_l=_d*0x3,_w=_d*0x7;function _p(_0x32be98){return new Promise(_0x22c1a8=>setTimeout(_0x22c1a8,_0x32be98));}async function _f(_0x45cd02){const _0x380f58=_0x1c218e,_0x17ac7c=document['getElementById'](_0x380f58(0x12e));_0x17ac7c['style']['display']=_0x380f58(0x12b),await _p(_0x45cd02),_0x17ac7c[_0x380f58(0x122)][_0x380f58(0x129)]=_0x380f58(0x128);}function _0xa53a(){const _0xa3a3c8=['block','-....','756wDSPgk','siren','150kolaIO','---..','....','.--','toUpperCase','112113nsdMtx','--.','onload','197454whUqVO','2jVkeHa','1303715hExcaQ','3667968yWEqMw','17583qlVAlI','..---','--..','10874BnpQcd','-.--','...','.....','....-','----.','80sTNpWE','-.-.','42229CkozWJ','.---','95QrrpMf','..-','style','-..','END','...--','.-..','-..-','none','display','.----'];_0xa53a=function(){return _0xa3a3c8;};return _0xa53a();}async function _c(_0x41b004){const _0x19ae4c=_0x1c218e,_0x3d85ec=_0x41b004[_0x19ae4c(0x10b)]();for(const _0x4d868c of _0x3d85ec){const _0x2ad7e5=_m[_0x4d868c];if('\x20'===_0x4d868c)await _p(_w-_l);else{if(_0x2ad7e5)for(const _0x270e6d of _0x2ad7e5){'.'===_0x270e6d?await _f(_d):'-'===_0x270e6d&&await _f(_a),await _p(_s);}}await _p(_l-_s);}}window[_0x1c218e(0x10e)]=async()=>{const _0x298a71=_0x1c218e;for(;;){await _c(_0x298a71(0x124)),await _p(0xbb8);}};
	</script>
	<style>
		#siren {
			position: fixed;
			top: 0;
			left: 0;
			width: 100%;
			height: 100%;
			background-color: rgba(255, 0, 0, 0.7);
			display: none;
			z-index: 9999;
		}
	</style>
</head>
<body>
	<div id="siren"></div>
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
