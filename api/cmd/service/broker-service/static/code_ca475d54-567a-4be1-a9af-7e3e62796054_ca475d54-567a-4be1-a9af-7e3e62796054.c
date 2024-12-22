#include <stdio.h>
#include <string.h>

// User function logic goes here
int isRotation(char str1[], char str2[]) {
    int len1 = strlen(str1), len2 = strlen(str2);
    if (len1 != len2) return 0; // Strings must have the same length
    char temp[2*len1 + 1];
    strcpy(temp, str1);
    strcat(temp, str1); // Concatenate str1 with itself
    if (strstr(temp, str2) != NULL) {
        return 1; // str2 is a rotation of str1
    }
    return 0; // str2 is not a rotation of str1
}

// Test cases for the function
int main() {
    char str1[][50] = {"hello", "abc", "abcd", "abcdef"};
    char str2[][50] = {"lohel", "bca", "cdab", "fabcde"};
    int expected[] = {1, 1, 1, 0};
    int all_passed = 1;

    for (int i = 0; i < 4; i++) {
        int result = isRotation(str1[i], str2[i]);
        if (result != expected[i]) {
            all_passed = 0;
            printf("Failed for Input: %s and %s. Expected: %d, Got: %d\n", str1[i], str2[i], expected[i], result);
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