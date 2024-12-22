#include <stdio.h>
#include <string.h>

// User function logic goes here
void sortString(char str[]) {
    int len = strlen(str);
    for (int i = 0; i < len - 1; i++) {
        for (int j = i + 1; j < len; j++) {
            if (str[i] > str[j]) {
                char temp = str[i];
                str[i] = str[j];
                str[j] = temp;
            }
        }
    }
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"golang", "python", "abcdef"};
    char expected[][50] = {"aglno", "hnopty", "abcdef"};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        sortString(test_cases[i]);
        if (strcmp(test_cases[i], expected[i]) != 0) {
            all_passed = 0;
            printf("Failed for Input: %s. Expected: %s, Got: %s\n", test_cases[i], expected[i], test_cases[i]);
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