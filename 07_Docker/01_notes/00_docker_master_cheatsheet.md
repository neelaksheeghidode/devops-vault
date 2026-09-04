# Complete Master Docker Commands & Concepts Cheat Sheet (Zero to Advanced)

## 1. Docker Environment & System Management
| Command / Syntax | What it does |
| :--- | :--- |
| `docker --version` | Checks the currently installed version of Docker |
| `docker info` | Displays detailed system-wide Docker setup and configurations |
| `docker --help` | Shows general help and global options for docker CLI |
| `sudo systemctl start docker` | Starts the Docker daemon service on Linux |
| `sudo systemctl stop docker` | Stops the Docker daemon service |
| `sudo systemctl restart docker` | Restarts the Docker daemon service |
| `sudo systemctl status docker` | Checks the health and running status of Docker service |
| `sudo systemctl enable docker` | Enables Docker service to start automatically on system boot |

## 2. Docker Image Management
| Command / Syntax | What it does |
| :--- | :--- |
| `docker images` | Lists all locally stored Docker images |
| `docker images -a` | Lists all images including intermediate layers |
| `docker pull <image_name>` | Downloads an image from Docker Hub registry |
| `docker pull <image>:<tag>` | Downloads a specific version or tag of an image |
| `docker build -t <image_name> .` | Builds an image using the Dockerfile in the current directory (`.`) |
| `docker build -f <custom_filename> -t <image> .` | Builds an image using a custom-named Dockerfile |
| `docker build --no-cache -t <image> .` | Builds an image completely from scratch without using cached layers |
| `docker rmi <image_id>` | Deletes/removes a specific local Docker image |
| `docker rmi -f <image_id>` | Forcefully deletes an image even if it's referenced |
| `docker tag <image> <username>/<repo:tag>` | Tags a local image to prepare it for a remote registry push |
| `docker history <image_name>` | Shows the step-by-step history and size of layers in an image |

## 3. Container Lifecycle Operations
| Command / Syntax | What it does |
| :--- | :--- |
| `docker run <image>` | Creates and starts a container in foreground mode |
| `docker run -d <image>` | Runs a container in detached (background) mode |
| `docker run -it <image> bash` | Runs a container interactively and launches a bash shell |
| `docker run -p <host_port>:<container_port> <img />` | Maps a port from the host machine to the container |
| `docker run --name <container_name> <img />` | Assigns a custom readable name to the container |
| `docker run --rm <image>` | Automatically deletes the container the moment it stops |
| `docker run -e KEY=VALUE <image>` | Passes environment variables directly into the container |
| `docker ps` | Lists all currently active/running containers |
| `docker ps -a` | Lists all containers (both running and stopped) |
| `docker ps -q` | Quiet mode: displays only container IDs |
| `docker start <container_id>` | Starts an existing stopped container |
| `docker stop <container_id>` | Gracefully stops a running container |
| `docker restart <container_id>` | Restarts a running or stopped container |
| `docker pause <container_id>` | Pauses all active processes inside a running container |
| `docker unpause <container_id>` | Resumes paused processes inside a container |
| `docker rm <container_id>` | Deletes a stopped container |
| `docker rm -f <container_id>` | Forcefully terminates and deletes a running container |

## 4. Container Monitoring & Debugging
| Command / Syntax | What it does |
| :--- | :--- |
| `docker logs <container_id>` | Outputs standard logs/stdout of a container |
| `docker logs -f <container_id>` | Streams and follows container logs live in real-time |
| `docker exec -it <container_id> bash` | Opens a live interactive shell inside an already running container |
| `docker inspect <container_id>` | Outputs detailed low-level JSON configuration of a container or image |
| `docker top <container_id>` | Displays the active processes running inside a container |
| `docker stats` | Live real-time resource utilization stream (CPU, memory, network stats) |
| `docker port <container_id>` | Lists public port mappings for a specific container |

