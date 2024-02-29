import React, { useState, useEffect } from 'react';
import Card from './Card';
import SearchBar from './SearchBar';

const Dashboard = ({ token }) => {
  const [cards, setCards] = useState([]);

  useEffect(() => {
    const fetchCards = async () => {
        try {
          const response = await fetch('http://localhost:8080/api/card', {
            method: 'GET',
            headers: {
              Authorization: `${token}`,
            },
          });
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          const data = await response.json();
          const cards = data.cards;
          setCards(cards);
        } catch (error) {
          console.error('Error fetching cards:', error);
        }
      };

    fetchCards();
  }, [token]);


  const handleAddCard = async (value,label ) => {
    try {
        const response = await fetch('http://localhost:8080/api/card', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `${token}`,
          },
          body: JSON.stringify({ 
            breedPath: value,
            breedLabel: label
        }),
        });
  
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
  
        const data = await response.json();
        const card = (data.card);
        const newCard = { breed: label, photo: card.photo };
        setCards(prevCards => [...prevCards, newCard]);
      } catch (error) {
        console.error('Error getting photo:', error);
      }
    
  };
  const deleteAllCards = async() => {
    try {
    const response = await fetch('http://localhost:8080/api/card', {
          method: 'DELETE',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `${token}`,
          }});
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          setCards([])
    } catch (error) {
        console.error('Error getting photo:', error);
      }
  };
  const handleDeleteCard =async(id)=>{
    try {
        console.log("Deleting " + id);
        const response = await fetch(`http://localhost:8080/api/card/${id}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `${token}`,
              }
            }
        );
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

    setCards(prevState => prevState.filter(c => c.id !==id));
    } catch (error) {
        console.error('Error deleting card:', error);
      }
  };


  return (
    <div>
      <h2>Doggo Collector</h2>
      <SearchBar onAddCard={handleAddCard} />
      <button onClick={deleteAllCards}>Clear All Cards</button>
      <h2>Caught Breeds</h2>
      <div className="card-list">
        {cards.map((card, index) => (
          <Card key={index} title={card.breed} photo={card.photo} id={card.id} onDeleteCard={handleDeleteCard} />
        ))}
      </div>
    </div>
  );
};

export default Dashboard;
