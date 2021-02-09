import hashlib

def find_hash(input, prefix):

    counter = 1
    while 1:
        candidate = input + str(counter)

        m = hashlib.md5();
        m.update(candidate)

        if m.hexdigest().startswith(prefix):
            return counter

        counter = counter + 1

# Part 1
print find_hash('yzbqklnj', '00000')

# Part 2
print find_hash('yzbqklnj', '000000')
