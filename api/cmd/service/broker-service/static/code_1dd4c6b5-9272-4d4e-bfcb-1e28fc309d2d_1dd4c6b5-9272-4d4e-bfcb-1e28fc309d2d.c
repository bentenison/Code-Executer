#include <stdio.h>
#include <ctype.h>

// User function logic goes here
int containsSpecialCharacter(char str[]) {
    for (int i = 0; str[i] != '\0'; i++) {
        if (!isalnum(str[i])) {
            return 1;
        }
    }
    return 0;
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"Hello", "Hello@World", "abc123", "!@#"};
    int expected[] = {0, 1, 0, 1};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int result = containsSpecialCharacter(test_cases[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %d, Got: %d\n", test_cases[i], expected[i], result);
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