user attempt logs and number of attempts must be saved

- hint user after 3 consicutive failures with the correct answer add penulty to overall score 10 points
- user will have 3 attempts per question in a day per challenge
- user can change the question 3 times per challenge this will give penulty and hit to overall score
ES is used for analytics it holds challenge performance per challenge per user per language
user performnace per question and language
Code execution stats per user per question per execution


indexes for these must be created 
kafka topics must also be created
in the broker the call happens 
in the admin the consumer consumrs the data

we will have some analytics running from the ES

17*3 = 51  17 challenges 3 questions per challlenge
17*30 = 510
500 points for reaching next rank 500/30 points  per challnege easy

700/60  12 challenges
12*3 36 questions medium level
20 points each

1200/90  14 challenges
14*3 42 questions 
30 points each

- on server start fetch the users global performace and metrics into redis  if needed update redis if rank changes or score changes
needed for accessing previous rank

- seperate API from broker which internally calls GRPC of diffrent services
mark question complete with is completed true
add score to challenge
add total score 
add penulty points if needed
calculate the speed and accuracy based on the number of answers and submissions
while challenge creation add the totalscore of the challenge by rank





- if hint used there will bge penulty
hint can be pop out after 3 consicutive submission failure per question
the score will be half for the hint used
keep track of questionids and failed attempts on frontend if limit exceeds give popup with instruction to copy and paste the code 
after the submission mark hint used true in the payload
based on this the feature will decide what to do

exponential backoff till the dbs and ES not started



1. Multiple Choice Questions (MCQs)

    Present a code snippet or concept and ask users to select the correct answer from a list of options.

2. True/False Questions

    Pose a statement related to programming concepts, and users must decide if it's true or false.

3. Code Output Questions

    Provide a code snippet and ask what the output will be when the code is executed.

4. Code Correction Questions

    Present a piece of code with one or more errors, and ask users to identify and correct the mistakes.

5. Fill-in-the-Blank (as you have)

    Provide code snippets with blanks for users to fill in the correct keywords, variable names, or methods.

6. Short Answer Questions

    Ask for explanations of concepts, such as "What is polymorphism?" or "Explain the difference between a stack and a queue."

7. Debugging Questions

    Provide code that contains bugs and ask users to debug it, identifying issues and proposing fixes.

8. Code Implementation Questions

    Ask users to implement a specific algorithm or data structure from scratch.

9. Code Refactoring Questions

    Provide a piece of code and ask users to refactor it for better readability, performance, or maintainability.

10. Pair Programming Scenarios

    Present a coding challenge that requires collaborative input, asking users to select the best approach or solution collaboratively.

11. Scenario-based Questions

    Describe a real-world problem and ask how to approach solving it using programming concepts.

12. Algorithm Complexity Questions

    Ask users to analyze the time and space complexity of given algorithms or code snippets.

13. Data Structure Operations

    Provide a scenario involving data structures (like stacks, queues, trees) and ask how to perform specific operations.

14. Version Control Questions

    Ask about best practices and commands related to Git or other version control systems.

15. Framework/Library-Specific Questions

    Focus on specific frameworks (like React, Django) or libraries (like NumPy, jQuery) and ask about their features or functionalities.

These question types can engage users in different ways and assess a variety of skills and knowledge levels in programming.