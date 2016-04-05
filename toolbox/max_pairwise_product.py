highest = 0
second_highest = 0

with open('./pairwise_data.txt', 'r') as data:
    indy = 0
    count = 0

    for line in data:
        if (indy == 0):
            count = int(line)
            indy += 1
        else:
            print line
            # iterate over range of integers
            for num in line.split(' '):
                if (int(num) > highest):
                    second_highest = highest
                    highest = int(num)
                elif (int(num) > second_highest):
                    second_highest = int(num)

product = highest * second_highest

print 'highest numbers are {0} {1}\n'.format(highest, second_highest)
print 'product {0}'.format(product)
