# ps-tag-onboarding-go


**Run Application**
go run main.go

**Run Mongo**
docker-compose up

**Hit endpoints**

curl localhost:9091/find/222 | jq

curl localhost:9091/users | jq   

curl -v localhost:9091/save -d '{"id" :"444", "firstname":"Ham", "lastname":"Cheese", "email": "a@a.a", "age": 20}'