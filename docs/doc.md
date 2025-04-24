# CapybaraClub

## Navigation:

<details>
  <summary>Navigation</summary>
  <ol>
    <li> <a href="#info"> About project </a> </li>
    <li> <a href="#user-service"> About user-microservice </a> 
        <ul> 
            <li> <a href="#user-service-configyaml"> About configuration file in user-service </a> </li>
            <li> <a href="#routes-for-user-microservice"> About routes for user-microservice</a> </li>
        </ul>
    </li>
    
</details>

### Info
* **Current version:** 0.1 (BETA)
* **Tested on:** ubuntu-latest 
* **Author:** wnderbin
* **Go version:** 1.22.2

## User-service
The task of this microservice is to work with the user.
### User-microservice structure:
* **config** - Working with and storing the configuration. Here you can change the microservice settings in the configuration file *config.yaml*
* **database** - Initialization of microservice operation with redis & postgres databases
* **handlers** - Handlers for urls
* **logger** - Logging tool
* **migrations** - Database migrations
* **migrator** - Performs database migrations
* **models** - Defining a user and their fields as an object
* **routes** - URL's
* **ui** - User interface
* **utils** - Utilities required for user security and authorization
#### The user microservice can only be launched if the startup-status status is changed to 1 in the configuration file. 
#### If 0 is specified, only the logger will be initialized, this status is necessary to run tests for the microservice.

### User-service config.yaml
```
env: "local" # "local"/"dev"/"prod"
startup-status: 0 # 1/0. 1 - run. 0 - tests
address: 0.0.0.0 # the address where the microservice will be launched
service-port: 8081 # the port on which the microservice will be launched
jwt_key: 9BP4natUsI2miUcYVt2go9VFMM3Ayca0K8YN1F5tI0A= # The JWT key is best stored in a location that is not accessible to third-party users, such as a config file that won't be included in the repository via .gitignore, but since this project is a simple microservices + practices prototype, I don't think it will make any difference if I put it here.

postgres:
  host: "localhost"
  port: 5432
  user: "wnd"
  password: "123"
  dbname: "wnd"
  sslmode: "disable"

redis:
  address: "localhost:6379"
  password: ""
  db: 0
```
**[ENV]** - Affects the format and information in messages that will be sent by the logger.
* **local** - Text/LevelDebug
* **debug** - JSON/LevelDebug
* **prod** - JSON/LevelInfo

**[JWT-KEY]** - Required for user authentication. When registering, a user is created, and when logging into an account, this key is assigned to him for 5 minutes, after which time the key is no longer relevant. The key itself is stored in cookies.

**If you want to change the key for security purposes, you can generate it and assign it to the jwt_key variable in the configuration file.**
```
openssl rand -base64 32
```

### Routes for user microservice
#### GET
* **/main** - Home page
* **/register/** - Registration form
* **/login/** - Account login form
* **/user** - Getting authorized user data
#### POST
* **/register/postform** - Sends data from the registration form
* **/login/postform** - Sends data from the login form
#### PUT
* **/user** - Changes user data. Path: /user?name=X&username=X&email=X&password=X \
**X** - your new data 
#### DELETE
* **/user** - Deletes the current user

