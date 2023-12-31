import React, { useState, useEffect } from "react";
import axios from 'axios';
import Card from "../components/Card";
import { backendUrl } from "../constants/global"

const AlbumList = () => {
    const [albums, setAlbums] = useState([]);

    useEffect(() => {
        axios({
          method: 'get',
          url: `${backendUrl}/api/v1/albums`,
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