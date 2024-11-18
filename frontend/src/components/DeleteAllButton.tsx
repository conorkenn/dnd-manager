import React, {useState} from "react";
import {deleteAllCharacters, DeleteAllResponse} from '../api/deleteAllCharactersApi';

const DeleteAllButton: React.FC = () => {
    const [result, setResult] = useState<DeleteAllResponse | null>(null)
    const [error, setError] = useState<string | null>(null)

    const handleDelete = async () =>{
        try{
            const deleteResult = await(deleteAllCharacters())
            setResult(deleteResult)
            setError(null)
        } catch(err){
            setError("failed to delete characters")
            setResult(null)
        }
    }

    return(
        <div>
            <button onClick={handleDelete}>Delete All</button>
        </div>
    );
};

export default DeleteAllButton