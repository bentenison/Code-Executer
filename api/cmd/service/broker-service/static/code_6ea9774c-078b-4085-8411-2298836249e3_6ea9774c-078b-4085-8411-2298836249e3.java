public class Main {

    // User function logic goes here
    public static String checkOddEven(int num) {
    if (num % 2 == 0) {
        return "Even";
    } else {
        return "Odd";
    }
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    int[][] testCases = {
        {5, "Odd"},  // Test case 1: 5 is Odd
        {8, "Even"}  // Test case 2: 8 is Even
    };
    for (int[] testCase : testCases) {
        int input = testCase[0];
        String expectedOutput = testCase[1];
        String output = checkOddEven(input);
        if (!output.equals(expectedOutput)) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}