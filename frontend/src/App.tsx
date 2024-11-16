import React from 'react';
import './App.css';
import DiceRoller from './components/DiceRoller';
import CharacterCreator from './components/CharacterCreator';
import ListCharactersContainer from './components/ListCharactersContainer';


const App: React.FC = () => {
  return (
    <div>
      <h1>D&D Character Manager</h1>
      <DiceRoller />
      <CharacterCreator />
      <ListCharactersContainer />
    </div>
  );
};

export default App;
