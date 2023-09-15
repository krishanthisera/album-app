import React from "react";

import Card from "../components/Card";

export const getServerSideProps = async (context) => {
  const res = await fetch("http://localhost:8080/api/v1/albums");
  const albums = await res.json();
  return { props: { albums } };
};

const AlbumList = ({ albums }) => {
  return (
    <div>
      <h1>Album List</h1>
      {albums.map((album) => (
        <Card key={album.id} album={album} />
      ))}
    </div>
  );
};

export default AlbumList;
