import re

reindeers = []

for line in open('advent14.txt').readlines():
    r = re.match('(.*) can fly (.*) km/s for (.*) seconds, but then must rest for (.*) seconds.', line)

    reindeer = {}
    reindeer['name'] = r.group(1)
    reindeer['speed'] = int(r.group(2))
    reindeer['fly_time'] = int(r.group(3))
    reindeer['rest_time'] = int(r.group(4))
    reindeer['state'] = 'fly'
    reindeer['state_timer'] = reindeer['fly_time']
    reindeer['position'] = 0
    reindeer['points'] = 0
    
    reindeers.append(reindeer)

def get_max_position(reindeers):
    max_position = 0
            
    for reindeer in reindeers:
        max_position = max(max_position, reindeer['position'])
        
    return max_position

    
for i in range(2503):

    for reindeer in reindeers:
    
        if reindeer['state'] == 'fly':
            reindeer['position'] += reindeer['speed']
            reindeer['state_timer'] -= 1
            if reindeer['state_timer'] == 0:
                reindeer['state'] = 'rest'
                reindeer['state_timer'] = reindeer['rest_time']
        else:
            reindeer['state_timer'] -= 1
            if reindeer['state_timer'] == 0:
                reindeer['state'] = 'fly'
                reindeer['state_timer'] = reindeer['fly_time']

    max_position = get_max_position(reindeers)
    for reindeer in reindeers:
        if reindeer['position'] == max_position:
            reindeer['points'] += 1


# Part 1
print get_max_position(reindeers)
    
# Part 2
max_points = 0
        
for reindeer in reindeers:
    max_points = max(max_points, reindeer['points'])
    
print max_points
        