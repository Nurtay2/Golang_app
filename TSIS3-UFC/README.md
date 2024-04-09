# UFC
Welcome to the League of Legends Project, a simple Go project inspired by the world of League of Legends. 

## UFC REST API

```
GET /champions
POST /champions
GET /champions/:id
PUT /champions/:id
DELETE /champions/:id
```

## DB Structure

```
// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs

// Use DBML to define your database structure
// Docs: https://dbml.org/docs

Table champions {
  id serial [primary key]
  name varchar(255)
  class varchar(255)
  price int
}

Table items {
  id serial [primary key]
  name varchar(255)
  description text
  cost int
}

Table matches {
  id serial [primary key]
  date timestamp
  duration interval
  winner_team_id int
}

Table players {
  id serial [primary key]
  summoner_id int
  champion_id int
  role varchar(50)
}

Table teams {
  id serial [primary key]
  player_ids int[]
}
```

## Team:

Nurzhan Kurmangazy 22B030492