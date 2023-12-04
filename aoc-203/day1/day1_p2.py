import re



nums= ['1', '2', '3', '4', '5', '6', '7', '8', '9']
numLiterals = {'one':1, 'two':2, 'three':3, 'four':4, 'five':5, 'six':6, 'seven':7, 'eight':8, 'nine':9}



def convertToInt(x):
    if x in nums:
        return int(x)
    else:
        return numLiterals[x]




# lookahead regex to include edge case of overlapping matches
numsReg = r'(?=([1-9]|one|two|three|four|five|six|seven|eight|nine))'

prog = re.compile(numsReg)

f = open("input1.txt", "r")
input = f.readlines()

sum = 0
for line in input:
    found = re.findall(prog, line)
    if not found:
        continue

    sum = sum + convertToInt(found[0]) * 10 + convertToInt(found[-1])

print(sum)


