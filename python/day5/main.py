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
