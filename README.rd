# Listing Docker Image Tag
List up docker image tags from docker-hub api.

## Abstract
You can search docker image tags by using below command.

      $ curl https://registry.hub.docker.com/v1/repositories/<image-name>/tags
      
But response is Json and it it not easy to parse Json by one liner command.

So, I write trivial script for getting image tags from docker-hub API by Go lang.

(But I wonder why docker have not implemented this function.)


## Usage
To install
