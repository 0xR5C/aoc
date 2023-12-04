

def possible(x):

    sets = x.split("; ")
    for i in sets:
        counts = {"red":0, 'green':0, "blue":0}
        grab = i.split(", ")
        for j in grab:
            counts[j.split()[1]] += int(j.split()[0])
            if counts["red"] > 12 or counts["green"] > 13 or counts["blue"] > 14:
                return False
    return True

f = open("input2.txt", "r")
input = f.readlines()

sum = 0
for line in input:
    input = line.split(': ')
    id = int(input[0].split(" ")[1])
    if possible(input[1]):
        sum += id


print(sum)