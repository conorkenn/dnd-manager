import React from 'react'
import { Link } from 'react-router-dom'

const HomePage: React.FC = () => {
    return(
        <div>
            <h1> DND Manager</h1>
            <p><Link to="/create-character">Create Character</Link></p>
            <p><Link to="/characters">Characters</Link></p>
        </div>
    )

}

export default HomePage;