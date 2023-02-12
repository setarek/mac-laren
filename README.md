## Mac Laren
Mac laren is ordering system with two service:
* order which contains order api that get order data and queue them
* order processor which consume from queue and save them into mysql

## How to run project
you should just run these two command to build and run project quickly

```
./script/build.sh
./script/run.sh
```

if you want just build or run one containers, pass the container name to these script like below:

```
./script/build.sh order
```

Note: Don't forget about change sh file's permission


## API

```
curl --location --request POST 'localhost:8080/api/v1/order' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"Downtown Burgur",
    "price":5000
}

```