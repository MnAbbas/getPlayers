# getPlayers  
get a list of players and formatting the output and array

# Run

``` go  
go run main.go

```

# Test

``` go  
go  test -v test/*.go

```

# Challenges
There were two challeges in doing the task

## Challenge 1
Teamid limit which must be a unsigned interger which means we have tree options , uint , uint8 , uint16 , uint32 , uint64 
but base on my guess i used uint8 which means maximum 255 

``` go  
MaxUint = ^uint8(0)

```

## challenge 2
speed , I could use chan or Waitgroup to incrrease speed of retrieving data , but because I had to wait to all retrieve data be finished I decided to use waitgroup 

``` go  
var wg sync.WaitGroup

```

## challenge 3

Timeout of http Get client

``` go
timeout := time.Duration(10 * time.Second)
client := http.Client{
Timeout: timeout,
}

```