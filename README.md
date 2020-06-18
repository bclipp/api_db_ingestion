![Python application](https://github.com/bclipp/api_db_ingestion/workflows/Python%20application/badge.svg)  
![Go](https://github.com/bclipp/api_db_ingestion/workflows/Go/badge.svg)

## API DB Ingestion

## Implementation
[Go](https://github.com/bclipp/api_db_ingestion/tree/master/go)  
[Python](https://github.com/bclipp/api_db_ingestion/tree/master/python)  
## Objective:
Write an application that combines the data in the local company database with an external API.

## Requirements:

The app should be able to run on a system with limited resources.

## Data:

1. Local Database running PostgreSQL provided via a docker container.
Which will contain the following customer data tables:
* customers
* stores

2. The External API:

swagger doc:

https://geo.fcc.gov/api/census/#!/block/get_block_find

example api call:  

https://geo.fcc.gov/api/census/area?lat=37.299590&lon=-76.742290&format=json


Outcome:

Customers and stores are updated with the following information
blockID or block fips id, state_fips, state code ,and block population.


## Docker Setup

1. Install docker and docker-compose

2. run docker-compose up

## Connecting to Database UI

*URL*: http://localhost:8080/  
*SYSTEM*: PostgreSQL  
*USER*: project01  
*PASSWORD*: project01  
*DB*: project01  

## Running the app

### Local Usage

export POSTGRES_DB=project01  
export POSTGRES_USER=project01  
export POSTGRES_PASSWORD=project01  
export DB_IP_ADDRESS=127.0.0.1  

#### setup python environment

##### Initial setup

```apt-get update && \
sudo apt install -y python3-pip python3

sudo apt-get install -y make build-essential libssl-dev zlib1g-dev libbz2-dev \
  libreadline-dev libsqlite3-dev wget curl llvm libncurses5-dev libncursesw5-dev \
  xz-utils tk-dev libffi-dev liblzma-dev python-openssl git 

curl https://pyenv.run | bash 

export PATH="$HOME/.pyenv/bin:$PATH" 
eval "$(pyenv init -)" 
eval "$(pyenv virtualenv-init -)"

pyenv install 3.8.0 

pyenv virtualenv 3.8.0 app_3.8
```

##### using pyenv environment
pyenv activate app_3.8  
pip3 install -r requirements.txt  
sudo --preserve-env=POSTGRES_DB,POSTGRES_USER,POSTGRES_PASSWORD docker-compose up  

#### if you need to access the docker container
sudo docker ps
sudo docker exec -it <container name> bash

### Continuous Integration
[Github Actions CI YAML](https://github.com/bclipp/api_db_ingestion/blob/master/.github/workflows/python-app.yml)
