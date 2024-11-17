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