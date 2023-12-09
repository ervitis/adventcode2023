"""
seed 55


source 50 -> 97 inclusives
destination 52 -> 99 inclusives
 ---- position is from source -> 55 - 50 = 5 + 1 because they are inclusive
 ---- then is 56
 ---- destination is 52 + 5 = 57
 ---- so the soil is 57

"""


raw_data = open('input').read().strip()
parts = raw_data.split('\n\n')

seeds, *instructions = parts
others = instructions

seeds = [int(seed) for seed in seeds.split(':')[1].split()]

res = []


def calculate_seed_area(seed, instr):
    for ins in instr:
        dest, src, rng = [int(x) for x in ins.split()]
        if src <= seed < src + rng:
            return dest+(seed - src)
    return seed


for seed in seeds:
    for instruction in instructions:
        section = instruction.split('\n')[1:]
        seed = calculate_seed_area(seed, section)
    res.append(seed)


print(min(res))


class Function:
    def __init__(self, S):
        lines = S.split('\n')[1:] # throw away name
        # dst src sz
        self.tuples: list[tuple[int,int,int]] = [[int(x) for x in line.split()] for line in lines]
        #print(self.tuples)

    def apply_one(self, x: int) -> int:
        for (dst, src, sz) in self.tuples:
            if src<=x<src+sz:
                return x+dst-src
        return x

    # list of [start, end) ranges
    def apply_range(self, R):
        A = []
        for (dest, src, sz) in self.tuples:
            src_end = src+sz
            NR = []
            while R:
                # [st                                     ed)
                #          [src       src_end]
                # [BEFORE ][INTER            ][AFTER        )
                (st,ed) = R.pop()
                # (src,sz) might cut (st,ed)
                before = (st,min(ed,src))
                inter = (max(st, src), min(src_end, ed))
                after = (max(src_end, st), ed)
                if before[1]>before[0]:
                    NR.append(before)
                if inter[1]>inter[0]:
                    A.append((inter[0]-src+dest, inter[1]-src+dest))
                if after[1]>after[0]:
                    NR.append(after)
            R = NR
        return A+R

Fs = [Function(s) for s in others]

def f(R, o):
    A = []
    for line in o:
        dest,src,sz = [int(x) for x in line.split()]
        src_end = src+sz


P2 = []
pairs = list(zip(seeds[::2], seeds[1::2]))
for st, sz in pairs:
    # inclusive on the left, exclusive on the right
    # e.g. [1,3) = [1,2]
    # length of [a,b) = b-a
    # [a,b) + [b,c) = [a,c)
    R = [(st, st+sz)]
    for f in Fs:
        R = f.apply_range(R)
    #print(len(R))
    P2.append(min(R)[0])
print(min(P2))