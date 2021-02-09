import itertools
import sys

def magic_missile(player1, player2):
    player1['mana'] -= 53
    player2['hp'] -= 4
    
def drain(player1, player2):
    player1['mana'] -= 73
    player1['hp'] += 2
    player2['hp'] -= 2

def shield(player1, player2):
    player1[

