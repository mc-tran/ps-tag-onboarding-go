# ps-tag-onboarding-go

This application is an API based on the Java application that can be found here: https://github.com/wexinc/ps-tag-onboarding

The application is written in Golang, utilising the Gorilla Mux framework. This framework was chosen in prder to simplify the router and handlers.


**Run Application**

docker-compose up


**Hit endpoints**

***Create a User*** 
curl -v localhost:8080/save -d '{"id" :"222", "firstname":"Ham", "lastname":"Cheese", "email": "a@a.a", "age": 20}'

***Find a User***
curl localhost:8080/find/222

