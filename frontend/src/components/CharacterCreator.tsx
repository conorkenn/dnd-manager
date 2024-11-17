import React, {useState} from "react";
import {createCharacter, CreateCharacterResponse} from '../api/createCharacterApi';


const CharacterCreator: React.FC = () => {
    const [name, setName] = useState("")
    const [race, setRace] = useState("")
    const [characterClass, setcharacterClass] = useState("")
    const [level, setLevel] = useState(1)
    const [result, setResult] = useState<CreateCharacterResponse | null>(null)
    const [error, setError] = useState<string | null>(null)

    const handleCreation = async () => {
        try{
            const creationResult = await createCharacter({name, race, characterClass, level})
            setResult(creationResult)
            setError(null)
        } catch(err){
            setError('failed to create character')
            setResult(null)
        }
    };

    return(
        <div>
            <h3>character creation</h3>
            <div>
                <label>
                    name:
                    <input type="string"
                    value ={name}
                    onChange={(e) => setName(e.target.value)}
                    />
                </label>
            </div>
            <div>
                <label>
                    race:
                    <input type="string"
                    value ={race}
                    onChange={(e) => setRace(e.target.value)}
                    />
                </label>
            </div>
            <div>
                <label>
                    class:
                    <input type="string"
                    value ={characterClass}
                    onChange={(e) => setcharacterClass(e.target.value)}
                    />
                </label>
            </div>
            <div>
                <label>
                    level:
                    <input type="number"
                    value={level}
                    onChange={(e) => setLevel(Number(e.target.value))}
                    min="1"
                    />
                </label>
            </div>
            <button onClick={handleCreation}>create</button>

            {error && <p style={{color: 'red'}}>{error}</p>}
            {result && (
                <div>
                    <h2>result</h2>
                    <p>name: {result.character.name}</p>
                    <p>race: {result.character.race}</p>
                    <p>class: {result.character.characterClass}</p>
                    <p>level: {result.character.level}</p>
                    <p>Attributes: </p>
                    <p>str: {result.character.attributes.strength}</p>
                    <p>dex: {result.character.attributes.dexterity}</p>
                    <p>cons: {result.character.attributes.constitution}</p>
                    <p>int: {result.character.attributes.intelligence}</p>
                    <p>wis: {result.character.attributes.wisdom}</p>
                    <p>cha: {result.character.attributes.charisma}</p>
                    <p></p>
                </div>
            )}
        </div>
    )
};


export default CharacterCreator;