import axios from 'axios'
import React, { useEffect, useState } from 'react'

const WhoAmI = () => {
    const [whoAmI, setWhoAmI] = useState('')

    useEffect(() => {
        const token = localStorage.getItem("token")
        axios.get(
            `${process.env.REACT_APP_API_URL}/v1/users/whoami`,{
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': token !== '' ? `Bearer ${token}`: ''
                }
            }
        ).then(response => {
            console.log(response.data)
            setWhoAmI(response.data)
        })
        .catch(error => alert(error))
    }, [whoAmI])

    return (
        <div>
            <h3>Who AM I Page {whoAmI}</h3>
        </div>
    )
}

export default WhoAmI;
