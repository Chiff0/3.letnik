REDOVALNICA - DOCUMENTATION
INSTALLATION
Clone the repository:
git clone https://github.com/yourusername/student-module.git
cd student-module
Download dependencies:
go mod download
Build the application:
go build -o redovalnica ./cmd
BASIC USAGE
Run the application with default settings:
./redovalnica
Show help:
./redovalnica --help
FLAGS
The application supports three flags:
--stOcen : Minimum number of grades required for a positive grade (default: 1)
--minOcena : Minimum possible grade (default: 6)
--maxOcena : Maximum possible grade (default: 10)
Example:
./redovalnica --stOcen=3 --minOcena=5 --maxOcena=10
COMMANDS

izpis - Display all students

Shows all students and their grades.
./redovalnica izpis
Example output:
REDOVALNICA
63220059: Filip Dobnikar - Grades: [10 10 10 10]
65454: Rok Dobnikar - Grades: [7 1 2 6]
1111111: Burek Cmurek - Grades: [9 8 9 9]

uspeh - Display final success

Shows average grades and final success for all students.
./redovalnica uspeh
Example with settings:
./redovalnica --stOcen=2 --minOcena=6 uspeh
Example output:
FINAL SUCCESS
Filip Dobnikar: average grade 10.00 -> EXCELLENT STUDENT!!!!!
Rok Dobnikar: average grade 4.00 -> BAD student >:
Burek Cmurek: average grade 8.75 -> Average student :|

dodaj - Add a grade to a student

Adds a new grade to a specific student.
./redovalnica dodaj <student_id> <grade>
Examples:
Add grade 9 to student Filip:
./redovalnica dodaj 63220059 9
Add grade with different limits:
./redovalnica --minOcena=1 --maxOcena=5 dodaj 63220059 4
USAGE EXAMPLES
Example 1: Check success with minimum number of grades
Student must have at least 3 grades:
./redovalnica --stOcen=3 uspeh
Example 2: Different grade range
Grades from 1 to 5:
./redovalnica --minOcena=1 --maxOcena=5 uspeh
Example 3: Combining commands
Display students:
./redovalnica izpis
Add a grade:
./redovalnica dodaj 63220059 8
Check success:
./redovalnica --stOcen=2 uspeh
PROJECT STRUCTURE
student-module/
go.mod              - Go module definition
go.sum              - Dependencies
student.go          - Main module logic
cmd/
main.go           - CLI application
GRADING SYSTEM
Grade >= 9: EXCELLENT STUDENT!!!!!
Grade >= minOcena: Average student :|
Grade < minOcena: BAD student >:
Not enough grades: Insufficient number of grades
