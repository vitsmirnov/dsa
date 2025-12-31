import random


def main() -> None:
    for _ in range(10):
        # arr = [2, 0, 5, 4, 3, 1]
        arr = [random.randint(-1_000_000, 1_000_000) for _ in range(100_000)]
        # arr = [i for i in range(10)]
        # print(arr)
        # print(is_sorted(arr))
        # arr = quicksort(arr)
        quicksort2(arr)
        # print(arr)
        print(is_sorted(arr))
        # print()


def is_sorted(arr: list) -> bool:
    if len(arr) <= 1:
        return True

    for i in range(1, len(arr)):
        if arr[i] < arr[i-1]:
            return False

    return True


def quicksort(arr: list) -> list:
    arr_len = len(arr)
    if arr_len <= 1:
        return arr

    pivot_idx = random.randrange(arr_len)
    pivot = arr[pivot_idx]
    left = list()
    right = list()
    middle = list()
    for v in arr:
        if v < pivot:
            left.append(v)
        elif v > pivot:
            right.append(v)
        else:
            middle.append(v)
    # print(left, middle, right)

    left = quicksort(left)
    right = quicksort(right)

    return left + middle + right


def quicksort2(arr: list, start_idx: int=0, end_idx: int | None=None) -> None:
    if end_idx is None:
        end_idx = len(arr) - 1

    arr_len = end_idx - start_idx + 1
    if arr_len <= 1:
        return

    pivot_idx = partition2(arr, start_idx, end_idx)
    quicksort2(arr, start_idx, pivot_idx-1)
    quicksort2(arr, pivot_idx+1, end_idx)


# 3, 0, 5, 2, 4  (3: 2)


def partition2(arr: list, start_idx: int, end_idx: int) -> int:
    """ Returns pivot index """

    pivot_idx = random.randint(start_idx, end_idx)
    pivot = arr[pivot_idx]
    arr[pivot_idx], arr[end_idx] = arr[end_idx], arr[pivot_idx]
    
    pivot_idx = start_idx
    for i in range(start_idx, end_idx):
        if arr[i] < pivot:
            arr[i], arr[pivot_idx] = arr[pivot_idx], arr[i]
            pivot_idx += 1

    arr[end_idx], arr[pivot_idx] = arr[pivot_idx], arr[end_idx]
    return pivot_idx


def partition(arr: list, start_idx: int, end_idx: int) -> int:
    """ Returns pivot index """

    pivot_idx = random.randint(start_idx, end_idx)
    pivot = arr[pivot_idx]
    s, e = start_idx, end_idx
    print("1", arr[start_idx:end_idx+1], pivot)
    while True:
        while arr[start_idx] < pivot and start_idx < end_idx:
            start_idx += 1
        while arr[end_idx] >= pivot and start_idx < end_idx:
            end_idx -= 1
        if start_idx >= end_idx:
            break
        arr[start_idx], arr[end_idx] = arr[end_idx], arr[start_idx]
        print(" ", arr[s:e+1], start_idx, end_idx)
        start_idx += 1
        end_idx -= 1
    pivot_idx = end_idx
    arr[pivot_idx], arr[e] = arr[e], arr[pivot_idx]
    print("2", arr[s:e+1], pivot)
    print()

    return pivot_idx


if __name__ == "__main__":
    main()
