#include <stdio.h>
#include <string.h>

// User function logic goes here
void removeChar(char str[], int index) {
    int len = strlen(str);
    for (int i = index; i < len; i++) {
        str[i] = str[i + 1];
    }
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"Hello", "World", "abcdef", "abc"};
    int indices[] = {1, 2, 3, 0};
    char expected[][50] = {"Hllo", "Wold", "abcef", "bc"};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        removeChar(test_cases[i], indices[i]);
        if (strcmp(test_cases[i], expected[i]) != 0) {
            all_passed = 0;
            printf("Failed for Input: %s, Index: %d. Expected: %s, Got: %s\n", test_cases[i], indices[i], expected[i], test_cases[i]);
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