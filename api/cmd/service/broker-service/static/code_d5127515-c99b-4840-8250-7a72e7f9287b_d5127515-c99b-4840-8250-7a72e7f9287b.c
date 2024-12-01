#include <stdio.h>

// User function logic goes here
int findMax(int a, int b, int c) {
    if (a >= b && a >= c) {
        return a;
    } else if (b >= a && b >= c) {
        return b;
    } else {
        return c;
    }
}

// Test cases for the function
int main() {
    int test_cases[5][3] = {{10, 20, 15}, {1, 1, 1}, {5, 2, 8}, {10, 25, 18}, {99, 100, 101}};
    int expected[] = {20, 1, 8, 25, 101};
    int all_passed = 1;

    for (int i = 0; i < 5; i++) {
        int result = findMax(test_cases[i][0], test_cases[i][1], test_cases[i][2]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %d, %d, %d. Expected: %d, Got: %d\n", test_cases[i][0], test_cases[i][1], test_cases[i][2], expected[i], result);
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