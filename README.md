[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)



# DiscordHackWeek

## Using

There are currently two commands:

- ```-allChat sampleString```  

  Saves the inputted message into a PostgreSQL DB. It also stores the Author'name, the guild ID, and the channel ID
  
- ```-allChat```

  Outputs the most recent ten messages into the current channel formmated as (ID# - Author: sampleString)

## Purpose

I wanted to create a way for different communities to interact. Sometimes these interactions can be funny, wholesome, or create interesting patterns. My first idea was to have each message sent to the bot tweeted out on a public Twitter account. However I still haven't had my developer account approved so that was added to the pending. I think the idea of merging guilds is an interesting concept and hope it can be explored more in the future.

## Creating

To run this bot you will need to have Golang(https://golang.org/) and PosrgreSQL installed

1. You will need to run your own PostgreSQL DB and create a file. Here is the script I used..

2. Clone the repo and add config.js in the root directory

```
config.js

{
    "host"      : "localhost",
    "port"      : 5433,
    "user"      : "postgres",
    "password"  : "admin1",
    "dbname"    : "my-db-name",
    "bottoken"  : "G1kqODE4NTY1MjM4MTI4ZjQx.XQ2_Yw.KNdeYQez9VxFnWaXU2Q9NWXF9_E"
}
```

3. After cloning run ```go run .\discordHackWeek.go``` in the terminal. You should see mulitple consoles indicating the server is running.

## // TODO

- add a timestamp field
- output a detailed list 
- tweet new messages with Twitter API
- open API to grab all data
