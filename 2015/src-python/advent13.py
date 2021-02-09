import re
import itertools

guests = []
hapiness_stats = {}

for line in open('advent13.txt').readlines():
    r = re.match('(.*) would (.*) ([0-9]*) happiness units by sitting next to (.*).', line)
    
    guest1 = r.group(1)
    if r.group(2) == 'gain':
        hapiness = int(r.group(3))
    else:
        hapiness = -int(r.group(3))
    guest2 = r.group(4)

    if guest1 not in guests:
        guests.append(guest1)
        
    if guest2 not in guests:
        guests.append(guest2)

    hapiness_stats[(guest1, guest2)] = hapiness

    
def calculate_total_hapiness(guests):
    
    total_hapiness = 0
    
    for i in range(len(guests)):
    
        guest1 = guests[i]
        guest2 = guests[(i + 1) % len(guests)]
    
        total_hapiness += hapiness_stats[(guest1, guest2)]
        total_hapiness += hapiness_stats[(guest2, guest1)]

    return total_hapiness
    
# Part 1
    
optimal_total_hapiness = 0

for permutation in itertools.permutations(guests):
    optimal_total_hapiness = max(optimal_total_hapiness, calculate_total_hapiness(permutation))
    
print optimal_total_hapiness

# Part 2

for guest in guests:
    hapiness_stats[('me', guest)] = 0
    hapiness_stats[(guest, 'me')] = 0

guests.append('me')
    
optimal_total_hapiness = 0

for permutation in itertools.permutations(guests):
    optimal_total_hapiness = max(optimal_total_hapiness, calculate_total_hapiness(permutation))
    
print optimal_total_hapiness
