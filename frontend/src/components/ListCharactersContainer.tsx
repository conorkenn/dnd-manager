import React, {useEffect, useState} from "react";
import {listCharacters} from '../api/listCharacterApi'
import { Character } from "../types";

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
                            <p>Attributes</p>
                            <p>str: {character.attributes.strength}</p>
                            <p>dex: {character.attributes.dexterity}</p>
                            <p>cons: {character.attributes.constitution}</p>
                            <p>int: {character.attributes.intelligence}</p>
                            <p>wis: {character.attributes.wisdom}</p>
                            <p>cha: {character.attributes.charisma}</p>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default ListCharactersContainer