
nums= ['1', '2', '3', '4', '5', '6', '7', '8', '9']

f = open("input1.txt", "r")
input = f.readlines()

sum = 0
for line in input:
    found = [int(x) for x in line if x in nums]
    if len(found) == 0:
        continue
    else:
        sum = sum + (found[0] * 10 + found[-1])

print(sum)

