# sketchtest

Test submission for the sketch takehome assessment

## Features
- A server that draws ascii art based on payload
- Canvas is stored with an identifier and can be fetched with the same endpoint without a body.
- Database migrations
- Docker compose for easy running

## Running
cd into the project folder and run `docker compose up` after which you should have a postgres service and the web server running

## Server documentation
A single endpoint is exposed on the server `POST /draw?id=<arbitrary_id>`.

If a json body is passed, then that body is uded to generate the canvas and art, if a body is not passed, then the body from the previous request is used 
(if there is no body for that identifier, then a 404 is returned)

### body format
Here is a sample response 

```
{
  "rectangles": [
    {
      "start_x": 15,
      "start_y": 0,
      "width": 7,
      "height": 6,
      "outline": "",
      "fill": "."
    },
    {
      "start_x": 0,
      "start_y": 3,
      "width": 8,
      "height": 4,
      "outline": "O",
      "fill": ""
    },
    {
      "start_x": 5,
      "start_y": 5,
      "width": 5,
      "height": 3,
      "outline": "X",
      "fill": "X"
    }
  ],
  "fills": [
    {
      "start_x": 0,
      "start_y": 0,
      "character": "-"
    }
  ]
}
```

And here is the output for the above request
```
---------------.......-
---------------.......-
---------------.......-
OOOOOOOO-------.......-
O      O-------.......-
O    XXXXX-----.......-
OOOOOXXXXX-------------
-----XXXXX-------------
-----------------------

```

### Validation rules
- `Rectangles` is required
- `Fills` is optional
- start coords and dimensions are required
- either one of fill/outline is required


## Running tests
To run tests set the following env variables and run `go test ./...`
```
	POSTGRES_HOST
	POSTGRES_PASSWORD 
	POSTGRES_USERNAME 
	POSTGRES_DBNAME
	POSTGRES_PORT
	POSTGRES_MIGRATIONS (path to the migrations folder)
```
