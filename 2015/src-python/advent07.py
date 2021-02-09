instructions = {}
for line in open('advent07.txt').readlines():
    parsed = line.split('>')
    
    wire = parsed[1].strip()
    value = parsed[0][:-1].strip()
    
    instructions[wire] = value
    

known_signals = {}

# Part 1
def calculate_signal(value):

    if value.isdigit():
        return int(value)
        
    if known_signals.has_key(value):
        return known_signals[value]

    instruction = instructions[value]

    if "AND" in instruction:
        lhs = calculate_signal(instruction.split(" ")[0])
        rhs = calculate_signal(instruction.split(" ")[2])
        known_signals[value] = lhs & rhs
        return known_signals[value]
    
    if "OR" in instruction:
        lhs = calculate_signal(instruction.split(" ")[0])
        rhs = calculate_signal(instruction.split(" ")[2])
        known_signals[value] = lhs | rhs
        return known_signals[value]
    
    if "NOT" in instruction:
        value = calculate_signal(instruction.split(" ")[1])
        known_signals[value] = ~value
        return known_signals[value]
    
    if "RSHIFT" in instruction:
        value = calculate_signal(instruction.split(" ")[0])
        shift = int(instruction.split(" ")[2])
        known_signals[value] = int(value) >> shift
        return known_signals[value]
    
    if "LSHIFT" in instruction:
        value = calculate_signal(instruction.split(" ")[0])
        shift = int(instruction.split(" ")[2])
        known_signals[value] = int(value) << shift
        return known_signals[value]

    known_signals[value] = calculate_signal(instruction)
    return known_signals[value]
    
print calculate_signal("a")

# Part 2
known_signals = {}
instructions["b"] = "16076"

print calculate_signal("a")
