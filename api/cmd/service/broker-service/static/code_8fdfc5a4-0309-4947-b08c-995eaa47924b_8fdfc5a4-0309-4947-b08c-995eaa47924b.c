#include <stdio.h>
#include <math.h>

// User function logic goes here
float calculateCompoundInterest(float principal, float rate, float time, int n) {
    return principal * pow((1 + rate / (100 * n)), n * time) - principal;
}

// Test cases for the function
int main() {
    int passed = 1;
    float principal = 1000, rate = 5, time = 2;
    int n = 4;
    float expected = 104.486;
    if (calculateCompoundInterest(principal, rate, time, n) != expected) {
        passed = 0;
    }
    printf("%s\n", passed ? "true" : "false");
    return 0;
}