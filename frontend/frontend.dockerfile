# Use the official Node.js image as the base image
FROM node:18

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json (if available)
COPY package.json package-lock.json ./

# Install the dependencies
RUN npm install

# Copy the rest of the frontend source code to the working directory
COPY . .

# Build the frontend application
RUN npm run build

# Use a smaller base image for serving the application
FROM nginx:alpine

# Copy the build output to the Nginx HTML directory
COPY --from=0 /app/dist /usr/share/nginx/html

# Expose the port the app runs on
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]