# Part 1

total_memory_size = 0
total_string_size = 0

for line in open('advent08.txt', 'r').readlines():

    encoded_line = line.strip()

    total_memory_size += len(encoded_line)

    line_with_quotes = encoded_line[1:-1]
    decoded_line = line_with_quotes.decode('string_escape')
    total_string_size += len(decoded_line)
    
    
print total_memory_size - total_string_size


# Part 2

total_memory_size = 0
total_new_memory_size = 0

for line in open('advent08.txt', 'r').readlines():

    encoded_line = line.strip()
    
    new_encoded_line = line.strip()
    new_encoded_line = new_encoded_line.replace('\\', '\\\\')
    new_encoded_line = new_encoded_line.replace('"', '\\"')
    new_encoded_line = "\"" + new_encoded_line + "\""
    
    total_memory_size += len(encoded_line)
    total_new_memory_size += len(new_encoded_line)
    
print total_new_memory_size - total_memory_size
