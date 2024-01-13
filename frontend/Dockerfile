FROM node:20-alpine

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the package.json and package-lock.json (or yarn.lock if using yarn) to the container
COPY package*.json ./

# Install the app's dependencies inside the container
RUN npm ci

# Copy the rest of the application's files into the container
COPY --chown=node:node . .

# Expose port 3000 for the application
EXPOSE 3000

# Start the application
ENTRYPOINT [ "/bin/sh","-c","npm run dev:start" ] 