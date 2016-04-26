l = []

def fib_fast(n):
    for i in range(n):
        index = i-1

        if (i <= 1):
            l.append(1)
        else:
            l.append(l[i-1] + l[i-2])

    return l[n-1]

print("Which fibonnaci number would you like?\n")
n = int(raw_input())
a = fib_fast(n)
print "The {0}th fibonacci number is {1}".format(n, a)
