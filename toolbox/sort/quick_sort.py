import dummy_data
import math
import time

'''
Quick sort based on first, middle, last value average
rather than random index
'''
def quick_sort(arr):
    l = len(arr)
    f = []
    m = []
    s = []

    # handle easy cases
    if not l or l <= 1:
        return arr

    # get basis value from first, middle, last index positions
    comparison_value = int((arr[0] + arr[int(math.floor(l/2))] + arr[l-1]) / 3)

    for i in range(l):
        val = int(arr[i])

        if val < comparison_value:
            f.append(arr[i])
        elif val == comparison_value:
            m.append(arr[i])
        else:
            s.append(arr[i])

    # python makes merging arrays quite easy
    return quick_sort(f) + m + quick_sort(s)


# milliseconds time
ms_time = lambda: int(round(time.time() * 1000))


dd = dummy_data.get_evenly_distributed_array()
start = ms_time()
result = quick_sort(dd)
end = ms_time()
print "\n\n Quicksort for evenly distributed array took {0} ms\n\n".format(end-start)


dd = dummy_data.get_bottom_heavy_array()
start = ms_time()
result = quick_sort(dd)
end = ms_time()
print "\n\n Quicksort for bottom heavy array took {0} ms\n\n".format(end-start)


dd = dummy_data.get_large_array_with_few_values()
start = ms_time()
result = quick_sort(dd)
end = ms_time()
print "\n\n Quicksort for large array w/ few values took {0} ms\n\n".format(end-start)
