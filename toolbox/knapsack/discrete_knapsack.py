items = [
    (2, 30),
    (3, 94),
    (6, 95),
    (5, 82)
]

'''
Items given to the algorithm are a list of tuples
(weight, value)
total_allowed_weight integer for total weight of knapsack

returns the total possible value from the knapsack
'''
def get_discrete_knapsack(items, total_allowed_weight):
    length = len(items)
    lookup_table = [[0 for i in range(total_allowed_weight+1)] for j in range(length+1)]
    pp.pprint(lookup_table)

    for i in range(1, length+1):
        for w in range(1, total_allowed_weight+1):
            # name values for readability
            current_item_weight = items[i-1][0]
            current_item_value = items[i-1][1]

            # initialize to previously best value
            lookup_table[i][w] = lookup_table[i-1][w]

            # make sure this item can fit in knapsack
            if (current_item_weight <= w):
                # calculate the value of including it in knapsack at current weight
                temp = current_item_value + lookup_table[i-1][w-current_item_weight]

                if (temp > lookup_table[i][w]):
                    # set new max value if this is optimal at given constraints (weight, value)
                    lookup_table[i][w] = temp

    return lookup_table[length][total_allowed_weight]

total_value = get_discrete_knapsack(items, 7)
