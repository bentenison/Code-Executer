#include <stdio.h>
#include <string.h>

// User function logic goes here
void evenIndexCharacters(char str[]) {
    for (int i = 0; str[i] != '\0'; i += 2) {
        printf("%c", str[i]);
    }
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"abcdef", "abcd", "hello", "world"};
    char expected[][50] = {"ace", "ac", "hlo", "wr"};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        char result[50];
        evenIndexCharacters(test_cases[i], result);
        if (strcmp(result, expected[i]) != 0) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %s, Got: %s\n", test_cases[i], expected[i], result);
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