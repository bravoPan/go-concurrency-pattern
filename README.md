# Three most typical Go concurrency pattern

### Runner
The purpose of the runner package is to show how channels can be used to moni- tor the amount of time a program is running and terminate the program if it runs too long. This pattern is useful when developing a program that will be scheduled to run as a background task process. This could be a program that runs as a cron job, or in a worker-based cloud environment like Iron.io.

### Pool
The purpose of the pool package is to show how you can use a buffered channel to pool a set of resources that can be shared and individually used by any number of goroutines. This pattern is useful when you have a static set of resources to share, such as database connections or memory buffers. When a goroutine needs one of these resources from the pool, it can acquire the resource, use it, and then return it to the pool.

### Work
The purpose of the work package is to show how you can use an unbuffered channel to create a pool of goroutines that will perform and control the amount of work that gets done concurrently. This is a better approach than using a buffered channel of some arbitrary static size that acts as a queue of work and throwing a bunch of gorou- tines at it. Unbuffered channels provide a guarantee that data has been exchanged between two goroutines. This approach of using an unbuffered channel allows the user to know when the pool is performing the work, and the channel pushes back when it can’t accept any more work because it’s busy. No work is ever lost or stuck in a queue that has no guarantee it will ever be worked on.