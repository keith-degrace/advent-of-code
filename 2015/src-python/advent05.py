strings = open('advent05.txt').readlines()

# Part 1
def has_no_unwanted_letters(string):
    return not ("ab" in string or "cd" in string or "pq" in string or "xy" in string)

def has_3_vowels(string):
    vowel_count = 0
    for char in string:
        if char in "aeiou":
            vowel_count = vowel_count + 1
            
        if vowel_count == 3:
            return True

    return False

def has_repeating_letters(string):
    for i in range(1, len(string)):
        if string[i-1] == string[i]:
            return True
            
    return False

def is_nice_v1(string):
    return has_no_unwanted_letters(string) and has_3_vowels(string) and has_repeating_letters(string)
    
nice_string_count = 0

for string in strings:
    if is_nice_v1(string):
        nice_string_count = nice_string_count + 1
            
print nice_string_count
 

#Part 2

def has_repeating_non_overlapping_pair(string):
    for i in range(1, len(string)):
        pair = string[i-1] + string[i]

        rest = string[i+1:]
        
        if pair in rest:
            return True
            
    return False
    
def has_repeating_letter_seperating_by_exactly_one_letter(string):
    for i in range(2, len(string)):
        if string[i-2] == string[i]:
            return True

    return False
    
def is_nice_v2(string):
    return has_repeating_non_overlapping_pair(string) and has_repeating_letter_seperating_by_exactly_one_letter(string)

    for i in range(1, len(string)):
        pair = string[i-1] + string[i]

        rest = string[i+1:]
        
        if pair in rest:
            found = True
    
nice_string_count = 0

for string in strings:
    if is_nice_v2(string):
        nice_string_count = nice_string_count + 1
            
print nice_string_count
