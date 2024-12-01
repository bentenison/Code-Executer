public class Main {

    // User function logic goes here
    public static int findLargest(int[] arr) {
    int largest = arr[0];
    for (int i = 1; i < arr.length; i++) {
        if (arr[i] > largest) {
            largest = arr[i];
        }
    }
    return largest;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    int[][] testCases = {
        {1, 5, 3, 9, 2},  // Test case 1: Largest number is 9
        {10, 20, 30, 40}  // Test case 2: Largest number is 40
    };
    for (int[] testCase : testCases) {
        int input[] = testCase;
        int expectedOutput = testCase[testCase.length - 1];
        int output = findLargest(input);
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}