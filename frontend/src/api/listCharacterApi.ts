import { Character } from "../types"

export interface ListCharactersResponse {
    characters: Character[]
    message: string
}

export async function listCharacters(): Promise<ListCharactersResponse>{
    const response = await fetch('http://localhost:8080/characters',{
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })

    if (!response.ok){
        throw new Error('failed to list characters')
    }

    const result: ListCharactersResponse = await response.json()
    return result;
}