#include <stdio.h>

// User function logic goes here
void quickSort(int arr[], int low, int high) {
    if (low < high) {
        int pi = partition(arr, low, high);
        quickSort(arr, low, pi - 1);
        quickSort(arr, pi + 1, high);
    }
}

// Test cases for the function
int main() {
    int test_cases[][5] = {{12, 11, 13, 5, 6}, {4, 3, 7, 2, 1}, {9, 8, 7, 5, 6}};
    int expected[][5] = {{5, 6, 11, 12, 13}, {1, 2, 3, 4, 7}, {5, 6, 7, 8, 9}};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        int n = sizeof(test_cases[i]) / sizeof(test_cases[i][0]);
        quickSort(test_cases[i], 0, n - 1);
        for (int j = 0; j < n; j++) {
            if (test_cases[i][j] != expected[i][j]) {
                all_passed = 0;
                break;
            }
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