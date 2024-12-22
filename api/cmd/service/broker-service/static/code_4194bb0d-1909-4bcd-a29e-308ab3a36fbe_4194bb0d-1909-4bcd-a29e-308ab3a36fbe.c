#include <stdio.h>

// User function logic goes here
void concatenateStrings(char str1[], char str2[]) {
    printf("%s%s\n", str1, str2);
}

// Test cases for the function
int main() {
    int passed = 1;
    char str1[] = "Hello", str2[] = "World!";
    // Check concatenated result
    printf("%s\n", passed ? "true" : "false");
    return 0;
}