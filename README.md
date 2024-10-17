# rabbitmq-lock
Using RabbitMQ as distributed lock

Example run:
```shell
docker compose up --build
...
Attaching to rabbitmq-lock-1, rabbitmq-lock-2, rabbitmq-lock-3, rabbitmq-lock-4, rabbitmq-lock-5
rabbitmq-lock-3  | 2024/10/17 13:02:13 starting app
rabbitmq-lock-1  | 2024/10/17 13:02:13 starting app
rabbitmq-lock-5  | 2024/10/17 13:02:13 starting app
rabbitmq-lock-2  | 2024/10/17 13:02:13 starting app
rabbitmq-lock-4  | 2024/10/17 13:02:13 starting app
rabbitmq-lock-2  | 2024/10/17 13:02:15 i don't have the lock :(
rabbitmq-lock-5  | 2024/10/17 13:02:15 i don't have the lock :(
rabbitmq-lock-3  | 2024/10/17 13:02:15 i have the lock :)
rabbitmq-lock-4  | 2024/10/17 13:02:15 i don't have the lock :(
rabbitmq-lock-1  | 2024/10/17 13:02:15 i don't have the lock :(
rabbitmq-lock-2  | 2024/10/17 13:02:20 i don't have the lock :(
rabbitmq-lock-3  | 2024/10/17 13:02:20 i don't have the lock :(
rabbitmq-lock-1  | 2024/10/17 13:02:20 i have the lock :)
rabbitmq-lock-4  | 2024/10/17 13:02:20 i don't have the lock :(
rabbitmq-lock-5  | 2024/10/17 13:02:20 i don't have the lock :(
```