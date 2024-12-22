#include <stdio.h>

// User function logic goes here
void reverseArray(int arr[], int size) {
    int left = 0, right = size - 1;
    while (left < right) {
        int temp = arr[left];
        arr[left] = arr[right];
        arr[right] = temp;
        left++;
        right--;
    }
}

// Test cases for the function
int main() {
    int test_cases[][5] = {{1, 2, 3, 4, 5}, {7, 8, 9}, {10, 20}};
    int expected[][5] = {{5, 4, 3, 2, 1}, {9, 8, 7}, {20, 10}};
    int all_passed = 1;

    for (int i = 0; i < 3; i++) {
        reverseArray(test_cases[i], 5);
        for (int j = 0; j < 5; j++) {
            if (test_cases[i][j] != expected[i][j]) {
                all_passed = 0;
                printf("Failed for Input: {%d, %d, %d, %d, %d}. Expected: {%d, %d, %d, %d, %d}, Got: {%d, %d, %d, %d, %d}\n",
                        test_cases[i][0], test_cases[i][1], test_cases[i][2], test_cases[i][3], test_cases[i][4],
                        expected[i][0], expected[i][1], expected[i][2], expected[i][3], expected[i][4],
                        test_cases[i][0], test_cases[i][1], test_cases[i][2], test_cases[i][3], test_cases[i][4]);
            }
        }
    }
    if (all_passed) {
        printf("true\n");
    } else {
        printf("false\n");
    }
    return 0;
}