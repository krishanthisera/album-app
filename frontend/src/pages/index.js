import React from "react";

import Card from "../components/Card";

export const getServerSideProps = async (context) => {
  const BACKEND_URL = process.env.BACKEND_URL || "http://localhost:3000";
  const res = await fetch(`${BACKEND_URL}/api/v1/albums`);
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
