# RESTful-API-Go-Endpoint

## What does this project do? :bulb:
As the name suggests, the program is an Endpoint set up for demonstrating REST API using Golang web server, routed through an Nginx server and the basic CRUD are done to and from Mongo database server. All the three modules are encapsulated in Docker containers. This dosen't solve any challenges or disclose new features, but just implements the following flow in practice. The aim is to input the data and get the same data as output passing through all the modules as follows:<br><br>

![alt type](https://github.com/VagueCoder/RESTful-API-Go-Endpoint/blob/master/RESTful-API-Go-Endpoint.png)

## Knowledge (or) Technologies Used :books:
**Sno.** | **Name** | **Usage**
-------: | :------: | :--------
1 | Go (Golang) | Go is a statically typed, compiled programming language designed at Google. Helps in speed and concurrency. More info at [golang.org](https://golang.org/doc/).
2 | REST API | Representational State Transfer (REST) is a software architectural style that defines a set of constraints to be used for creating Web services. The rule Zero of using this is, you'll send everything as JSON objects over the application endpoints (API socket ports). More info at [restfulapi.net](https://restfulapi.net/).
3 | MongoDB | MongoDB is a cross-platform document-oriented database program. Classified as a NoSQL database program, MongoDB uses JSON-like documents with optional schemas. More info at [mongodb.com](https://www.mongodb.com/).
4 | Docker | Docker is a set of platform as a service products that use OS-level virtualization to deliver software in packages called containers. These containers are lightweight virtual-machine kind of platforms which make use of same host OS kernel, but are very scalable and efficient. More info at [docker.com](https://www.docker.com/).
5 | Nginx | Nginx is open source software for web serving, reverse proxying, caching, load balancing, media streaming, and more. Find more at [nginx.com](https://www.nginx.com/).

## Base System Configurations :wrench:
**Sno.** | **Name** | **Version/Config.**
-------: | :------: | :------------------
1 | Operating System | Windows 10 x64 bit + WSL2 Ubuntu-20.04 
2 | Language | Go Version 1.14.7 Windows/amd64
3 | IDE | Visual Studio Code Version 1.52.1
4 | Containerization | Docker Version 20.10.0, Docker-Compose Version 1.27.4
5 | Database | MongoDB Version 4.4.2
6 | Network Proxy | Nginx Version 1.19.6

> These doesn't probably effect the usage, as the application runs in Docker containers. But of course, the development steps may differ as per your configuration. The required softwares/configurations are mentioned under **Prerequisites** section.

## Prerequisites :file_folder:
**Sno.** | **Software** | **Detail** | **Download Links/Steps** |
-------: | :----------: | :--------: | :----------------------: |
1 | Docker Version 20.10.0 or Higher | Containerizes the application modules for using them as services. This also creates containers of Golang, Nginx and MongoDB avoiding to download stand-alones for the same. | [docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop)
2 | Docker-Compose Version 1.27.4 or Higher | A CLI of Docker that helps to run [docker-compose.yml](https://github.com/VagueCoder/RESTful-API-Go-Endpoint/blob/master/docker-compose.yml) which is helps in building/starting/stopping all the containers at once and with ease. If using Windows (or) Mac, the Docker-Compose automatically gets downloaded along with Docker. | [docs.docker.com/compose/install/](https://docs.docker.com/compose/install/)
3 | Postman (or any equivalent) | Making the GET requests, and most importantly, the POST requests to the API are made easy with Postman. | [postman.com/downloads/](https://www.postman.com/downloads/)

> If you're using Postman, for avoiding the issue with Proxy while running this application, go to **Postman -> File -> Settings -> Proxy** and uncheck "Use the system proxy" option.

## Useful Socket-Ports: :handshake:
**Sno.** | **Port Number** | **Endpoint** | **Exposed to Host** | **Comment**
-------: | :-------------: | :----------: | :---------: | :----------
1 | 8080 | Nginx Proxy | Yes | This is statically binded in [docker-compose.yml](https://github.com/VagueCoder/RESTful-API-Go-Endpoint/blob/master/docker-compose.yml) file. However, you may use a different port for the application.
2 | 8081 | Go Web API | No | This is internal to docker network alone, and not exposed to host as we expect the routing to happen through the Nginx. There is no fixed port for this. You may change the number in [docker-compose.yml](https://github.com/VagueCoder/RESTful-API-Go-Endpoint/blob/master/docker-compose.yml) under <br>services -> restful-api -> environment -> API_PORT.  
3 | 27017 | MongoDB | No | This is internal to docker network alone and not exposed to host as the host mongod server's default port 27017 may collide. However, you can map the internal 27017 with any other port on the host system.

## Setup Application in Local :bookmark_tabs:
Following the steps to recreate the application in your local (in Docker Containers) to scrape and save to database.
1. Download the whole repo and place anywhere in the local. Go's location constraint doesn't apply as application runs in containers. 
2. Open any terminal in the same location and build using docker-compose.
```
docker-compose build
```
3. Run the application (includes all module containers along with network) at once using the command:
```
docker-compose up -d
```
The option `-d` (short form of `--detach`) here means the container runs in background.

Output should be:
```
Creating network "rest_api_go_endpoint_network" with the default driver
Creating MongoDB ... done
Creating RESTful-API-Go-Endpoint ... done
Creating Nginx-Proxy             ... done
```

4. To verify the services/containers processes running, run the following commands:
```
docker-compose ps
docker ps
```
The first command, `docker-compose ps` gives the status of all the services running in the same directory and the second, `docker ps` shows all the docker processes running on the desktop.

5. Close the services/application:
```
docker-compose down
```
This closes all the services gracefully (the word it uses for non-force shutting). And removes the containers and network so created.

`Note: The automatic removal process wipes out only the containers and network that got created, and leaves the 2 images that are built (Scraper-API & Collector-API) and of course, the downloaded images (mongo & golang) remain. This is good. If not removed, the containers might collide with the new containers that will be build and might also lead to build failures.`

## Funcitionality Testing :mag:

Simple 2 step process:
#### 1. Load the data using POST request
* The data should be sent as form data JSON to `localhost:8080` which can be of any structure. A sample:
```
{
	"records": [
		{
			"name": "Alfreds Futterkiste",
			"city": "Berlin",
			"country": "Germany",
			"famous foods":	{
				"count": 3,
				"foods": [
					"Currywurst",
					"Döner Kebab",
					"Bockwurst"
				]
			}
		},
		{
			"name": "Ana Trujillo Emparedados y helados",
			"city": "México D.F.",
			"country": "Mexico"
		}
	]
}
```

You should get an ID from Mongo server something like `{"InsertedID":"5ff46d7b9927aae97da877b8"}`. This confirms the load operation.

#### 2. Retrieve the data using GET request
* Just hit the url `localhost:8080` using any browser or Postman, and you should get the data loaded earlier, probably in flattened style.
```
[{"records":[{"city":"Berlin","country":"Germany","famous foods":{"count":3,"foods":["Currywurst","Döner Kebab","Bockwurst"]},"name":"Alfreds Futterkiste"},{"city":"México D.F.","country":"Mexico","name":"Ana Trujillo Emparedados y helados"}]}]
```

## Using Postman :email:
As explained above, postman helps in making the calls, especially the POST method calls to hosts. Make the proxy disable as explained under `Prerequisites` section.
The steps are pretty simple.
1. Download Postman from official site [postman.com/downloads/](https://www.postman.com/downloads/).
2. Install and launch the application in local.
3. Select the method (GET/POST) from the drop-down and enter the above mentioned URLs (one at a time) in the corresponding space.
4. If POST requests, you'll find the `Body` option just below the URL. Click and go as follows:<br>
   **Body (Sub Menu) -> Raw (Radio Button) -> JSON (from Drop-down)** <br>
   This will enable the description form where you can copy-paste JSON data. This works more efficiently than selecting individual Key-Value pairs.
Use the URLs, methods, form data and check for outputs in this with that of mentioned under `Making the Calls` section above.


#### This concludes everything that is required to check and make use of the [Amazon-Scraper-Collector](https://github.com/VagueCoder/Amazon-Scraper-Collector). The code walk-throughs will be added in the future developments on this. For any issues, queries or discussions, please update in issues menu or write to `vaguecoder0to.n@gmail.com`.

## Happy Coding !! :metal:
