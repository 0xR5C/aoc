

def min_power(x):

    sets = x.split("; ")
    max_red = 0
    max_green = 0
    max_blue = 0
    for i in sets: 
        temp_vals = {"red":0, "green":0, "blue":0}
        grab = i.split(", ")
        for j in grab:
            temp_vals[j.split()[1]] += int(j.split()[0])
            
        if max_red < temp_vals["red"]:
            max_red = temp_vals["red"]
        if max_green < temp_vals["green"]:
            max_green = temp_vals["green"]
        if max_blue < temp_vals["blue"]:
            max_blue = temp_vals["blue"]
    return max_red * max_green * max_blue

f = open("input2.txt", "r")
input = f.readlines()

sum = 0
for line in input:
    input = line.split(': ')
    sum += min_power(input[1])


print(sum)