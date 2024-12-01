public class Main {

    // User function logic goes here
    public static boolean isPalindrome(String str) {
    String reversed = new StringBuilder(str).reverse().toString();
    return str.equals(reversed);
}

    // Debugging execution of user function
    public static void main(String[] args) {
    boolean allPassed = true;
    // Loop over all test cases
    String[][] testCases = {
        {"madam", "true"},  // Test case 1: 'madam' is a palindrome
        {"hello", "false"}    // Test case 2: 'hello' is not a palindrome
    };
    for (String[] testCase : testCases) {
        String input = testCase[0];
        boolean expectedOutput = Boolean.parseBoolean(testCase[1]);
        boolean output = isPalindrome(input);
        if (output != expectedOutput) {
            allPassed = false;
            break;
        }
    }
    System.out.println(allPassed ? "true" : "false");  // Output true if all test cases pass, false otherwise
}
}