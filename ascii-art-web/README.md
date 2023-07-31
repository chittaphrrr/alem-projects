# ASCII-ART-WEB


## Description

Ascii-art-web is creating and running a server, in which it is possible to use a web GUI (graphical user interface) version of the  ascii-art project.

The webpage allows the use of the different banners:

- shadow
- standard
- thinkertoy
The following HTTP endpoints are implemented:

1. GET /: Sends HTML response, the main page.

2. POST /: that sends data to Go server (text and a banner)

The result from the POST is displayed in the route / after the POST is completed.

The webpage consists of the following elements:
- text input
- select object to switch between banners
- submit button, which sends a POST request to '/' and outputs the result on the page.

## Usage: how to run

To deploy this project run

bash
  go run ./cmd/web



## Authors

- [@nsabyrov](https://01.alem.school/git/nsabyrov)