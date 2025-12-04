total_jolts = 0

with open(0) as f:
    for line in f:
        bank = line[:-1]
        max_jolts = 0
        for i in range(len(bank)):
            for j in range(i+1, len(bank)):
                jolts = int(bank[i] + bank[j])
                if jolts > max_jolts:
                    max_jolts = jolts
        print(bank, "->", max_jolts)
        total_jolts += max_jolts
print("total", total_jolts)
