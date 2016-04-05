'''
Lemma

gcd(a, b) = gcd(b, a`) = gcd(b, a)
    - a' being the remainder of a % b

Ex

481 % 117 = 13
117 % 13 = 0

gcd (481, 117) = 13
'''

x = 481;
y = 117;

def gcd(a, b):
    r = a % b

    if (r == 0):
        return b
    else:
        return gcd(b, r)

print gcd(x,y)
