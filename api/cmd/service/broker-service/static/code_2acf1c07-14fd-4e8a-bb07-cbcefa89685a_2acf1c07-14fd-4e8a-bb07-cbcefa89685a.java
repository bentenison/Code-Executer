public class Main {

    // User function logic goes here
    public static int findSmallest(int[] arr) {
    int smallest = arr[0];
    for (int i = 1; i < arr.length; i++) {
        if (arr[i] < smallest) {
            smallest = arr[i];
        }
    }
    return smallest;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    int[][] testCases = {
        {3, 5, 1, 9, 2},  // Test case 1: Smallest number is 1
        {10, 20, 30, 5}   // Test case 2: Smallest number is 5
    };
    for (int[] testCase : testCases) {
        int input[] = testCase;
        int expectedOutput = testCase[testCase.length - 1];
        int output = findSmallest(input);
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}