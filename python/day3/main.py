from collections import defaultdict


data = open('input').read().strip()
lines = data.split('\n')
matrix = [[c for c in line] for line in lines]
N = len(matrix)
M = len(matrix[0])

p1 = 0
nums = defaultdict(list)
for i in range(len(matrix)):
    gears = set()
    num = 0
    c_part = False
    for j in range(len(matrix[i])+1):
        if j < M and matrix[i][j].isdigit():
            num = num*10 + int(matrix[i][j])
            for rr in [-1, 0, 1]:
                for cc in [-1, 0, 1]:
                    if 0 <= i+rr < N and 0<j+cc<M:
                        ch = matrix[i+rr][j+cc]
                        if not ch.isdigit() and ch != '.':
                            c_part = True
                        if ch == '*':
                            gears.add((i+rr, j+cc))
        elif num > 0:
            for gear in gears:
                nums[gear].append(num)
            if c_part:
                p1 += num
            num = 0
            c_part = False
            gears = set()

print(p1)

p2 = 0
for k, v in nums.items():
    if len(v) == 2:
        p2 += v[0] * v[1]
print(p2)
