# Use an official Node.js runtime as the base image
FROM node:14

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the local package.json and package-lock.json (if available) to the container's working directory
COPY package*.json ./

# Install the dependencies inside the container
RUN npm install

# Copy the rest of the application to the container's working directory
COPY . .

# Expose the port the app runs on
EXPOSE 3000

# The command to run the application
CMD ["npm", "start"]