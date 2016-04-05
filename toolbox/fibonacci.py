n = 500
l = []

for i in range(n):
    index = i-1

    if (i <= 1):
        l.append(1)
    else:
        l.append(l[i-1] + l[i-2])

print "The {0}th fibonacci number is {1}".format(n, l[n-1])
