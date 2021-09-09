# autocomplete-with-elasticsearch

A simple scalabale implimentation of search autocomplete using elasticsearch 7 . On the backend I used Golang.

This solution can suggest top N word based on search frequency .

## How To Run
Using Docker -

git clone https://github.com/mdnurahmed/autocomplete-with-elasticsearch
cd autocomplete-with-elasticsearch
docker-compose up 
Then go to localhost:3000 in the browser . I have included kibana with the docker-compose file. So if you wanna see how the search strings are stored in elasticsearch go to localhost:5601 . 