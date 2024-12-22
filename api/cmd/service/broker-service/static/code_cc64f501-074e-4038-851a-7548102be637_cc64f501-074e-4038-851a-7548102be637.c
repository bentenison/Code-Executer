#include <stdio.h>

// User function logic goes here
int largestPrimeFactor(int n) {
    int largest_factor = -1;
    while (n % 2 == 0) {
        largest_factor = 2;
        n /= 2;
    }
    for (int i = 3; i * i <= n; i += 2) {
        while (n % i == 0) {
            largest_factor = i;
            n /= i;
        }
    }
    if (n > 2) largest_factor = n;
    return largest_factor;
}

// Test cases for the function
int main() {
    int test_cases[] = {56, 13195, 29, 49, 77};
    int expected[] = {7, 29, 29, 7, 11};
    int all_passed = 1;

    for (int i = 0; i < 5; i++) {
        int result = largestPrimeFactor(test_cases[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %d. Expected: %d, Got: %d\n", test_cases[i], expected[i], result);
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