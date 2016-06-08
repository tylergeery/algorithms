from __future__ import division

'''
Get The optimal value based on weights, capacity, and values of each items
Greedy algorithm value/weight

Items can't be split
'''
def get_optimal_value(capacity, items):
    value = 0.

    # Sort items by value/weight
    items.sort(key=lambda item: item[0]/item[1], reverse=True)

    # figure out how much value can fit in capacity
    i = 0
    while (capacity > 0 and i < len(items)):
        space_for_item = min(items[i][1], capacity)
        value_from_item = space_for_item * items[i][0] / items[i][1]

        capacity -= space_for_item
        value += value_from_item
        i += 1

    return value


if __name__ == "__main__":
    data = []
    while True:
        txt = raw_input()

        if (txt == ""):
            break
        data.append(map(int, txt.split()))

    # first line is # items && capacity
    n, capacity = data[0][0:2]

    opt_value = get_optimal_value(capacity, data[1:])
    print("{:.10f}".format(opt_value))
