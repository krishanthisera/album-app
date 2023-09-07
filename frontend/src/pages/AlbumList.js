import React, { useState, useEffect } from "react";
import axios from 'axios';
import Card from "../components/Card";

const AlbumList = () => {
    const [albums, setAlbums] = useState([]);

    useEffect(() => {
        axios({
          method: 'get',
          url: 'http://localhost:8080/api/v1/albums',
          headers: {
            
            'Content-Type': 'application/json',
          },
        })
        .then(response => {
            setAlbums(response.data);
        })
        .catch(error => {
            console.error('There was an error fetching the data', error);
        });
    }, []);

    return (
        <div>
          <h1>Album List</h1>
          {albums.map(album => (
            <Card key={album.id} album={album} />
          ))}
        </div>
      );
    }
    
    export default AlbumList;