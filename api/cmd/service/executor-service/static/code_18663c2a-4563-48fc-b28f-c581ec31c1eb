#include <stdio.h>

// User function logic goes here
int findMax(int arr[], int n) {
    int max = arr[0];
    for (int i = 1; i < n; i++) {
        if (arr[i] > max) {
            max = arr[i];
        }
    }
    return max;
}

// Test cases for the function
int main() {
    int test_cases[][6] = {
        {1, 2, 3, 5},
        {1, 3, 4, 5, 6},
        {1, 2, 4, 5, 6},
        {1, 3}
    };
    int expected[] = {5, 6, 6, 3};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int n = sizeof(test_cases[i]) / sizeof(test_cases[i][0]);
        int result = findMax(test_cases[i], n);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: ");
            for (int j = 0; j < n; j++) printf("%d ", test_cases[i][j]);
            printf(". Expected: %d, Got: %d\n", expected[i], result);
        }
    }
    if (all_passed) {
        return 1; // Return 1 to indicate true
    } else {
        return 0; // Return 0 to indicate false
    }
}