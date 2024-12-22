#include <stdio.h>

// User function logic goes here
int sumOfEvens(int arr[], int n) {
    int sum = 0;
    for (int i = 0; i < n; i++) {
        if (arr[i] % 2 == 0) {
            sum += arr[i];
        }
    }
    return sum;
}

// Test cases for the function
int main() {
    int test_cases[][5] = {{1, 2, 3, 4, 5}, {2, 4, 6, 8, 10}, {1, 3, 5, 7, 9}, {10, 15, 20, 25, 30}};
    int expected[] = {6, 30, 0, 60};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int sum = sumOfEvens(test_cases[i], 5);
        if (sum != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: [");
            for (int j = 0; j < 5; j++) {
                printf("%d ", test_cases[i][j]);
            }
            printf("] Expected: %d, Got: %d\n", expected[i], sum);
        }
    }
    if (all_passed) {
        printf("true\n");
        return 1;
    } else {
        printf("false\n");
        return 0;
    }
}