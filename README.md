# Three most typical Go concurrency pattern

### Runner
The purpose of the runner package is to show how channels can be used to moni- tor the amount of time a program is running and terminate the program if it runs too long. This pattern is useful when developing a program that will be scheduled to run as a background task process. This could be a program that runs as a cron job, or in a worker-based cloud environment like Iron.io.

1. Pool
2. Work