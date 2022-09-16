# Get Average Blocks per Minute

## Start Prysm

Run the following commands to create directories:

```
mkdir ethereum
cd ethereum
mkdir execution
mkdir consensus
```

Create `jwt.hex` file inside consensus directory and put data from https://seanwasere.com/generate-random-hex/

The start `geth` using:

```
docker run -it -d -p 8545:8545 -p 30303:30303 -v /host/path/ethereum:/root/.ethereum  \
    ethereum/client-go:v1.10.23 \
    --http --http.api eth,net,engine,admin \
    --authrpc.jwtsecret /root/.ethereum/consensus/jwt.hex \
    --datadir /root/.ethereum/execution \
    --keystore /root/.ethereum/execution
```

Then start `prysm` using:

```
docker run -d -v /host/path/ethereum:/root/.ethereum -p 3500:3500 -p 4000:4000 -p 13000:13000 -p 12000:12000/udp --name beacon-node \
  gcr.io/prysmaticlabs/prysm/beacon-chain:stable \
  --datadir=/root/.ethereum/consensus \
  --jwt-secret=/root/.ethereum/consensus/jwt.hex \
  --rpc-host=0.0.0.0 \
  --grpc-gateway-host=0.0.0.0 \
  --monitoring-host=0.0.0.0 \
  --execution-endpoint=/root/.ethereum/execution/geth.ipc \
  --accept-terms-of-use
```

## Start PostgresQL

Use the below command to run PostgresQL

```
docker run -p 5455:5432 -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test -e POSTGRES_DB=tracker -d postgres
```

## Run API Server

```
git clone git@github.com:narayanprusty/average-blocks.git
cd average-blocks
go build && ./average-blocks
```

### Curl Requests for Testing

#### Register 

```
curl --location --request POST 'http://localhost:8000/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "narayan",
    "password": "password"
}'
```

#### Login

```
curl --location --request POST 'http://localhost:8000/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "narayan",
    "password": "password"
}'
```

#### Generate API Key

```
curl --location --request GET 'http://localhost:8000/key' \
--header 'Token: <token>'
```


#### Fetch Rate

```
curl --location --request GET 'http://localhost:8000/rate' \
--header 'Api-Key: <api-key>'
```



