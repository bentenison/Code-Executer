#include <stdio.h>

// User function logic goes here
int linearSearch(int arr[], int n, int target) {
    for (int i = 0; i < n; i++) {
        if (arr[i] == target) {
            return i;
        }
    }
    return -1;
}

// Test cases for the function
int main() {
    int test_cases[][5] = {{10, 20, 30, 40, 50}, {5, 15, 25, 35, 45}, {1, 2, 3, 4, 5}};
    int targets[] = {30, 25, 6};
    int expected[] = {2, 2, -1};
    int n = 5;
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        int result = linearSearch(test_cases[i], n, targets[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: { ");
            for (int j = 0; j < n; j++) {
                printf("%d%s", test_cases[i][j], j == n - 1 ? " " : ", ");
            }
            printf("}. Target: %d. Expected: %d, Got: %d\n", targets[i], expected[i], result);
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