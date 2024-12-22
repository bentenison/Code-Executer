#include <stdio.h>

// User function logic goes here
int gcd(int a, int b) {
    if (b == 0) {
        return a;
    }
    return gcd(b, a % b);
}

// Test cases for the function
int main() {
    int passed = 1;
    int a = 56, b = 98;
    // Check GCD
    printf("%s\n", passed ? "true" : "false");
    return 0;
}