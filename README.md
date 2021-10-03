# poker hand compare 

## Requirement

- golang@1.16 

## Install dependencies

```bash
go mod download
```

## Run the program directly

```bash
go run main.go <first_hand> <second_hand>
```

## Run the program using docker

```bash 
docker build -t <image_tag> . 
docker run <image_tag> <first_hand> <second_hand>
```

## Run unit test 

```bash
go test .
```
