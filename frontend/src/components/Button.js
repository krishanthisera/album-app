import React from 'react';

const styles = {
  button: {
    padding: '10px 20px',
    margin: 20,
    border: 'none',
    borderRadius: '4px',
    background: '#3f51b5',
    color: 'white',
    fontSize: '1rem',
    cursor: 'pointer',
    transition: 'background 0.3s',

   
    '&:hover': {
      background: '#303f9f'
    }
  }
};

function Button({ label, onClick }) {
  return (
    <button style={styles.button} onClick={onClick}>
      {label}
    </button>
  );
}

export default Button;
