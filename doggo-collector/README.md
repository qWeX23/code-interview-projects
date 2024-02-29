# User Login Code Challenge

Using a JavaScript framework of your choice (preferably [React](https://reactjs.org/)), create a simple login screen that allows users to enter their username and password and submit the login form to a backend process.

Create a backend (preferably using [GoLang](https://go.dev/), but not required) that processes the login information and checks if the username and password are valid. If the login information is valid, the backend should return a success message to the user. If the login information is invalid, the backend should return an error message to the user.


## Instructions
1. Click "Use this template" to create a copy of this repository in your personal github account.  
1. Update the README in your new repo with:
    * a `How-To` section containing any instructions needed to execute your program.
    * an `Assumptions` section containing documentation on any assumptions made while interpreting the requirements.
1. Send an email to Scoir (code_challenge@scoir.com) with a link to your newly created repo containing the completed exercise (preferably no later than one day before your next interview).

## Expectations
1. This exercise is meant to drive a conversation between you and Scoir's hiring team.  
1. Please invest only enough time needed to demonstrate your approach to problem solving and code design.  
1. Within reason, treat your solution as if it would become a production system.
1. If you have any questions, feel free to contact us at code_challenge@scoir.com

# How-To

## Prerequisites 
- Docker Installed
- NodeJS + NPM installed
- MongoSh installed

## Setup Locally 
1. clone repo this repo
    - `git clone git@github.com:qWeX23/doggo-collector.git`
1. Run the backend services
    - `docker compose up -d --build`
1. Add a user to the database 
    - login to the database `mongosh -u root -p example`    
    - you should now be using the mongo shell
    - create the DC-App Db `use DC-App`
    - insert user `db.Users.insert({username:"user123",password:"password123"})`
1. Run the frontend 
    - `cd dc-app`
    - run the app `npm start`
1. You should automatically be redirected to `localhost:3000`


# Assumptions

## UI 
I took some liberties with the UI. I named the title of the app "Doggo Collector", instead of "Dog Catcher". The former seems a bit more jovial and light hearted. 

This was my first attempt into using a real, modern front end framework like React.  The learning curve was not as intense as i thought it would be. The styling of the app is non existent , and it shows. I certainly would hope to improve this in the future and would want to lean on colleagues that are more in tune with the technical and business sides here. "I am a back-end guy" is not an excuse for the way the app looks, but i hope it is acceptable for demonstrating the tech working together.

I am mostly satisfied with the organization of the app. It could use some rearranging into folders if it continues to grow. If i were to change anything, I would make a new component called card list that dealt with the operations and organizations of the Cards, and put that in the dashboard, to help reduce the size of the current dashboard. 

This app needs a lot of TLC to be production ready, but maybe the biggest area of concern is to add tests. We will want to write tests that mock API output and verify the correct text shows on the screen. 

Lastly, we will need to move away from hardcoded strings for the api routes. I think we would want to set up some sort of proxy on our end to communicate with the API without having to know the URL.

## API-App

I decided to use Gin frame work for this app, it seems to be a solid library for making these kinds of APIs. I wanted to first make an app that "worked" for what we needed, and then focus on the internals. This included a RESTful HTTP interface for the front end and a Mongo DB backend. Adding the dockerfile and docker compose was easy enough, and helped speed up the development process. 

### API

I wanted to follow good restful design as much as I could. I would be open to feedback about the naming of my routes and overall design here but I think it is in a good spot. We would want to add versioning from the beginning, instead of when we need to publish a new version. This design covers the need for what we need now and should allow for expansion in the future. 

#### Code organization

I assumed that it would be ok to deviate from the best practices of the actual code organization. For example we have lots of logic around the operations of the card, we would want to _not_ put it in the handler function where it is now. We should put it into its own interface and create a struct that implements this for mocking (via the "magic of dependency injection"). I have some of this code stubbed out that shows this, but i did not implement it fully in the interest of time. 

Furthermore, I would want to begin separating the code out int separate modules, instead of cramming everything into main. I think a good start would be to make a module for each "entity" ( in this case card and dog. ) The module would include the service, data models, and database integrations for the object. The main function would call the new API, which would organize all of the routes and controllers. The API module would also handle all of the dependencies for all the different entities. 

#### Database

This is my first use of mongo DB outside of college 8 years ago (I used it briefly for my capstone project, a fun ed-tech app!). I love how easy it was to set up and get going. I do not think I am in love with my decisions to organize the databases, but let me explain. 

The main Database is called DC-App. The Users collection in this database holds all of the users data. for now this is username, password and token. Lets overlook the obvious security risk here for now. This collection is used by the login route and the auth middleware to validate logins. 

The secondary database is called Cards. It holds the cards data for each user. Each collection is named after the Users Object ID. the collection then has all of the cards the user currently has in their collection. Deleted cards are removed forever. I used this data model to facilitate what i think to be an efficient data access. A normal person may have just added `CurrentCards` as an array attribute to the main users collection, but whats the fun if you can't optimize for a problem that does not exist yet. (haha)

#### DevSecOps and other thoughts

We need to update the app to not allow all origins like it currently does. I think the proxy mentioned earlier would help that. The database needs to be updated and maintained with real service accounts and under no circumstances should we be using root as we are. The app needs unit tests and a functional test suite. I would soon after that refactor is complete I would make a CI pipeline to start getting this automatically tested and into a container registry (would need this for the front end too). DevSecOps is a made up word. I have no other thoughts for now, I hope this is enough to cover the bases for our interview! 

