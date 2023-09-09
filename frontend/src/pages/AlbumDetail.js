import React, { useState, useEffect } from "react";
import axios from 'axios';
import { useParams, useNavigate } from 'react-router-dom';
import AlbumCard from "../components/AlbumCard";
import Button from "../components/Button";


const AlbumDetail = () => {
    const [album, setAlbum] = useState(null);
    const navigate = useNavigate();
    const { id: albumId } = useParams(); // Destructuring to get the 'id' param and renaming it to 'albumId'
  
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
            <AlbumCard album={album} />
          ) : (
            <p>Loading...</p>
          )}
          <Button label="Go Back" onClick={() => navigate(-1)} />
        </div>
        
      );
  
}

export default AlbumDetail;
