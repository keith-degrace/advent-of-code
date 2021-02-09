import re

aunts = []

for line in open('advent16.txt').readlines():
    r = re.match('Sue (.*): (.*): (.*), (.*): (.*), (.*): (.*)', line)

    aunt = {}
    aunt['number'] = r.group(1)
    aunt[r.group(2)] = int(r.group(3))
    aunt[r.group(4)] = int(r.group(5))
    aunt[r.group(6)] = int(r.group(7))
    
    aunts.append(aunt)

# Part 1
def is_candidate(aunt):
    
    if aunt.has_key('children') and aunt['children'] != 3:
        return False
    
    if aunt.has_key('cats') and aunt['cats'] != 7:
        return False

    if aunt.has_key('samoyeds') and aunt['samoyeds'] != 2:
        return False
        
    if aunt.has_key('pomeranians') and aunt['pomeranians'] != 3:
        return False
    
    if aunt.has_key('akitas') and aunt['akitas'] != 0:
        return False
    
    if aunt.has_key('vizslas') and aunt['vizslas'] != 0:
        return False
    
    if aunt.has_key('goldfish') and aunt['goldfish'] != 5:
        return False

    if aunt.has_key('trees') and aunt['trees'] != 3:
        return False
        
    if aunt.has_key('cars') and aunt['cars'] != 2:
        return False
    
    if aunt.has_key('perfumes') and aunt['perfumes'] != 1:
        return False

    return True
        
for aunt in aunts:
    if is_candidate(aunt):
        print aunt['number']
        
# Part 2
def is_candidate_v2(aunt):
    
    if aunt.has_key('children') and aunt['children'] != 3:
        return False
    
    if aunt.has_key('cats') and aunt['cats'] <= 7:
        return False

    if aunt.has_key('samoyeds') and aunt['samoyeds'] != 2:
        return False
        
    if aunt.has_key('pomeranians') and aunt['pomeranians'] >= 3:
        return False
    
    if aunt.has_key('akitas') and aunt['akitas'] != 0:
        return False
    
    if aunt.has_key('vizslas') and aunt['vizslas'] != 0:
        return False
    
    if aunt.has_key('goldfish') and aunt['goldfish'] >= 5:
        return False

    if aunt.has_key('trees') and aunt['trees'] <= 3:
        return False
        
    if aunt.has_key('cars') and aunt['cars'] != 2:
        return False
    
    if aunt.has_key('perfumes') and aunt['perfumes'] != 1:
        return False

    return True
        
for aunt in aunts:
    if is_candidate_v2(aunt):
        print aunt['number']
        