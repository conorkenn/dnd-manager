import React from 'react'
import ListCharactersContainer from '../components/ListCharactersContainer'
import DeleteAllButton from '../components/DeleteAllButton'

const ListCharactersPage: React.FC = () => {
    return (
        <div>
            <ListCharactersContainer />
            <DeleteAllButton />
        </div>
    )
}

export default ListCharactersPage