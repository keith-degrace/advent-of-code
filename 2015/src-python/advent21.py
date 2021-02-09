import itertools
import sys

weapons = [[8,4], [10,5], [25,5], [40,7], [74,8]]
armors  = [[0,0], [13,1], [31,2], [53,3], [75,4], [102,5]]
rings   = [[0,0,0], [0,0,0], [25,1,0], [50,2,0], [100,3,0], [20,0,1], [40,0,2], [80,0,3]]

def attack(player1, player2):
    player2['hp'] -= max(1, player1['damage'] - player2['armor'])

def fight(player1, player2):
    while True:
        attack(player1, player2)
        if player2['hp'] <= 0:
            return True
            
        attack(player2, player1)
        if player1['hp'] <= 0:
            return False

suits = []
for weapon in weapons:
    for armor in armors:
        for ring1 in rings:
            for ring2 in rings:
                if ring2 == ring1:
                    continue
                
                suit = {}
                suit['damage'] = weapon[1] + ring1[1] + ring2[1]
                suit['armor'] = armor[1] + ring1[2] + ring2[2]
                suit['cost'] = weapon[0] + armor[0] + ring1[0] + ring2[0]
                
                suits.append(suit)

min_cost = sys.maxint

for suit in suits:
    player1 = {}
    player1['name'] = 'Player'
    player1['hp'] = 100
    player1['damage'] = suit['damage']
    player1['armor'] = suit['armor']
                   
    boss = {}
    boss['name'] = 'Boss'
    boss['hp'] = 109
    boss['damage'] = 8
    boss['armor'] = 2
    
    if fight(player1, boss):
        min_cost = min(min_cost, suit['cost'])

print min_cost


max_cost = 0

for suit in suits:
    player1 = {}
    player1['name'] = 'Player'
    player1['hp'] = 100
    player1['damage'] = suit['damage']
    player1['armor'] = suit['armor']
                   
    boss = {}
    boss['name'] = 'Boss'
    boss['hp'] = 109
    boss['damage'] = 8
    boss['armor'] = 2
    
    if not fight(player1, boss):
        max_cost = max(max_cost, suit['cost'])

print max_cost