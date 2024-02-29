import React from 'react';
import './Card.css'; // Import the CSS file

const Card = ({ title, photo,id, onDeleteCard }) => {
    const deleteCard = (id) =>{
        console.log("deleting "+id)
        onDeleteCard(id)
    }
  return (
    <div className="card">
      <h3>{title}</h3>
      <img src={photo} alt={title} />
      <button onClick={()=>deleteCard(id)}>Delete</button>
    </div>
  );
};

export default Card;
