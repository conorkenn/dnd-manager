import { Character } from "../types";


export interface GetCharacterResponse {
    character: Character
    message: string
}

export interface GetCharacterRequest{
    name: string 
}

export async function getCharacter(data: GetCharacterRequest): Promise<GetCharacterRequest>{
    const response = await fetch(`http://localhost:8080/characters/${data.name}`)
    if (!response.ok){
        throw new Error('Failed to fetch character')
    }
    return response.json()
}