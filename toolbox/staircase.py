# import pprint
# pp = pprint.PrettyPrinter(indent=4)
'''
A person can walk up a staircase by 1 or 2 steps at a time
Find the total number of ways the person can climb a staircase with a given number of steps
'''
number_of_steps = 50000

'''
n - number of steps
'''
def climb(n):
    lookup = {
        0: 0,
        1: 1,
        2: 2
    }

    if n <= 2:
        return lookup[n]

    for i in range(3,n):
        lookup[i] = lookup[i-1] + lookup[i-2]

    # pp.pprint(lookup)
    print "Total possible ways: " + str(lookup[n-1] + lookup[n-2])


climb(number_of_steps)
