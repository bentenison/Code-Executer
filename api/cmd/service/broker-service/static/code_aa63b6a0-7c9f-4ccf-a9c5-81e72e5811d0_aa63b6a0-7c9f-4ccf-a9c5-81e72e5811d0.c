#include <stdio.h>

// User function logic goes here
int findLargest(int a, int b) {
    if (a > b) {
        return a;
    }
    return b;
}

// Test cases for the function
int main() {
    int test_cases[5][2] = {{10, 20}, {15, 10}, {25, 25}, {5, 8}, {9, 4}};
    int expected[] = {20, 15, 25, 8, 9};
    int all_passed = 1;

    for (int i = 0; i < 5; i++) {
        int result = findLargest(test_cases[i][0], test_cases[i][1]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %d, %d. Expected: %d, Got: %d\n", test_cases[i][0], test_cases[i][1], expected[i], result);
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