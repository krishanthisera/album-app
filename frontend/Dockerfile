FROM node:18-alpine

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the package.json and package-lock.json (or yarn.lock if using yarn) to the container
COPY package*.json ./

# Install the app's dependencies inside the container
RUN npm ci

# Copy the rest of the application's files into the container
COPY . .

# Start the application
CMD [ "npm", "run", "dev:start" ]

EXPOSE 3000