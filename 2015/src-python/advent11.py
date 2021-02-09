def has_one_increasing_straight_of_three_characters(password):
    for i in range(2, len(password)):
        c1 = ord(password[i-2])
        c2 = ord(password[i-1])
        c3 = ord(password[i])
    
        if c1 == (c2-1) and c2 == (c3-1):
            return True

    return False

def has_no_invalid_characters(password):
    return 'i' not in password and 'o' not in password and 'l' not in password

def has_two_different_non_overlapping_pairs(password):
    pair_count = 0

    i = 1
    while i < len(password): 
        c1 = ord(password[i])
        c2 = ord(password[i-1])
    
        if c1 == c2:
            pair_count += 1
            i += 2
            if pair_count == 2:
                return True
        else:
            i += 1

    return False
    
def is_valid(password):
    return has_one_increasing_straight_of_three_characters(password) and \
           has_no_invalid_characters(password) and \
           has_two_different_non_overlapping_pairs(password)

def increment(password):

    if password[-1] != 'z':
        next_char = chr(ord(password[-1]) + 1)
        if not has_no_invalid_characters(str(next_char)):
            next_char = chr(ord(next_char) + 1)
    
        return password[:-1] + next_char
    else:
        return increment(password[:-1]) + 'a'

def get_next_password(password):
    password = increment(password)
    
    while not is_valid(password):
        password = increment(password)
    
    return password
        
print get_next_password('hepxcrrq')
print get_next_password(get_next_password('hepxcrrq'))
