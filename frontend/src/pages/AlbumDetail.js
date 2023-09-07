import React, { useState, useEffect } from "react";
import axios from 'axios';

const AlbumDetail = (props) => {
    const [album, setAlbum] = useState(null);
    const albumId = props.match.params.id; 
  

    useEffect(() => {
        axios({
          method: 'get',
          url: `http://localhost:8080/api/v1/album/${albumId}`,
          headers: {
            'Content-Type': 'application/json',
          },
        })
      .then(response => {
        setAlbum(response.data);
      })
      .catch(error => {
        console.error("There was an error fetching the data", error);
      });
    }, [albumId]);
  
    return (
      <div>
        {album ? (
          <div>
            <h1>{album.title}</h1>
            <h1>{album.artist}</h1>
            <h1>{album.price}</h1>
          </div>
        ) : (
          <p>Loading...</p>
        )}
      </div>
    );
  }
  
  export default AlbumDetail;