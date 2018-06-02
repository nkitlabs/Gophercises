# Exercise #1: Quiz Game

you can follow the details and instruction more on the [link](https://gophercises.com/exercises/quiz) 

## Exercise details

Develop a program that simulate the quiz to a tester. The quiz is started after pressing the enter. The question will be shown one by one provided from given CSV file (more details below). The quiz will end after a tester answers all questions or the test meets the time limit (the default time limit is 30 seconds). At the end, the program will show the result how many questions correct from the total number of questions in the quiz.

The CSV file should default to `problems.csv` (example shown below), but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.

**Big Credit : (https://gophercises.com/exercises/)**
