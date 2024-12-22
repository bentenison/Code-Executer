#include <stdio.h>

// User function logic goes here
int findSecondLargest(int arr[], int n) {
    int largest = arr[0], second_largest = arr[0];
    for (int i = 1; i < n; i++) {
        if (arr[i] > largest) {
            second_largest = largest;
            largest = arr[i];
        } else if (arr[i] > second_largest && arr[i] != largest) {
            second_largest = arr[i];
        }
    }
    return second_largest;
}

// Test cases for the function
int main() {
    int test_cases[][5] = {{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {10, 20, 30, 40, 50}, {7, 1, 8, 3, 2}};
    int expected[] = {4, 4, 40, 7};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int second_largest = findSecondLargest(test_cases[i], 5);
        if (second_largest != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: [");
            for (int j = 0; j < 5; j++) {
                printf("%d ", test_cases[i][j]);
            }
            printf("] Expected: %d, Got: %d\n", expected[i], second_largest);
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