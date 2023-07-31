# Objectives
The purpose of this project is for you to calculate the following:
    
    *Average
    *Median
    *Variance
    *Standard Deviation

# Instructions
Your program must be able to read from a file and print the result of each statistic mentioned above. In other words, your program must be able to read the data present in the path passed as argument. The data in the file will be presented as the following example:
```sh
189
113
121
114
145
110
...
```

This data represents a statistical population: each line contains one value.

To run your program a command similar to this one will be used if your project is made in Go:

```sh
go run . data.txt
```

After reading the file, your program must execute each of the calculations asked above and print the results in the following manner (the following numbers are only examples):
```sh
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```