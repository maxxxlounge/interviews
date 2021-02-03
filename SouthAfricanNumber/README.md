# South African Mobile Numbers

Please find an attached CSV file "South African Mobile Numbers". This file contains both
correctly and incorrectly formed South African Mobile Numbers.
Using any Open Source Framework and making use of Object Oriented Coding design
principles.

* Consume the provided file via any of the following means eg. upload / console call / API.
* Test each number and check for correctness, attempt to correct incorrectly formed numbers
and reject numbers that are invalid. (27831234567 is the correct format for this exercise).
* Store the results appropriately to DB / Temporary File - as per your discretion.

Display results by the following means - divide the states as follows:
a. Display acceptable numbers
b. Corrected numbers + what was modified
c. Incorrect numbers.

* Use the following methods to display the data by whichever means you feel make best sense
for a user to interpret, or machine in the case of an API.
a. HTML
b. File (single file or divided - your choice)
c. API endpoint

* Include a user form to test a single number and get a response with a success status/error
message - if incorrect, why? And if self corrected, what was done.

* Application should include basic Documentation (readme) & Unit Tests. Making use of existing
libraries for the heavy lifting is permitted. If you do use a database, please include seeding files.

Conclusion:
Be creative and have loads of fun, project shouldn't take much longer than a couple of hours,
we are not looking for perfection, just an insight into how you work.
-

## RUN PROJECT

```sh

```

## DEPLOYMENT STRATEGY

The input file is intended as first row headers, and comma separated values
The correct format is assumed formed from : fixed prefix "2783" + 7 digits (0-9)

the prefix is called Prefix
the 7 digit after prefix is called the "core" part

### Configuration:

| param | default | description |
|---|---|---|
|-p | "80" | listen port |
|-i | "input.csv" | input source file |
|-d | "output.json" | destination file (json) |

eg:

```sh
#with default parameters values
go run main.go

#with custom parameters values
go run main.go -d=pippo.json -i=other_input.csv -p=8899
```

### Description:

* Check the input data quality and validation
 - file can be opened
 - file has at less 2 column comma separated
 - indexes are unique

* Ingest the file and split the case in different groups (map with id index):
    * numbers valid at first check (regex check)
    * numbers with critical error, intended as UnFixable because some information are missing
    * numbers that can be fixed

Data format is map of reference and each record is :

| key | type |
|---|---|
| key | string | give index from file (it could be a unsigned big int if SQL db used) |
| Original | string | original imput number |
| Changed | string | the result of changes applied |
| Type | custom (like enum) | decription a string like field to control the 3 possible state (ValidFirstAttempt,InvalidCritical,InvalidButFixable)  |
| Errors | []string | string array with description about errors occurred |

An example of a record
```json
"103425772": {
        "Original": "27832719392",
        "Changed": "",
        "Type": "ValidFirstAttempt",
        "Errors":["errors1", "errors2"]
    }
```

The chosen format to best represent the data for a user to interpret and API endpoint is JSON

* Data it is stored in memory and in "output.json" for other uses,
so, it doesn't need DB, 'cause specifics don't ask for resource modification, so no store different from input file needed

Endpoints:

    ```
    GET localhost:8888/numbers  return all loaded numbers (NB: it doesn't used rest because is a search and resource returned is the same (numbers))
    GET localhost:8888/numbers?type={$type}   the values of type field is (valid,fixable,critical)

    GET localhost:8888/numbers?type=valid
    GET localhost:8888/numbers?type=critical
    GET localhost:8888/numbers?type=fixable

    GET localhost:8888/numbers/check?number=
    eg:     http://localhost:8888/numbers/check?number=123123
    The number check endpoint, if empty string passed return a 400 "missing number" error
    It returns always 400 in case of InvalidCritical type number and InvalidButFixable because it means a "Bad request"
    it return 200 for ValidFirstAttempt number type

    GET localhost:8888/check.html main check number page

    ```

### Correction strategy

minimum fixable number is consider having equal or more than "core" part (7) numberic digit
numbers with numeric part less than 7 are considered critical and un-fixable

| inserted | changes | description |
|---|---|---|
| __some text_27831234567 | 27831234567 | adding prefix to complete number, applied to all input |
| 1234567 | 27831234567 | adding prefix to complete number |
| 121234567 | 27121234567 | adding prefix missing digit to complete number (to separate error of wrong and partial prefix)|
| 27811234567 | 27831234567 | replace wrong perfix |
| 278112345672 | 27811234567 | cut exceed digits |
