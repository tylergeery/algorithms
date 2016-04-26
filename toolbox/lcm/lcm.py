def lcm(a, b):
    #write your code here
    return a*b

if __name__ == '__main__':
    a, b = map(int, raw_input().split())
    print(lcm(a, b))
