public class Main {

    // User function logic goes here
    public static boolean isLeapYear(int year) {
    if ((year % 4 == 0 && year % 100 != 0) || (year % 400 == 0)) {
        return true;
    }
    return false;
}

    // Debugging execution of user function
    public static void main(String[] args) {
    // Call the function with a test case
    System.out.println(isLeapYear(2020));  // Test case 1
}
}