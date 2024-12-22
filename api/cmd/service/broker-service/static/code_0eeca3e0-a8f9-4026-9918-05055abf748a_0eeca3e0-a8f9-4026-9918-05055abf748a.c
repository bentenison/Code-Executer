#include <stdio.h>

// User function logic goes here
int removeDuplicates(int arr[], int size) {
    if (size == 0) return 0;
    int j = 0;
    for (int i = 1; i < size; i++) {
        if (arr[i] != arr[j]) {
            j++;
            arr[j] = arr[i];
        }
    }
    return j + 1;
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
        return 1; // Return 1 to indicate true
    } else {
        printf("false\n");
        return 0; // Return 0 to indicate false
    }
}