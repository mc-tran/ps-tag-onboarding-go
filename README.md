# ps-tag-onboarding-go

This application is an API based on the Java application that can be found here: https://github.com/wexinc/ps-tag-onboarding

The application is written in Golang, utilising the Gorilla Mux framework. This framework was chosen in prder to simplify the router and handlers.


**Run Application**

docker-compose up


**Hit endpoints**

_Create a User_

curl -v localhost:8080/save -d '{"firstname":"Ham", "lastname":"Cheese", "email": "a@a.a", "age": 20}'

_Find a User_

use the id returned from create user

curl localhost:8080/find/6566a0d8595456f2953f62d1

