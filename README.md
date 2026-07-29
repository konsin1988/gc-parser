# GC Parsing project 

> The goal is to get data of goods and sellers from popular marketplaces, 
> to create an REST API to send this data to client and to develop UI 
> to check data about sellers and goods

### Steck
- Database - postgres (opensource and powerful);
- Api - golang (fast and simple);
- Agent - golang (fast and simple);
- UI - React + Typescript + Vite (most popular and typesafe);


## Worker
> It runs as docker container with specific command to get target data

### Commands to parse data

#### Ozon
##### Goods from search query 
```go run main.go search "Ноутбук xiaomi" --max-pages 10```

##### Brand goods
```go run main.go brand 551 --max-pages 20```

##### Seller goods
```go run main.go seller 3001138 --max-pages 5```

##### Category goods
```go run main.go category 12 --max-pages 5```

## Api

### Endpoints

#### Sellers List with filters
```/sellers?brand=1&min_goods=5&min_score=4.87&category=285```
