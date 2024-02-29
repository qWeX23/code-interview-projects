import React, { useState, useEffect } from 'react';
import Select from 'react-select';

const SearchBar = ({onAddCard}) => {
    const [options, setOptions] = useState([]);
    const [selectedOption, setSelectedOption] = useState(null);
    const handleAddCard = () => {
      if (selectedOption) {
        onAddCard(selectedOption.value, selectedOption.label);
        setSelectedOption(null);
      }
    };
    useEffect(() => {
        const fetchOptions = async () => {
            try {
                const response = await fetch('http://localhost:8080/api/dog/breed', {
                    method: 'GET'
                });

                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }

                const data = await response.json();
                const breeds = data.breeds;
                const options = breeds.map(b => ({
                    value: b.path,
                    label: b.display,
                }));

                setOptions(options);
            } catch (error) {
                console.error('Error fetching cards:', error);
            }
        };

        fetchOptions();
    }, []); 
    const handleAddRandomCard = () => {
        if (options.length > 0) {
          const randomIndex = Math.floor(Math.random() * options.length);
          const randomBreed = options[randomIndex];
          onAddCard(randomBreed.value, randomBreed.label);
        }
      };
    return(
    <div>
   <Select
        value={selectedOption}
        onChange={setSelectedOption}
        options={options}
      />
      <button onClick={handleAddCard}>Add Card</button>
      <button onClick={handleAddRandomCard}>Random</button>
    </div>
    )
};

export default SearchBar;
