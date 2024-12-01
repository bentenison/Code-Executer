#include <stdio.h>
#include <ctype.h>

// User function logic goes here
int countVowels(char str[]) {
    int count = 0;
    for (int i = 0; str[i] != '\0'; i++) {
        char ch = tolower(str[i]);
        if (ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u') {
            count++;
        }
    }
    return count;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"Hello World", "abcdefg", "AEIOU", "xyz"};
    int expected[] = {3, 2, 5, 0};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int result = countVowels(test_cases[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %d, Got: %d\n", test_cases[i], expected[i], result);
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