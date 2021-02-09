import re

all_replacements = {}
molecule = ""

for line in open('advent19.txt').readlines():
    r = re.match('(.*) => (.*)', line)
    
    if r is not None:
        value = r.group(1)
        replacement = r.group(2)
    
        if not all_replacements.has_key(value):
            all_replacements[value] = [replacement]
        else:
            all_replacements[value].append(replacement)
            all_replacements[value].sort(reverse=True)

    elif line != "":
        molecule = line

def get_new_molecules(molecule, value, replacements):
    new_molecules = []

    for replacement in replacements:
        pos = 0
        
        while True:
            pos = molecule.find(value, pos)
            if pos == -1:
                break
            
            pre_string = ""
            if pos > 0:
                pre_string = molecule[0:pos]
                
            post_string = ""
            if (pos + len(value)) < len(molecule):
                post_string = molecule[pos + len(value):]

            new_molecule = pre_string + replacement + post_string
            
            new_molecules.append(new_molecule)
            pos += len(value)
        
    return new_molecules
    
visited = []

def search(current_molecule, step=1):

    # print (step * ' ') + current_molecule

    if current_molecule in visited:
        return 10000000000000
   
    visited.append(current_molecule)
    
    for replacement in all_replacements['e']:
        if current_molecule == replacement:
            print "Found at " + str(step)
            return step
    
    min_step = 10000000000000
    
    for value,replacements in all_replacements.iteritems():
    
        if value == 'e':
            continue
    
        new_molecule = current_molecule
        
        for replacement in replacements:
            new_molecule = new_molecule.replace(replacement, value)
            
        min_step = min(min_step, search(new_molecule, step + 1))
            # pos = 0
            # while True:
                # pos = current_molecule.find(replacement, pos)
                # if pos == -1:
                    # break
                
                # pre_string = ""
                # if pos > 0:
                    # pre_string = current_molecule[0:pos]
                    
                # post_string = ""
                # if (pos + len(replacement)) < len(current_molecule):
                    # post_string = current_molecule[pos + len(replacement):]

                # new_molecule = pre_string + value + post_string
                # print (step * ' ') + pre_string + "["+ replacement + "]" + post_string + " => " + pre_string + "["+ value + "]" + post_string
                
                # if len(new_molecule) < current_molecule:
                    # min_step = min(min_step, search(new_molecule, step+1))
                
                # new_molecules.append(new_molecule)
                # pos += len(replacement)
                
    return min_step

# def search(current_molecule, step=1):

    # # print (step * ' ') + current_molecule

    # if current_molecule == molecule:
        # print "Found at " + str(step)
        # return step
        
    # if len(current_molecule) >= len(molecule):
        # return None

    # # if current_molecule in visited:
        # # return 10000000000000
   
    # # visited.append(current_molecule)
    
    # min_step = 10000000000000
    
    # for value,replacements in all_replacements.iteritems():
        # for replacement in replacements:
            # pos = 0
            # while True:
                # pos = current_molecule.find(value, pos)
                # if pos == -1:
                    # break
                
                # pre_string = ""
                # if pos > 0:
                    # pre_string = current_molecule[0:pos]
                    
                # post_string = ""
                # if (pos + len(value)) < len(current_molecule):
                    # post_string = current_molecule[pos + len(value):]

                # new_molecule = pre_string + replacement + post_string
                # # print (step * ' ') + pre_string + "["+ value + "]" + post_string + " => " + pre_string + "["+ replacement + "]" + post_string
                
                # if len(new_molecule) < current_molecule:
                    # min_step = min(min_step, search(new_molecule, step+1))
                
                # new_molecules.append(new_molecule)
                # pos += len(value)
                
    # return min_step
    
# def search_backwards2():

    # stack = [(molecule,0)]
    
    # smallest_steps = 1000000000
    
    # while len(stack) != 0:
    
        # current = stack.pop()
        # if current[0] in visited:
            # continue
       
        # visited.append(current[0])

        # if current[0] == 'e':
            # print "Found " + str(current[1])
            # smallest_steps = min(smallest_steps, current[1])
            
        # for value,replacements in all_replacements.iteritems():
            # for replacement in replacements:
                # for new_molecule in get_new_molecules(current[0], replacement, [value]):
                    # stack.append((new_molecule, current[1] + 1))
                    
    # return smallest_steps
       
# def search_backwards(current_molecule=molecule, level=0):

    # if current_molecule in visited:
        # return 0

    # visited.append(current_molecule)
    
    # if current_molecule == 'e':
        # print "Found!"

    # for value,replacements in all_replacements.iteritems():
        # for replacement in replacements:
            # for new_molecule in get_new_molecules(current_molecule, replacement, [value]):
                # if len(new_molecule) < current_molecule:
                    # search_backwards(new_molecule, level+1)
                    
    # return 0

# def search():

    # stack = ['e']
    
    # while len(stack) != 0:
    
        # print len(stack)
    
        # current_molecule = stack.pop()
        # if current_molecule == molecule:
            # print "Found in " + len(stack)

        # if current_molecule in visited:
            # continue
        
        # visited.append(current_molecule)

        # if len(current_molecule) > molecule:
            # continue
        
        # for value,replacements in all_replacements.iteritems():
            # for new_molecule in get_new_molecules(current_molecule, value, replacements):
                # stack.append(new_molecule)

# def search2(current_molecule, used_replacements):

    # if current_molecule in visited:
        # return
    
    # visited.append(current_molecule)
    
    # if current_molecule == molecule:
        # print "Found!"

    # for value,replacements in all_replacements.iteritems():
        # if value in used_replacements:
            # continue

        # for new_molecule in get_new_molecules(current_molecule, value, replacements):
            # search2(new_molecule, used_replacements + [value])

                
                
# Part 1
new_molecules = []

for value,replacements in all_replacements.iteritems():
    for new_molecule in get_new_molecules(molecule, value, replacements):
        if new_molecule not in new_molecules:
            new_molecules.append(new_molecule)

print len(new_molecules)

# Part 2
print search(molecule)