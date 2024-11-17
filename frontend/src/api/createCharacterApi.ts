import { Attributes, Character } from "../types"

export interface CreateCharacterResponse {
    character: Character
    message: string
}

export interface CreateCharacterRequest {
    name: string
    race: string
    characterClass: string
    level: number
    attributes?: Attributes
}


export async function createCharacter(data: CreateCharacterRequest): Promise<CreateCharacterResponse>{
    data.attributes = data.attributes || {
        strength: 10,
        dexterity: 10,
        constitution: 10,
        intelligence: 10,
        wisdom: 10,
        charisma: 10,
    };
    const response = await fetch('http://localhost:8080/characters', {
        method: 'POST',
        headers: {
            'Content-Type':'application/json'
        },
        body: JSON.stringify(data)
    })

    if (!response.ok){
        throw new Error('Failed to create character')
    }

    const result: CreateCharacterResponse = await response.json()
    return result
}