## 5. Docker Volumes & Data Persistence
| Command / Syntax | What it does |
| :--- | :--- |
| `docker volume create <vol_name>` | Creates a managed named Docker volume |
| `docker volume ls` | Lists all Docker volumes present on the system |
| `docker volume inspect <vol_name>` | Displays detailed info about a specific volume |
| `docker volume rm <vol_name>` | Deletes an unused Docker volume |
| `docker run -v <vol_name>:/container/path <img />` | Mounts a named volume into a container directory |
| `docker run -v /host/path:/container/path <img />` | Performs a bind-mount linking a host folder to a container directory |

## 6. Docker Networking
| Command / Syntax | What it does |
| :--- | :--- |
| `docker network ls` | Lists all active Docker networks (bridge, host, none) |
| `docker network create <net_name>` | Creates a custom user-defined bridge network |
| `docker network inspect <net_name>` | Displays detailed configuration of a custom network |
| `docker network rm <net_name>` | Deletes a custom user-defined network |
| `docker network connect <net> <container>` | Connects a running container to a specified network |
| `docker network disconnect <net> <container>` | Disconnects a container from a network |
| `docker run --network <net_name> <image>` | Starts a container directly attached to a custom network |

## 7. Docker Compose (Multi-Container Management)
| Command / Syntax | What it does |
| :--- | :--- |
| `docker compose up` | Builds, creates, and starts all services defined in `docker-compose.yml` |
| `docker compose up -d` | Starts all multi-container services in detached background mode |
| `docker compose up --build` | Forces rebuilding of images before starting containers |
| `docker compose down` | Stops and removes containers and default networks |
| `docker compose down -v` | Stops services and completely deletes associated named volumes |
| `docker compose ps` | Lists the status of containers managed by the compose file |
| `docker compose logs` | Views unified log output for all services |
| `docker compose logs -f` | Follows logs live for all compose services |
| `docker compose build` | Builds or updates all services defined in the compose file |
| `docker compose stop` | Stops compose containers without removing them |
| `docker compose restart` | Restarts all running compose services |

## 8. Dockerfile Essential Instructions (Minor & Major)
| Instruction | What it does |
| :--- | :--- |
| `FROM <image>` | Sets the base image for subsequent instructions |
| `WORKDIR /path` | Sets the working directory inside the container for subsequent commands |
| `COPY <src> <dest>` | Copies local files or directories into the container image |
| `ADD <src> <dest>` | Similar to COPY, but also supports URL downloads and auto-extraction of tar files |
| `RUN <command>` | Executes build-time commands (like package installation) and creates an image layer |
| `ENV KEY=VALUE` | Sets environment variables that persist during container runtime |
| `EXPOSE <port>` | Documents network ports the container listens on at runtime |
| `CMD ["executable","param"]` | Defines the default command executed when a container starts |
| `ENTRYPOINT ["executable"]` | Configures a container that will run as an executable |
| `USER <username>` | Sets the user name or UID to use for running subsequent instructions |
| `ARG <name>=<default>` | Defines variables that users can pass at build-time via `--build-arg` |

## 9. Docker Hub & Registry Integration
| Command / Syntax | What it does |
| :--- | :--- |
| `docker login` | Authenticates and logs into your Docker Hub registry account |
| `docker logout` | Clears credentials and logs out from the registry session |
| `docker push <username>/<repo>:<tag>` | Uploads/pushes a tagged local image up to Docker Hub |

## 10. System Cleanup & Pruning
| Command / Syntax | What it does |
| :--- | :--- |
| `docker system prune` | Removes all stopped containers, unused networks, and dangling images |
| `docker system prune -a` | Removes all unused images (not just dangling ones) and stopped containers |
| `docker image prune` | Cleans up unused or dangling local images |
| `docker container prune` | Deletes all currently stopped containers at once |
| `docker volume prune` | Deletes all unused Docker volumes safely |
