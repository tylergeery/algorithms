import sys

''' Gets the number of coins needed for changes using only 1, 5, & 10 '''
def get_change(n):
    coins = 0

    ''' Greedy solution starts with highest coin'''
    while (n > 0):
        if (n > 10):
            n = n - 10
        elif (n > 5):
            n = n - 5
        else:
            n = n - 1

        coins += 1

    return coins

if __name__ == '__main__':
    n = int(raw_input())
    print(get_change(n))
