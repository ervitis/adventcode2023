from collections import defaultdict
data = open('input').read().strip()
lines = data.split('\n')

my_cards_raw = []
winning_cards = []
for line in lines:
    raw = [cards.strip() for cards in line.split('|')[0].split(':')]
    winning_cards.append(dict.fromkeys(raw[1].split()))

    my_cards_raw.append([cards.strip() for cards in line.split('|')[1].strip().split()])

p1 = 0
p2 = defaultdict(int)
winning_values = []
for i in range(len(my_cards_raw)):
    l1 = my_cards_raw[i]
    l2 = winning_cards[i]

    winning_values.append([k for k in l1 if k in l2])
idx = 0
for values in winning_values:
    p2[idx] += 1
    t = 1
    if len(values) > 0:
        for v in range(2, len(values) + 1):
            t *= 2
        p1 += t

    for j in range(len(values)):
        p2[idx+j+1] += p2[idx]
    idx += 1

print(p1)
print(sum(p2.values()))
