import React from "react";
import { useRouter } from "next/router";

import AlbumCard from "../../components/AlbumCard";
import Button from "../../components/Button";

export const getServerSideProps = async (context) => {
  const res = await fetch(
    `http://localhost:8080/api/v1/album/${context.params.id}`
  );
  const album = await res.json();
  return { props: { album } };
};

const AlbumDetail = ({ album }) => {
  // const [album, setAlbum] = useState(null);

  const router = useRouter();

  return (
    <div>
      {album ? <AlbumCard album={album} /> : <p>Loading...</p>}
      <Button label="Go Back" onClick={() => router.back()} />
    </div>
  );
};

export default AlbumDetail;
