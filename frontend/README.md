
# React Album Display Module

This module provides components for displaying album details and lists in a React application.

## Description

The module contains two main React components:

1. **AlbumDetail** - A component to display the detailed view of an individual album.
2. **AlbumList** - A component to display a list of albums.

## Installation

To install the required dependencies for this module, navigate to the project root directory and run:

```bash
npm install
```

## Docker Setup

A Dockerfile has been provided for containerization. To build and run the application using Docker:

1. **Build the Docker Image**:

```bash
docker build -t album-app-frontend .
```

2. **Run the Docker Container**:

```bash
docker run -p 3000:3000 album-app-frontend
```

This will map port 3000 inside the container to port 3000 on your machine.

## Usage

1. **AlbumDetail**:

To use the `AlbumDetail` component, import it in your React file:

```javascript
import AlbumDetail from './path-to/AlbumDetail';
```

Then, you can use it in your JSX:

```jsx
<AlbumDetail album={yourAlbumObject} />
```

2. **AlbumList**:

Similarly, for the `AlbumList` component:

```javascript
import AlbumList from './path-to/AlbumList';
```

And in your JSX:

```jsx
<AlbumList albums={yourAlbumsArray} />
```

## Additional Information

Ensure you have the correct version of Node.js installed, as specified in `package.json`. Also, follow best practices when deploying Docker containers in production.
