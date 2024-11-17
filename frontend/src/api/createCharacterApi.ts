export interface Character {
    name: string
    race: string
    characterClass: string
    level: number
    attributes: Attributes
}

export interface Attributes {
    strength: number
    dexterity: number
    constitution: number
    intelligence: number
    wisdom: number
    charisma: number

}

export interface CreateCharacterResponse {
    character: Character
    message: string
}

export interface CreateCharacterRequest {
    name: string
    race: string
    characterClass: string
    level: number
}


export async function createCharacter(data: CreateCharacterRequest): Promise<CreateCharacterResponse>{
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