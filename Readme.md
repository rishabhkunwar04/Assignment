## Requirement

You are required to build the delivery service endpoint which will serve a list of campaigns.
1. This will be a GET endpoint available on a web server
2. You will have to keep an eye for performance issues



3. Think about scalability (horizontal as well as vertical)
4. Write test cases to check for correctness
5. Your architecture should comply with the read-heavy nature of this problem. The number
of campaigns will always be much smaller than the number of delivery requests. 1000s
of campaigns vs Billions of delivery request
6. Think about how your code complexity will increase if the number of dimensions will
increase
7. You are free to choose any database or cache of your choice. There is no constraint on
SQL/NoSQL DBs
8. You are not required to create end-points for storing Targeting rules and Campaign
details. Those can be directly inserted into the database. But your delivery service
should react to the changes in the DB. For example, if the status of a campaign is
changed to “inactive”, the delivery service should stop serving it. Similarly, if the link of a
campaign is changed, the delivery service should send the updated link