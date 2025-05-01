# CapybaraClub

## Navigation:

<details>
  <summary>Navigation</summary>
  <ol>
    <li> <a href="#info"> About project </a> </li>
    <li> <a href="#user-service"> About user-microservice </a> 
        <ul> 
            <li> <a href="#user-service-configyaml"> About configuration file in user-service </a> </li>
        </ul>
    </li>
    <li><a href="#admin-service"> About admin-microservice </a>
      <ul> 
            <li> <a href="#admin-service-configyaml"> About configuration file in admin-service </a> </li>
        </ul>
    </li>
    <li><a href="#restaurant-service">About restaurant-microservice</a></li>
    <li><a href="#order-service">About order-microservice</a></li>
    
</details>

### Info
* **Current version:** 0.8 (BETA)
* **Tested on:** ubuntu-latest 
* **Author:** wnderbin
* **Go version:** 1.24.2

## User-service
The task of this microservice is to work with the user.
### User-microservice structure:
* **config** - microservice settings, specified in the environment variable at startup
* **handlers** - requests
* **nats_client** - message broker and its methods
* **routes** - urls
* **ui** - user interface

### User-service config.yaml
```
address: 0.0.0.0 # the address where the microservice will be launched
service-port: 8081 # the port on which the microservice will be launched
jwt_key: 9BP4natUsI2miUcYVt2go9VFMM3Ayca0K8YN1F5tI0A= 
```

**[JWT-KEY]** - Required for user authentication. When registering, a user is created, and when logging into an account, this key is assigned to him for 5 minutes, after which time the key is no longer relevant. The key itself is stored in cookies.

**If you want to change the key for security purposes, you can generate it and assign it to the jwt_key variable in the configuration file.**
```
openssl rand -base64 32
```

-----

## Admin-service

The task of this microservice is to work with administrators

**Administrators** - users with advanced parameters. At the same time, the user cannot be given administrator rights, you can only give him an administrator account. Login to all administrator privileges is carried out using the login. Otherwise, the use of the administrator account will be prohibited. Login lasts the same time as the user - 5 minutes.

### Admin-microservice structure:

* **config** - microservice settings, specified in the environment variable at startup
* **handlers** - requests
* **nats_client** - message broker and its methods
* **routes** - urls
* **ui** - user interface


### Admin-service config.yaml

```
address: 0.0.0.0 # the address where the microservice will be launched
service-port: 8083 # the port on which the microservice will be launched
admin_name: admin
admin_password: pass123
admin_email: admin@mail
jwt_key: XU5AusKEA2MCVt5khTUsVvwHj90kBkLNyqqUCAZRixU= 
```

**The administrator specified in this configuration file is the root and cannot be deleted or changed. It is initialized together with the database.** 

**Administrator:**
* **admin_name:** admin
* **admin_password:** pass123
* **admin_email:** admin@mail

**[JWT-KEY]** - Required for admin authentication.

### Restaurant-service
The purpose of this microservice is to show the user a selection of restaurants.

### Order-service
The purpose of this microservice is to show active orders to the user.