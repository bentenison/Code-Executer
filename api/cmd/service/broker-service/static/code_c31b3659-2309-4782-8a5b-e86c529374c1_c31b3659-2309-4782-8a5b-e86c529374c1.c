#include <stdio.h>

// User function logic goes here
int findMissingNumber(int arr[], int n) {
    int sum = (n * (n + 1)) / 2;
    int arr_sum = 0;
    for (int i = 0; i < n - 1; i++) {
        arr_sum += arr[i];
    }
    return sum - arr_sum;
}

// Test cases for the function
int main() {
    int test_cases[][5] = {{1, 2, 4, 5}, {1, 3, 4}, {2, 3, 4, 5, 6}, {3, 4, 5, 6}};
    int expected[] = {3, 2, 1, 2};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int result = findMissingNumber(test_cases[i], 5);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %d. Expected: %d, Got: %d\n", test_cases[i], expected[i], result);
        }
    }
    if (all_passed) {
        printf("true\n");
        return 1; // Return 1 to indicate true
    } else {
        printf("false\n");
        return 0; // Return 0 to indicate false
    }
}