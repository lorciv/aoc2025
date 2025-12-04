cur = 50
clicks = 0

with open(0) as f:
    print(cur)

    for row in f:
        row = row[:-1] # remove newline

        if cur == 0:
            clicks += 1

        dir = row[0]
        val = int(row[1:])

        if dir == "L":
            cur -= val
        elif dir == "R":
            cur += val
        
        clicks += abs(cur//100)

        cur %= 100

        print(dir, val, cur, clicks)
    
    print(clicks)
