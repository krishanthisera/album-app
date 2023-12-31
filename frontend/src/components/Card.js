import React from 'react';
import { useNavigate } from 'react-router-dom';

// We declare an object called styles that will contain a few objects for card and heading styles
// Notice that each key lists CSS styles in camel case
const styles = {
  card: {
    margin: 20,
    background: '#e8eaf6',
    cursor: 'pointer',
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

function Card({ album }) {
  const navigate = useNavigate();

  // Function to handle card click
  const handleCardClick = () => {
    navigate(`/album/${album.id}`);
  };

  return (
    <div style={styles.card} onClick={handleCardClick}>
      <div style={styles.heading}>{album.title}</div>
    </div>
  );
}

export default Card;

