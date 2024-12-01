public class Main {

    // User function logic goes here
    public static String checkAllPalindromes(String[] strs) {
    for (String str : strs) {
        String reversed = new StringBuilder(str).reverse().toString();
        if (!str.equals(reversed)) {
            return "False";
        }
    }
    return "True";
}

    // Debugging execution of user function
    public static void main(String[] args) {
    // Test case 1
    String[] testCase1 = {"madam", "racecar", "level"};
    System.out.println(checkAllPalindromes(testCase1));  // Expected output: True
    // Test case 2
    String[] testCase2 = {"hello", "racecar", "level"};
    System.out.println(checkAllPalindromes(testCase2));  // Expected output: False
    // Test case 3
    String[] testCase3 = {"madam", "level"};
    System.out.println(checkAllPalindromes(testCase3));  // Expected output: True
    // Test case 4
    String[] testCase4 = {"racecar", "world"};
    System.out.println(checkAllPalindromes(testCase4));  // Expected output: False
}
}