import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { Icon, Image } from 'semantic-ui-react'

const Sponsor = ({sponsorName}) => {
    const [sponsor, setSponsor] = useState({})

    useEffect(() => {
        axios.get(`${process.env.REACT_APP_API_URL}/v1/socialmedia/${sponsorName}`)
        .then(resp => {
            setSponsor(resp.data)
        })
        .catch(error=>{
            alert(error)
        })
    }, [])

    return (
        <div style={{backgroundColor:'gainsboro'}}>
            <Image src={sponsor.profile_pic_url} avatar size="small" centered />
            <br/>
            <Icon name="instagram" size="big"/>
            <h2>{sponsor.full_name}</h2>
            <span><b>Media count: {sponsor.media_count}</b></span>
            <br></br>
            <span><b>Follower count: {sponsor.follower_count}</b></span>
            <br></br>
            <span><b>UserTags count: {sponsor.usertags_count}</b></span>
            <br/>
            <Icon name="facebook" size="big"/>
            <br/>
            <span><b>Facebook likes: {sponsor.likes?.split(" ")[0]}</b></span>
            <br/>
            <span><b>Facebook followers: {sponsor.followers?.split(" ")[0]}</b></span>
        </div>
    )
}

export default Sponsor
