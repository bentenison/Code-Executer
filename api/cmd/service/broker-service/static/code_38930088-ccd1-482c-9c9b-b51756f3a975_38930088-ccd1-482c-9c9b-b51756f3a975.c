#include <stdio.h>

// User function logic goes here
int searchElement(int arr[], int size, int element) {
    int left = 0, right = size - 1;
    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (arr[mid] == element) {
            return mid;
        } else if (arr[mid] < element) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return -1;
}

// Test cases for the function
int main() {
    int test_cases[][6] = {{1, 1, 2, 3, 3, 4}, {1, 1, 1, 1, 1, 1}, {2, 3, 4, 4, 5, 6}};
    int expected[] = {4, 1, 5};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        int result = removeDuplicates(test_cases[i], 6);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: {%d, %d, %d, %d, %d, %d}. Expected: %d, Got: %d\n", test_cases[i][0], test_cases[i][1], test_cases[i][2], test_cases[i][3], test_cases[i][4], test_cases[i][5], expected[i], result);
        }
    }

    if (all_passed) {
        printf("true\n");
    } else {
        printf("false\n");
    }
    return 0;
}