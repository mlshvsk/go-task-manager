# go-task-manager

## Swagger API doc
https://app.swaggerhub.com/apis-docs/mlshvsk/task-manager/1.0.0

## How to run project locally

1. Clone repository
2. Start docker wth `docker-compose up`
3. Run migrations (should be done only once)
   
   3.1 Enter container: `docker-compose run app_api bash`
   
   3.2 Run inside container: `goose -dir ./database/migrations/ mysql "app:password@(app_mysql:3306)/app" up`
4. App is available via http://localhost:8080/api/v1


## Try it out remotely
1. Use http://ec2-52-214-112-128.eu-west-1.compute.amazonaws.com:8080/api/v1 to work with remote project deployed on aws (Use Swagger doc for request references)
2. Please, keep it clean and do not mess it up
3. In case if project is down - ping me
