#include <stdio.h>
#include <string.h>

// User function logic goes here
void rotateString(char str[], int n) {
    int len = strlen(str);
    n = n % len;  // To handle cases where N is larger than the string length
    if (n == 0) return;
    char temp[len + 1];
    strcpy(temp, str + len - n);  // Copy last N characters
    strncat(temp, str, len - n);  // Append first (len - N) characters
    strcpy(str, temp);
}

// Test cases for the function
int main() {
    char test_cases[][50] = {"abcdef", "hello", "world", "abc"};
    int rotations[] = {2, 3, 1, 1};
    char expected[][50] = {"efabcd", "lohel", "dworl", "cab"};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        rotateString(test_cases[i], rotations[i]);
        if (strcmp(test_cases[i], expected[i]) != 0) {
            all_passed = 0;
            printf("Failed for Input: %s, N: %d. Expected: %s, Got: %s\n", test_cases[i], rotations[i], expected[i], test_cases[i]);
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