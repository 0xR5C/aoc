import re


def mapConversion(mapVal, value):
    for rangeVal in mapVal:
        destination, source, length = map(int, rangeVal.split())
        if source <= value and source+length-1 >= value:
            return destination + value-source
    return value

f = open("input5.txt")
input = f.read()

segments = input.split("\n\n")
seeds = re.findall(r'\d+', segments[0])

locations = []
for i in seeds:
    val = int(i)
    for mapVal in segments[1:]:
        ranges = re.findall(r'\d+ \d+ \d+', mapVal)
        convert = mapConversion(ranges, val)
        val = convert
    
    locations.append(val)

print(min(locations))
