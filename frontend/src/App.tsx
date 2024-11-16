import React from 'react';
import './App.css';
import DiceRoller from './components/DiceRoller';


const App: React.FC = () => {
  return (
    <div>
      <h1>D&D Character Manager</h1>
      <DiceRoller />
    </div>
  );
};

export default App;
