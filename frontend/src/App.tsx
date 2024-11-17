import React from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import HomePage from './pages/HomePage';
import CreateCharacterPage from './pages/CreateCharacterPage';
import ListCharactersPage from './pages/ListCharactersPage';
import DicePage from './pages/DicePage';


const App: React.FC = () => {
  return (
    <Router>
      <nav>
        <Link to="/"> Home </Link>
        <Link to="/create-character"> Create </Link>
        <Link to="/characters"> List </Link>
        <Link to="/dice"> Dice </Link>
      </nav>
      <Routes>
        <Route path="/" element={<HomePage/>}/>
        <Route path="/create-character" element={<CreateCharacterPage/>}/>
        <Route path="/characters" element={<ListCharactersPage/>}/>
        <Route path="/dice" element={<DicePage/>}/>
      </Routes>
    </Router>

  );
};

export default App;
