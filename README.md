# GC Parsing project 

> The goal is to get data of goods and sellers from popular marketplaces, 
> to create an REST API to send this data to client and to develop UI (in progress) 
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

#### AllSellers:
    - ID
    - Name
    - Slug

```/sellers/options```

#### SellerList:
    - ID
    - Name
    - Slug
    - Ogrnip
    - Inn
    - GoodsAmount
    - AverageReviewScore

> filterBy: GoodsAmount, Brands, AverageReviewScore, Category 
> orderBy: Name, GoodAmount, AverageReviewScore

```/sellers?brand=1&min_goods=5&min_score=4.2&category=285&sort=score:desc,goods:desc,name:asc```

```/sellers?brand=1&sort=score:desc,goods:desc,name:asc```

```/sellers?brand=1&min_goods=3&max_goods=90&max_score=4.8&sort=score:desc,goods:desc,name:asc```


#### Seller By Id
    - ID				  
    - Name			  	
    - Slug			  
    - Ogrn			  
    - Inn				  
    - GoodsAmount 	  
    - Brands			  
    - Goods			  
    - AverageReviewScore

```/seller/1229288```


#### All brands:
    - ID
    - Title 
    - Slug

```/brands/options```

#### Brand by Id:
	- Info
	- Brands
	- Goods

#### BrandList:
    - ID
    - Name
    - Slug
    - GoodsAmount
    - AverageReviewScore
    - sellerAmount

> filterBy: GoodsAmount, Sellers, AverageReviewScore, Category 
> orderBy: Name, GoodAmount, AverageReviewScore

```/brands?seller=3548727&sort=score:desc,goods:desc,name:asc```

```http://localhost:8000/brands?min_goods=5&sort=score:desc,goods:asc,name:asc&category=218```


#### All Goods
    - ID
    - Title 
    - Slug

```/brands/options```
