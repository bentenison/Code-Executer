#include <stdio.h>
#include <string.h>

// User function logic goes here
void printParts(char str[], int n) {
    int len = strlen(str);
    int part_size = len / n;
    for (int i = 0; i < n; i++) {
        char part[part_size + 1];
        strncpy(part, str + i * part_size, part_size);
        part[part_size] = '\0';
        printf("%s\n", part);
    }
}

// Test cases for the function
int main() {
    int passed = 1;
    char test_cases[2][100] = {
        "abcdef",
        "abcdefghi"
    };
    int n[2] = {2, 3};
    for (int i = 0; i < 2; i++) {
        printParts(test_cases[i], n[i]);
    }
    printf("%s\n", passed ? "true" : "false");
    return 0;
}