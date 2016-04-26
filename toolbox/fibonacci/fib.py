def calc_fib(n):
    if (n <= 1):
        return n

    return calc_fib(n - 1) + calc_fib(n - 2)

print("Which fibonnaci number would you like?\n")
n = int(raw_input())
a = calc_fib(n)
print "The {0}th fibonacci number is {1}".format(n, a)
