# BitlyBot
## Discord bot using DiscordGo that shortens long links that come into the server

# What?
* BitlyBot is designed to shorten links that people enter into a Discord channel. It deletes the link that's entered and replaced with a short link so that links don't clog up the discord chat. It scrapes the title, so that there is still some context as to what the link is for.  
* BitlyBot is a self-hosted bot which accepts a Discord Bot Token and a Bitly API token as command-line arguments. The main reason for making this self-hosted is because Bitly accounts are tied to a certain user with rate limits. Using a free account means that each account would get 1000 shortened links per hour and that would be used up very quickly if it was only hosted once.


## Installation 
*  Requires Go installation
*  Discord bot token auth
*  Bitly API token
*  Either host it on your own system, i.e. Rasp Pi or put it into a cloud provider

## To-Do
*  dockerize it 
*  Put it in GCP
*  CI/CD 

