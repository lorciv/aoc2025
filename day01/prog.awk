BEGIN {
	POS = 50
	COUNT = 0
	print("initial position", POS)
}

{
	DIR = substr($1, 1, 1)
	DIST = substr($1, 2)

	if (DIR == "L") {
		DIST *= -1
	}

	POS += DIST
	
	COUNT += int(POS/100)
	if (POS == 0) {
		COUNT += 1
	}
	
	print($1, "position", POS, "count", COUNT)
}

END {
	print("count", COUNT)
}
