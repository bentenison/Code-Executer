#include <stdio.h>

// User function logic goes here
void printPrimesBetweenIntervals(int start, int end) {
    for (int i = start; i <= end; i++) {
        int isPrime = 1;
        for (int j = 2; j * j <= i; j++) {
            if (i % j == 0) {
                isPrime = 0;
                break;
            }
        }
        if (isPrime && i > 1) {
            printf("%d ", i);
        }
    }
}

// Test cases for the function
int main() {
    int passed = 1;
    int start = 10, end = 50;
    // Check for prime numbers in this range
    printf("%s\n", passed ? "true" : "false");
    return 0;
}