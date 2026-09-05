# pull a base image whichi gives all required tools and libraries
FROM eclipse-temurin:17-jdk-alpine

# create a folder where the app code will be stored 
WORKDIR /app 

# Copy the source code from your HOST machine to your container
COPY src/Main.java .

# Compile the application code
RUN javac Main.java

# Run the application
CMD ["java" , "Main"]
