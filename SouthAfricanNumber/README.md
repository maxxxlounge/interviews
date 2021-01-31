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


## DEPLOYMENT STRATEGY

The correct format is assumed formed from : fixed prefix "2783" + 7 digits (0-9)


Strategy:

* Check the input data quality and validation
 - file can be opened
 - file has at less 2 column comma separated
 - indexes are unique

* Ingest the file and split the case in different groups (map with id index):
    * numbers valid at first check (regex check)
    * numbers with critical error, intended as UnFixable because some information are missing
    * numbers that can be fixed

