import dummy_data
import math
import time

'''
Perform a typical recursive merge Sort
'''
def merge_sort(a):
    # handle empty list or list of length 1
    if not a or len(a) == 1:
        return a

    mid = int(math.floor(len(a)/2))
    first_half = merge_sort(a[:mid])
    second_half = merge_sort(a[mid:])

    return merge(first_half, second_half)

'''
Merge two arrays
'''
def merge(a,b):
    result = []

    while len(a) or len(b):
        if len(a) and len(b):
            if (a[0] <= b[0]):
                result.append(a.pop(0))
            else:
                result.append(b.pop(0))
        elif len(a):
            # only a has values left
            result.append(a.pop(0))
        else:
            # only b has values left
            result.append(b.pop(0))

    return result

# milliseconds time
ms_time = lambda: int(round(time.time() * 1000))


dd = dummy_data.get_evenly_distributed_array()
start = ms_time()
merge_sort(dd)
end = ms_time()
print "\n\n Merge sort for evenly distributed array took {0} ms\n\n".format(end-start)


dd = dummy_data.get_bottom_heavy_array()
start = ms_time()
merge_sort(dd)
end = ms_time()
print "\n\n Merge sort for bottom heavy array took {0} ms\n\n".format(end-start)


dd = dummy_data.get_large_array_with_few_values()
start = ms_time()
merge_sort(dd)
end = ms_time()
print "\n\n Merge sort for large array w/ few values took {0} ms\n\n".format(end-start)
