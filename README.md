# Welcome
## Project struct
- cmd
- - main - endpoint
- internal
- - table - the main logic is here
- - validation - logic to validation of the main points
- tests
- - csv - csv files for tests
- - [csv_test.go] - the file itself with unit tests

## To run tests, write to the console
````
go test tests/csv_test.go -v
````
## To run program, write to the console
````
go run cmd/main/main.go

you can provide arguments, where -file test.csv
go run cmd/main/main.go -file test.csv
````
## You can compile a binary file using this command
````
go build csv_solver/cmd/main 
````
## You can run the program using a binary file
````
main.exe
````

#### Further instructions will be displayed in the console after program is launched


### Example:
    test.csv
    ````
    ,A,B,Cell
    1,1,0,1
    2,2,=A1+Cell30,0
    30,0,=B1+A1,5
    ````
go run cmd/main/main.go

Write in console test.csv

Output 
    
    ,A,B,Cell
    1,1,0,1
    2,2,6,0,
    30,0,1,5 
