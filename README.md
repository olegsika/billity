# billity

For start the application need
1. Run DB condtainer

  ```docker-compose up```
  
2. Run Rammit MQ

  ```docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management```
  
3. Init the migrations ms
  Run the migration ms with ```init``` argument
  Run the migration ms without aruments
  
4. Run API ms
5. Run Worker ms

For api you can use the Postman Collection by link https://www.getpostman.com/collections/1d514fc4dd1046ad0ec1
