import React, {useEffect, useState} from "react";
import {listCharacters, Character} from '../api/listCharacterApi'

const ListCharactersContainer: React.FC = () => {
    const [characters, setCharacters] = useState<Character[]>([])
    const [error, setError] = useState<string | null>(null)

    useEffect(() => {
        const fetchCharacters = async () => {
            try{
                const response = await listCharacters()
                setCharacters(response.characters)
                setError(null)
            } catch(err){
                setError('failed to load characters')
            }
        };
        fetchCharacters();
    }, [])

    return(
        <div>
            <h2>characters</h2>
            {error && <p>{error}</p>}
            {characters.length === 0 ? (
                <p>no characters</p>
            ):(
                <ul>
                    {characters.map((character, index)=> (
                        <li>
                            <p>{character.name} {character.race} {character.characterClass} {character.level}</p>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default ListCharactersContainer