

f = open("input4.txt")
input = f.readlines()

scores = [1 for element in range(len(input))]

points = 0
for i in range(len(input)):
    winning, yours = input[i].split(": ")[1].split(" | ")
    winning = set(winning.split())
    yours =  set(yours.split())
    wins = len(winning & yours)
    points += int(2**(wins-1))
    for j in range(i+1, i+wins+1):
        scores[j] += 1 * scores[i]

print("Part 1: ", points)
print("Part 2: ", sum(scores))