import re
import pprint
import sys
pp = pprint.PrettyPrinter(indent=4)

'''
This program will take an input such as:
    3 * 5 + 12 - 2 * 4

and decide the parentheses association to produce the highest total output
    i.e ((3 * (5 + 12)) - 2) * 4 = 156

This uses a dynamic programming algorithm to avoid an n! runtime
'''

'''
Gets the min and max solutions for given inputs
0 2 5 9
  1 4 8
    3 7
      6
'''
def get_min_and_max(mins, maxs, ops, i, j):
    min_val = sys.maxint
    max_val = -sys.maxint

    for k in range(i-j):
        if ops[i-1-k] == '+':
            a = mins[i][i-k] + mins[i-1-k][j]
            b = mins[i][i-k] + maxs[i-1-k][j]
            c = maxs[i][i-k] + mins[i-1-k][j]
            d = maxs[i][i-k] + maxs[i-1-k][j]
        elif ops[i-1-k] == '-':
            a = mins[i-1-k][j] - mins[i][i-k]
            b = maxs[i-1-k][j] - mins[i][i-k]
            c = mins[i-1-k][j] - maxs[i][i-k]
            d = maxs[i-1-k][j] - maxs[i][i-k]
        else:
            a = mins[i][i-k] * mins[i-1-k][j]
            b = mins[i][i-k] * maxs[i-1-k][j]
            c = maxs[i][i-k] * mins[i-1-k][j]
            d = maxs[i][i-k] * maxs[i-1-k][j]

        max_val = max(max_val, a, b, c, d)
        min_val = min(min_val, a, b, c, d)

    return [min_val, max_val]


'''
Gets the indexes of the appropriate parantheses spots
'''
def get_parentheses(expression):
    # parse valid expressions
    if not re.match(r'(\d+\s?[\+\-\*]\s?)+\d+', expression):
        raise Exception('Invalid Expression')

    numbers = map(int, re.findall('\d+', expression))
    operators = re.findall('[\+\-\*]', expression)
    sub_mins = [[0 for j in range(i+1)] for i in range(len(numbers))]
    sub_maxs = [[0 for j in range(i+1)] for i in range(len(numbers))]

    for i in range(len(numbers)):
        sub_mins[i][i] = numbers[i]
        sub_maxs[i][i] = numbers[i]

        for j in reversed(range(i)):
            print "looping through i:" + str(i) + ", j:" + str(j)

            min_and_max = get_min_and_max(sub_mins, sub_maxs, operators, i, j)
            sub_mins[i][j] = min_and_max[0]
            sub_maxs[i][j] = min_and_max[1]

        pp.pprint(sub_mins)
        pp.pprint(sub_maxs)

    recreate_expression_with_parentheses(sub_mins, sub_maxs, numbers, operators)



'''
Recreates the expression for the user to view
'''
def recreate_expression_with_parentheses(sub_mins, sub_maxs, numbers, ops):
    # find which operation was done last
    print "Max result: {0}, can be acheived by expression:\n".format(str(sub_maxs[len(ops)][0]))
    ops_reverse_order = []
    search_value = sub_maxs[len(ops)][0]

    for i in range(len(numbers)):
        for j in range(len(numbers) - i):
            #@todo finish recreating path
    return true

expression = raw_input('What expression would you like evaluated?\n\n')
get_parentheses(expression)
