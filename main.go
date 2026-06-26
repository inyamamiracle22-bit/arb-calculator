package main

import (
	"fmt"
	"net/http"
	"os"
)

func page(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head>
<title>Arb Calculator</title>
<meta name="viewport" content="width=device-width, initial-scale=1.0">

<style>
	body {
		font-family: Arial;
		background: #0d0b20ff;
		color: white;
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
		margin: 0;
	}

	.app {
		width: 100%;
		max-width: 900px;
		background: #111a2e;
		padding: 20px;
		border-radius: 16px;
	}

	.title {
		text-align: center;
		color: #00ff7b;
		font-weight: bold;
		margin-bottom: 10px;
	}

	.top {
		text-align: center;
		font-size: 18px;
		font-weight: bold;
		color: #00ff7b;
		margin-bottom: 6px;
	}

	input {
		width: 100%;
		padding: 10px;
		margin-top: 5px;
		border-radius: 8px;
		border: none;
		background: #0b1220;
		color: white;
		box-sizing: border-box;
	}

	.row {
		display: flex;
		gap: 10px;
	}

	.card {
		flex: 1;
		background: #16233d;
		padding: 10px;
		border-radius: 10px;
	}

	button {
		width: 100%;
		padding: 12px;
		margin-top: 10px;
		border-radius: 8px;
		border: none;
		font-weight: bold;
		cursor: pointer;
	}

	.calc { background: #00ff7b; }
	.quick { background: #ffa500; }
	.clear { background: #ff4d4d; }
</style>
</head>

<body>

<div class="app">

	<div class="title">ARBITRAGE CALCULATOR</div>

	<div class="top" id="profit">PROFIT: 0%</div>
	<div class="top" id="outcome">TOTAL OUTCOME: ₦0</div>

	<select id="mode">
		<option value="3">3-WAY MARKET</option>
		<option value="2">2-WAY MARKET</option>
	</select>

	<div class="row">
		<div class="card">
			Odd 1
			<input id="o1">
			Stake 1
			<input id="s1">
		</div>

		<div class="card">
			Odd 2
			<input id="o2">
			Stake 2
			<input id="s2">
		</div>

		<div class="card" id="box3">
			Odd 3
			<input id="o3">
			Stake 3
			<input id="s3">
		</div>
	</div>

	<br>

	Total Stake
	<input id="total">

	<!-- FULL CALC (YOUR ORIGINAL SYSTEM) -->
	<button class="calc" onclick="calcFull()">CALCULATE FULL STAKE SYSTEM</button>

	<!-- NEW QUICK MODE -->
	<button class="quick" onclick="calcQuick()">CHECK ARB % (ODDS ONLY)</button>

	<button class="clear" onclick="clearAll()">CLEAR</button>

</div>

<script>

let last = "total"

// keep original behavior
document.getElementById("total").addEventListener("input", () => last = "total")
document.getElementById("s1").addEventListener("input", () => last = "s1")
document.getElementById("s2").addEventListener("input", () => last = "s2")
document.getElementById("s3").addEventListener("input", () => last = "s3")

document.getElementById("mode").addEventListener("change", () => {
	let mode = document.getElementById("mode").value
	document.getElementById("box3").style.display = (mode == "2") ? "none" : "block"
})


// =========================
// YOUR ORIGINAL CALCULATOR
// =========================
function calcFull() {

	let mode = document.getElementById("mode").value

	let o1 = parseFloat(document.getElementById("o1").value)
	let o2 = parseFloat(document.getElementById("o2").value)
	let o3 = parseFloat(document.getElementById("o3").value)

	if (isNaN(o1) || isNaN(o2)) return
	if (mode == "3" && isNaN(o3)) return

	let s1 = parseFloat(document.getElementById("s1").value)
	let s2 = parseFloat(document.getElementById("s2").value)
	let s3 = parseFloat(document.getElementById("s3").value)
	let total = parseFloat(document.getElementById("total").value)

	let p1=0, p2=0, p3=0

	if (last === "total" && !isNaN(total)) {

		if (mode == "2") {
			let sum = (1/o1) + (1/o2)
			p1 = total * (1/o1) / sum
			p2 = total * (1/o2) / sum
		} else {
			let sum = (1/o1) + (1/o2) + (1/o3)
			p1 = total * (1/o1) / sum
			p2 = total * (1/o2) / sum
			p3 = total * (1/o3) / sum
		}

	} else if (last === "s1" && !isNaN(s1)) {

		p1 = s1
		p2 = s1 * (o1/o2)
		p3 = (mode == "3") ? s1 * (o1/o3) : 0
		total = p1 + p2 + p3

	} else if (last === "s2" && !isNaN(s2)) {

		p2 = s2
		p1 = s2 * (o2/o1)
		p3 = (mode == "3") ? s2 * (o2/o3) : 0
		total = p1 + p2 + p3

	} else if (last === "s3" && !isNaN(s3)) {

		p3 = s3
		p1 = s3 * (o3/o1)
		p2 = s3 * (o3/o2)
		total = p1 + p2 + p3
	}

	let payout = p1 * o1
	let profit = ((payout / total) - 1) * 100

	document.getElementById("s1").value = p1.toFixed(2)
	document.getElementById("s2").value = p2.toFixed(2)
	document.getElementById("s3").value = p3.toFixed(2)
	document.getElementById("total").value = total.toFixed(2)

	document.getElementById("profit").innerText =
	"PROFIT: " + profit.toFixed(2) + "%"

	document.getElementById("outcome").innerText =
	"TOTAL OUTCOME: ₦" + payout.toFixed(2)
}


// =========================
// NEW QUICK ODDS-ONLY MODE
// =========================
function calcQuick() {

	let mode = document.getElementById("mode").value

	let o1 = parseFloat(document.getElementById("o1").value)
	let o2 = parseFloat(document.getElementById("o2").value)
	let o3 = parseFloat(document.getElementById("o3").value)

	if (isNaN(o1) || isNaN(o2)) return
	if (mode == "3" && isNaN(o3)) return

	let total = 100 // virtual stake

	let p1=0, p2=0, p3=0

	if (mode == "2") {

		let sum = (1/o1) + (1/o2)
		p1 = total * (1/o1) / sum
		p2 = total * (1/o2) / sum

	} else {

		let sum = (1/o1) + (1/o2) + (1/o3)
		p1 = total * (1/o1) / sum
		p2 = total * (1/o2) / sum
		p3 = total * (1/o3) / sum
	}

	let payout = p1 * o1
	let profit = ((payout / total) - 1) * 100

	document.getElementById("profit").innerText =
	"ARB % (ODDS ONLY): " + profit.toFixed(2) + "%"

	document.getElementById("outcome").innerText =
	"TOTAL OUTCOME: ₦" + payout.toFixed(2)
}


// CLEAR
function clearAll() {

	document.getElementById("o1").value = ""
	document.getElementById("o2").value = ""
	document.getElementById("o3").value = ""

	document.getElementById("s1").value = ""
	document.getElementById("s2").value = ""
	document.getElementById("s3").value = ""
	document.getElementById("total").value = ""

	document.getElementById("profit").innerText = "PROFIT: 0%"
	document.getElementById("outcome").innerText = "TOTAL OUTCOME: ₦0"

	last = "total"
}

</script>

</body>
</html>`)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", page)

	fmt.Println("Server running on port", port)
	http.ListenAndServe(":"+port, nil)
}
