import React from 'react';

const styles = {
  card: {
    margin: 20,
    background: '#e8eaf6',
  },
  heading: {
    background: '#3f51b5',
    minHeight: 50,
    lineHeight: 3.5,
    fontSize: '1.2rem',
    color: 'white',
    padding: '0 20px',
  },
  content: {
    padding: 20,
  },
};

function AlbumCard({ album }) {
  return (
    <div style={styles.card}>
      <div style={styles.heading}>{album.title}</div>
      <div style={styles.content}>
        <h2>Artist: {album.artist}</h2>
        <h3>Price: ${album.price}</h3>
      </div>
    </div>
  );
}

export default AlbumCard;
