# Assumptions
### Data
This app makes the assumption that POSTing the same report will end up in multiple records being created.
This app assumes that reports are not needed to be retrieved in the future. 
This assumption was made from the requirement of the report endpoint being POST.

It will make sense to seek more clarification on this from the business. If we are to flesh out the implementation, we would need to asses the necessity of the report being persisted. We should also understand the context around the usage of this API and what it means.
Depending on how the business would need to use these resources, there are several design changes we could make. 

### Request-Response
We have some strict requirements around the requests that will be posted to the server. Those seem to be followed based on the knowledge we have. We took some liberties (shortcuts) with the implementation of the Response. Since we are using the same models for the DB and request and response, we have some overlap. That was assumed to be acceptable for this implementation. In the future we would want to separate these into different structs and possibly add a new service layer to maintain better separation of concern and testability. 

# Comments
### Testing
The current implementation lacking in tests. Here is an outline of the test we would need to be production ready.
- Unit tests for all modules, with proper mocking of dependencies. Perhaps using this -> https://github.com/DATA-DOG/go-sqlmock
- Functional/ Integration tests. Test the app with the database and verify. It would make sense to build this as a separate module or cmd app. 

### Code 
1. Typo: joynet should be joyent. Great job, me, for messing **that** up!
1. I tried to move everything into and interface and code around that. Ideally this would provide us some flexibility and separation of concerns. This should make testing much easier.
1. Adding the Env: this was created to answer the question "How can the controllers share a DB connection". Perhaps there is a cleaner way to do this, but this should work for now. The db interface should be thread safe and the connection should be able to scale concurrent requests. 
1. I would love to get a basic GHAs Workflow running for this, time permitting. I love this project for linting in CI https://golangci-lint.run/

### Path to production
- CICD here are a list of possible iterations we could have to get everything automated :D 
    1. Start with simple lint->build->test on push
    1. Add container registry and refine tagging strategy 
    1. Setup and run containers for tests in ephemeral environments for Pull Requests
    1. Run on `main` branch on push to deploy to staging environment 
    1. Setup Automated Deploys/ Rollbacks with manual approval gates
    1. Refine deployment strategies and remove manual gates
- Environment
    1. To run in kubernetes we will need to generate the appropriate K8s objects
    1. Should we run the DB in with our app or use a service?
    1. We would need to ensure that `GIN_MODE` is set to `release`
