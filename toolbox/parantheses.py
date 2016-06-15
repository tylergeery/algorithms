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
            min_and_max = get_min_and_max(sub_mins, sub_maxs, operators, i, j)
            sub_mins[i][j] = min_and_max[0]
            sub_maxs[i][j] = min_and_max[1]

    recreate_expression_with_parentheses(sub_mins, sub_maxs, numbers, operators)



'''
Recreates the expression for the user to view
'''
def recreate_expression_with_parentheses(sub_mins, sub_maxs, numbers, ops):
    # find which operation was done last
    h = 0
    v = len(ops)
    ops_reverse_order = []
    search_values = []
    search_values.append([h,v,sub_maxs[v][0]])

    print "Max result: {0}, can be acheived by expression:\n".format(str(sub_maxs[len(ops)][0]))

    '''
    1
    2 3
    4 5 6
    7 8 9 1
    1 2 3 4 5

    Go down horizontal and vertically from position of each matrix,
    find pair that create
    '''
    while (len(ops_reverse_order) < len(ops)):
        value = search_values.pop(0)

        h = value[0]
        v = value[1]
        target = value[2]

        for i in range(h, len(sub_maxs[v])):
            if (ops[i] == '+'):
                a = sub_maxs[i][h] + sub_maxs[v][i+1]
                b = sub_maxs[i][h] + sub_mins[v][i+1]
                c = sub_mins[i][h] + sub_maxs[v][i+1]
                d = sub_mins[i][h] + sub_mins[v][i+1]
            elif (ops[i] == '-'):
                a = sub_maxs[i][h] - sub_maxs[v][i+1]
                b = sub_maxs[i][h] - sub_mins[v][i+1]
                c = sub_mins[i][h] - sub_maxs[v][i+1]
                d = sub_mins[i][h] - sub_mins[v][i+1]
            else:
                a = sub_maxs[i][h] * sub_maxs[v][i+1]
                b = sub_maxs[i][h] * sub_mins[v][i+1]
                c = sub_mins[i][h] * sub_maxs[v][i+1]
                d = sub_mins[i][h] * sub_mins[v][i+1]

            if a == target:
                if h != i:
                    search_values.append([h, i, sub_maxs[i][h]])
                if (i+1) != v:
                    search_values.append([i+1, v, sub_maxs[v][i+1]])

                ops_reverse_order.append(i)

                break
            if b == target:
                if h != i:
                    search_values.append([h, i, sub_maxs[i][h]])
                if (i+1) != v:
                    search_values.append([i+1, v, sub_mins[v][i+1]])

                ops_reverse_order.append(i)

                break
            if c == target:
                if h != i:
                    search_values.append([h, i, sub_mins[i][h]])
                if (i+1) != v:
                    search_values.append([i+1, v, sub_maxs[v][i+1]])

                ops_reverse_order.append(i)

                break
            if d == target:
                if h != i:
                    search_values.append([h, i, sub_mins[i][h]])
                if (i+1) != v:
                    search_values.append([i+1, v, sub_mins[v][i+1]])

                ops_reverse_order.append(i)

                break

    generate_expression_string(ops_reverse_order, ops, numbers)

def generate_expression_string(ops_reverse_order, ops, numbers):
    pp.pprint(ops_reverse_order)
    op_expressions = {}
    first_op = ops_reverse_order.pop(0)
    op_expressions[first_op] = ops[first_op]

    for op in ops_reverse_order:
        if ((op-1) != first_op and (op-1) in op_expressions):
            before = op_expressions[op-1]

            op_expressions[op-1] = op

            # chase down head as if linked list
            while isinstance(before, int):
                print 'instance before:{0}'.format(str(before))
                tmp = before
                before = op_expressions[before]
                op_expressions[tmp] = op
        else:
            before = numbers[op]

        if ((op+1) != first_op and (op+1) in op_expressions):
            after = op_expressions[op+1]

            op_expressions[op+1] = op

            # chase down head as if linked list
            while isinstance(after, int):
                print 'instance after:{0}'.format(str(before))
                tmp = after
                after = op_expressions[after]
                op_expressions[tmp] = op
        else:
            after = numbers[op+1]

        op_expressions[op] = '({0}{1}{2})'.format(str(before), ops[op], str(after))

    expr = ''
    if (first_op == 0):
        expr += str(numbers[0])
    if (first_op-1) in op_expressions:
        tmp = op_expressions[first_op-1]
        while isinstance(tmp, int):
            tmp = op_expressions[tmp]

        expr += tmp

    expr += str(op_expressions[first_op])

    if (first_op+1) in op_expressions:
        tmp = op_expressions[first_op+1]
        while isinstance(tmp, int):
            tmp = op_expressions[tmp]

        expr += tmp
    if (first_op == len(ops_reverse_order)):
        expr += str(numbers[first_op+1])

    print expr

expression = raw_input('What expression would you like evaluated?\n\n')
get_parentheses(expression)
