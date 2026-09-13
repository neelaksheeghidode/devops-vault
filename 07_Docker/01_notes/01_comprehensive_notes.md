1. Introduction to Docker

    The Big Problem: In software development, we always face the classic excuse: "It works on my machine, but it’s broken on yours!"

    Why It Happens: Developers use different operating systems, libraries, and settings. Code that runs smoothly on Windows might crash on Mac or Linux.

    The Docker Solution: Docker packs your app along with all its required libraries and settings into a single isolated box called a Container.

    The Magic: Once it is inside a Docker container, your application runs identically on any computer in the world, completely eliminating environmental issues.

2. Containers vs. Virtual Machines (VMs)

    Virtual Machines (Heavy): Every VM needs its own full Operating System installed. Because of this, they consume massive amounts of RAM and CPU, making them heavy and slow.

    Docker Containers (Lightweight): Containers share the host machine's Operating System kernel. They are super lightweight and boot up in seconds, allowing you to run multiple containers easily even on a normal laptop (like an 8GB RAM machine).

3. Docker Installation & Setup
Local Setup (Your Laptop)

    Go to the official Docker website and download Docker Desktop.

    Install it, and you get a clean dashboard where you can easily view, manage, and monitor all your running containers.

Cloud Setup (AWS)

    If you don't want to load your local laptop, you can spin up a virtual machine instance on AWS (Amazon Web Services).

    Use an SSH (Secure Shell) command from your terminal to connect securely to your cloud machine and install Docker directly onto the cloud server for remote testing.

4. Deep Dive: How Docker Works (Under the Hood)

When you type a Docker command, a well-organized background machinery handles your request:

    Docker Client: This is where you sit and type commands (like docker run). It acts as the user interface.

    Docker API: The communication bridge that takes your commands from the client and sends them to the core brain.

    Docker Daemon (dockerd): The background service responsible for managing everything—creating, running, stopping, and deleting your containers.

    Containerd: The industry-standard engine tucked inside Docker that handles the actual lifecycle of the container.


## Docker

Docker was brought in so that when you run code on someone else's machine and it breaks , that problem gets solved. Docker carries its own environment, meaning it travels with its own little world 🤏. But a container is just an empty box—so how do we actually run it? That's where Docker Engine comes in.
It is not some heavy OS rather, it is a veryyy lightweight client-server application that quietly runs in the background using the current OS.
And to make it work we use
Docker CLI: This is where you give your Docker commands. Like docker run etc..
REST API: This acts as the bridge that carries your commands to the Docker Daemon. 
Docker Daemon or dockerd: This is the main part—the mastermind that handles and does everything else.
