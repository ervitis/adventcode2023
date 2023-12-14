raw_data = open('input').read().strip()
parts = [d.split(':')[1].strip().split() for d in raw_data.split('\n')]

p1 = 1
j = 0
while j < len(parts[0]):
    t, dmax = int(parts[0][j]), int(parts[1][j])
    n = 0

    for i in range(t-1):
        tt = t - i
        d = i * tt
        if d > dmax:
            n += 1

    p1 *= n
    j += 1

print(p1)
