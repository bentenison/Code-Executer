# User function logic goes here
def count_inversions(arr):
    def merge_count(arr, temp_arr, left, right):
        if left == right:
            return 0
        mid = (left + right) // 2
        inv_count = merge_count(arr, temp_arr, left, mid)
        inv_count += merge_count(arr, temp_arr, mid + 1, right)
        inv_count += merge(arr, temp_arr, left, mid, right)
        return inv_count

    def merge(arr, temp_arr, left, mid, right):
        i = left
        j = mid + 1
        k = left
        inv_count = 0
        while i <= mid and j <= right:
            if arr[i] <= arr[j]:
                temp_arr[k] = arr[i]
                i += 1
            else:
                temp_arr[k] = arr[j]
                inv_count += (mid-i + 1)
                j += 1
            k += 1
        while i <= mid:
            temp_arr[k] = arr[i]
            i += 1
            k += 1
        while j <= right:
            temp_arr[k] = arr[j]
            j += 1
            k += 1
        for i in range(left, right + 1):
            arr[i] = temp_arr[i]
        return inv_count

    n = len(arr)
    temp_arr = [0] * n
    return merge_count(arr, temp_arr, 0, n-1)

# Test cases for the function
if __name__ == '__main__':
    all_passed = True
    test_cases = [
        ([1, 20, 6, 4, 5],),
        ([5, 4, 3, 2, 1],),
    ]
    expected_outputs = [
        5,
        10,
    ]
    for test_input, expected in zip(test_cases, expected_outputs):
        result = count_inversions(test_input[0])
        if result != expected:
            all_passed = False
            print(f'Failed for Input: {test_input}. Expected: {expected}, Got: {result}')
    print(all_passed)