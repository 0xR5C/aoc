
# TODO: REFACTOR CODE


nums= ['1', '2', '3', '4', '5', '6', '7', '8', '9', '0']
gearsDict = {}

def isSymbol(x):
    return x not in nums and x != '.' and x != '\n'

f = open("input3.txt", "r")
input = f.readlines()

sum = 0
for i in range(0, len(input)):
    temp_num = -1
    temp_num_start = 0
    for j in range(0, len(input[i])):
        if temp_num == -1 and (input[i][j] == '.' or input[i][j] == '\n'):
            continue            
        elif input[i][j] in nums:
            if temp_num == -1:
                temp_num_start = j
                temp_num = 0
            temp_num = temp_num * 10 + int(input[i][j])
        elif isSymbol(input[i][j]) and temp_num != -1:
            if input[i][j] == "*":
                if (i, j) in gearsDict:
                    gearsDict[(i, j)].append(temp_num)
                else:
                    gearsDict[(i, j)] = [temp_num]
            sum += temp_num
            temp_num = -1
            continue
        elif temp_num != -1 and (input[i][j] == '.' or input[i][j] == '\n'):
            # check if it is a part number
            temp_num_end = j+1 if j+1 < len(input[i]) else j
            temp_num_start = temp_num_start-1 if j-1 >= 0 else temp_num_start

            #check same line (input[i])
            if isSymbol(input[i][temp_num_start]):
                if input[i][temp_num_start] == "*":
                    if (i, temp_num_start) in gearsDict:
                        gearsDict[(i, temp_num_start)].append(temp_num)
                    else:
                        gearsDict[(i, temp_num_start)] = [temp_num]
                sum += temp_num
                temp_num = -1
                continue

        
            # check previous line if you can (input[i-1])
            if i-1 >= 0:
                for k in range(temp_num_start, temp_num_end):
                    if isSymbol(input[i-1][k]):
                        if input[i-1][k] == "*":
                            if (i-1, k) in gearsDict:
                                gearsDict[(i-1, k)].append(temp_num)
                            else:
                                gearsDict[(i-1, k)] = [temp_num]
                        sum += temp_num
                        temp_num = -1
                        break  
                if temp_num == -1:
                    continue      
            
            if i+1 != len(input):
                for k in range(temp_num_start, temp_num_end):
                    if isSymbol(input[i+1][k]):
                        if input[i+1][k] == "*":
                            if (i+1, k) in gearsDict:
                                gearsDict[(i+1, k)].append(temp_num)
                            else:
                                gearsDict[(i+1, k)] = [temp_num]
                        sum += temp_num
                        temp_num = -1
                        break  
            temp_num = -1

sum2 = 0
for keys in gearsDict:
    print(keys, gearsDict[keys])
    if len(gearsDict[keys]) == 2:
        sum2 += gearsDict[keys][0] * gearsDict[keys][1]


print("Part 1 result: ", sum)
print("Part 2 result: ", sum2)