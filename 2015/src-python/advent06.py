import re

instructions = []
re = re.compile('(turn on|turn off|toggle) ([0-9]+),([0-9]+) through ([0-9]+),([0-9]+)')
for line in open('advent06.txt').readlines():

    instruction = {}
    instruction['action'] = re.match(line).group(1)
    instruction['from_x'] = int(re.match(line).group(2))
    instruction['from_y'] = int(re.match(line).group(3))
    instruction['to_x'] = int(re.match(line).group(4))
    instruction['to_y'] = int(re.match(line).group(5))

    instructions.append(instruction)

# Part 1

grid = [[0 for x in range(1000)] for x in range(1000)] 

for instruction in instructions:

    if instruction['action'] == "turn on":
        for x in range(instruction['from_x'], instruction['to_x']+1):         
            for y in range(instruction['from_y'], instruction['to_y']+1):
                grid[x][y] = 1
    elif instruction['action'] == "turn off":
        for x in range(instruction['from_x'], instruction['to_x']+1):         
            for y in range(instruction['from_y'], instruction['to_y']+1):
                grid[x][y] = 0
    elif instruction['action'] == "toggle":
        for x in range(instruction['from_x'], instruction['to_x']+1):         
            for y in range(instruction['from_y'], instruction['to_y']+1):
                if grid[x][y] == 0:
                    grid[x][y] = 1
                else:
                    grid[x][y] = 0
            
light_count = 0

for x in range(1000):
    for y in range(1000):
        if grid[x][y] == 1:
            light_count = light_count + 1
            
print light_count            

# Part 2
grid = [[0 for x in range(1000)] for x in range(1000)] 

for instruction in instructions:

    if instruction['action'] == "turn on":
        for x in range(instruction['from_x'], instruction['to_x']+1):         
            for y in range(instruction['from_y'], instruction['to_y']+1):
                grid[x][y] = grid[x][y] + 1
    elif instruction['action'] == "turn off":
        for x in range(instruction['from_x'], instruction['to_x']+1):         
            for y in range(instruction['from_y'], instruction['to_y']+1):
                if grid[x][y] > 0:
                    grid[x][y] = grid[x][y] - 1
    elif instruction['action'] == "toggle":
        for x in range(instruction['from_x'], instruction['to_x']+1):         
            for y in range(instruction['from_y'], instruction['to_y']+1):
                grid[x][y] = grid[x][y] + 2
            
total_brightness = 0

for x in range(1000):
    for y in range(1000):
        total_brightness = total_brightness + grid[x][y]
            
print total_brightness            

    
    