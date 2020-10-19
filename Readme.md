The repo is to understand how GIT workw along with basic appreciation of REST API.

To reproduce the training

1. take the code with TAG : TRAIN_BASELINE
2. create a new branch - say training1
3. Checkin this code

Tasks to do:

Volunteer 1:
create a new brach training1_v2a
Copy the code from main_v2a.go into main.go
Commit and push

Volunteer 2:
create a new brach training1_v2b
Copy the code from main_v2b.go into main.go
Commit and push

Volunteer 3:
create a new brach training1_v2c
Copy the code from main_v2c.go into main.go
Commit and push

Volunteer 1:
Raise a pull request into training1 and self approve

Volunteer 2:
Raise a pull request into training1 and ask for approval

Volunteer 3:
Raise a pull request into training1 - you will see conflicts
Merge training1 branch with branch training1_v2c
Resolve conflicts
Push
Raise pull request and ask for approval
