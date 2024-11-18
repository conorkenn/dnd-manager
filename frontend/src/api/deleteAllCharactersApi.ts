export interface DeleteAllResponse {
    message: string
}
export async function deleteAllCharacters(): Promise<DeleteAllResponse> {
    const response = await fetch('http://localhost:8080/delete',{
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        }
    })

    if (!response.ok){
        throw new Error('failed to delete characters')
    }

    const result: DeleteAllResponse = await response.json()
    return result;
}