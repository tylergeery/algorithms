import math
import random

LIST_LENGTH = 100000

'''
Returns an array of size 10000
Should be more or less evenly distributed
'''
def get_evenly_distributed_array():
    arr = []

    for i in range(LIST_LENGTH):
        arr.append(int(math.ceil(LIST_LENGTH * random.random())))

    return arr

'''
Returns an array of size 10000
Lower numbers are much more probable
'''
def get_bottom_heavy_array():
    arr = []

    for i in range(LIST_LENGTH):
        arr.append(int(math.ceil(i * random.random())))

    return arr

'''
Returns an array of size 10000
Should have no more than 5 unique values
'''
def get_large_array_with_few_values():
    arr = []
    options = [100, 500, 1000, 2500, 2]

    for i in range(LIST_LENGTH):
        arr.append(options[int(math.floor(len(options) * random.random()))])

    return arr
