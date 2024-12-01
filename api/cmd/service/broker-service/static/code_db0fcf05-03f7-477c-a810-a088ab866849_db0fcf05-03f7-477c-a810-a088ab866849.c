#include <stdio.h>

// User function logic goes here
int isLeapYear(int year) {
    if ((year % 400 == 0) || (year % 4 == 0 && year % 100 != 0)) {
        return 1;
    }
    return 0;
}

// Test cases for the function
int main() {
    int test_cases[] = {2020, 1900, 2000, 2024, 2100};
    int expected[] = {1, 0, 1, 1, 0};
    int all_passed = 1;

    for (int i = 0; i < 5; i++) {
        int result = isLeapYear(test_cases[i]);
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