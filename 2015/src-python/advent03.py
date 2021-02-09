moves = open('advent03.txt').read()

def get_new_position(position, move):
    if move == '<':
        return [position[0] - 1, position[1]]
    if move == '>':
        return [position[0] + 1, position[1]]
    if move == '^':
        return [position[0], position[1] + 1]
    if move == 'v':
        return [position[0], position[1] - 1]

    print move

# Part 1
position = [0,0]
delivered = []

for move in moves:
    if position not in delivered:
        delivered.append(position)
        
    position = get_new_position(position, move)

print len(delivered)

# Part 1
santaPosition = [0,0]
robotPosition = [0,0]
delivered = []

for i in range(0, len(moves)):

    if i % 2 == 0:
        if santaPosition not in delivered:
            delivered.append(santaPosition)

        santaPosition = get_new_position(santaPosition, moves[i])
    else:
        if robotPosition not in delivered:
            delivered.append(robotPosition)

        robotPosition = get_new_position(robotPosition, moves[i])

print len(delivered)
