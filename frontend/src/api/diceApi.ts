export interface RollResponse {
    sides: number
    results: number[]
    num_rolls: number
    message: string
}

export interface RollRequest {
    sides: number
    num_rolls?: number
}

export async function rollDice(data: RollRequest): Promise<RollResponse> {
    const response = await fetch('http://localhost:8080/roll', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })

    if(!response.ok){
        throw new Error('Failed to roll dice')
    }

    const result: RollResponse = await response.json()
    return result;
}