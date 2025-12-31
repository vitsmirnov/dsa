import random
from datetime import datetime as dt
# import timeit


def main() -> None:
    # arr = [3, 5, 2, 1, 0, 4]
    arr = [i for i in range(10)]
    random.shuffle(arr)

    print(arr)
    # arr = quick_sort(arr)
    # permutate(arr)
    quick_sort4(arr)
    print(arr)

    test_sort(quick_sort4)


def timer(func):
    def wrapper(*args, **kwargs):
        start = dt.now()
        res = func(*args, **kwargs)
        t = dt.now() - start
        # print(f"t: {t.min}:{t.seconds}:{t.microseconds}")
        print(f"t: {t.total_seconds()}")
        return res

    return wrapper


@timer
def test_sort(sort_func) -> None:
    if sort_func(list()) != list():
        raise Exception("Sorting does not work (1)!")

    if sort_func([3]) != [3]:
        raise Exception("Sorting does not work (2)!")

    for _ in range(100):
        arr = [random.randint(-100_000, 100_000) for _ in range(10_000)]
        arr = sort_func(arr)
        if not is_sorted(arr):
            print(arr)
            raise Exception("Sorting does not work (3)!")


def is_sorted(arr: list) -> bool:
    if len(arr) <= 1:
        return True

    for i in range(1, len(arr)):
        if arr[i] < arr[i-1]:
            return False

    return True


def quick_sort(arr: list) -> list:
    if len(arr) <= 1:
        return arr

    pivot_idx = random.randrange(0, len(arr))
    pivot = arr[pivot_idx]
    left = list()
    right = list()

    for i, v in enumerate(arr):
        if i == pivot_idx:
            continue

        if v < pivot:
            left.append(v)
        else:
            right.append(v)

    left = quick_sort(left)
    right = quick_sort(right)

    return left + [pivot] + right


# 3 4 2 1
# 1 4 2 3
# 1 2 4 3

def partition(arr: list, start: int=0, end: int=None) -> int:
    # arr_len = len(arr)

    # if end is None:
    #     end = arr_len - 1

    # # if end - start + 1 <= 1 or arr_len <= 1:
    # if start >= end or arr_len <= 1:
    #     return start

    pivot_idx = random.randint(start, end)
    pivot = arr[pivot_idx]
    arr[end], arr[pivot_idx] = arr[pivot_idx], arr[end]

    # i1, i2 = start, start
    # while i2 < end:
    #     if arr[i2] < pivot:
    #         arr[i1], arr[i2] = arr[i2], arr[i1]
    #         i1 += 1
    #     i2 += 1

    # arr[end], arr[i1] = arr[i1], arr[end]
    # return i1

    b = start
    for i in range(start, end+0):
        if arr[i] < pivot:
            # if i != b:
            arr[i], arr[b] = arr[b], arr[i]
            b += 1

    arr[end], arr[b] = arr[b], arr[end]
    return b


def quick_sort2(arr: list, start: int=0, end: int=None) -> list:
    if end is None:
        end = len(arr) - 1

    # if end - start + 1 <= 1 or len(arr) <= 1:
    if start >= end:
        return arr

    pivot_idx = partition(arr, start=start, end=end)

    quick_sort2(arr, start=start, end=pivot_idx-1)
    quick_sort2(arr, start=pivot_idx+1, end=end)

    return arr



def quick_sort3(arr: list, start: int=0, end: int | None=None) -> list:
    if end is None:
        end = len(arr) - 1

    if end <= start:
        return arr

    pivot_idx = partition2(arr, start, end)
    quick_sort3(arr, start=start, end=pivot_idx-1)
    quick_sort3(arr, start=pivot_idx+1, end=end)

    return arr

# 3 4 1 2
# 1 4 3 2  # b=1
# 1 2 3 4

def partition2(arr: list, start: int, end: int) -> int:
    pivot_idx = random.randint(start, end)
    pivot = arr[pivot_idx]
    arr[end], arr[pivot_idx] = arr[pivot_idx], arr[end]

    b = start
    for i in range(start, end):
        if arr[i] < pivot:
            arr[b], arr[i] = arr[i], arr[b]
            b += 1

    arr[b], arr[end] = arr[end], arr[b]

    return b


def quick_sort4(arr: list, start: int=0, end: int | None=None) -> list:
    if end is None:
        end = len(arr) - 1

    if start >= end:
        return arr

    pivit_idx = partition4(arr, start, end)
    quick_sort4(arr, start=start, end=pivit_idx-1)
    quick_sort4(arr, start=pivit_idx+1, end=end)

    return arr


def partition4(arr: list, start: int, end: int) -> int:
    pivot_idx = random.randint(start, end)
    pivot = arr[pivot_idx]
    arr[end], arr[pivot_idx] = arr[pivot_idx], arr[end]

    pivot_idx = start
    for i in range(start, end):
        if arr[i] < pivot:
            arr[i], arr[pivot_idx] = arr[pivot_idx], arr[i]
            pivot_idx += 1

    arr[pivot_idx], arr[end] = arr[end], arr[pivot_idx]

    return pivot_idx


if __name__ == "__main__":
    main()
