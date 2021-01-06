import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { Card, Container, Icon, Image } from 'semantic-ui-react'
import Sponsor from '../sponsor/Sponsor'

const Product = ({product}) => {
    const [likes, setLikes] = useState(15)

    const increaseLikes = () => {
        setLikes(likes => {
                return likes + 1
            })
     }

    const extra = (
        <div>
        <a onClick={increaseLikes}>
            <Icon name='heart'/>
            {likes} Likes
        </a>
        <h5>Price: {product.Price} $</h5>
    </div>
    )

    

    return (
       <div style={{position:"revert", top:'50%',left:'50%'}}>
           <br></br>
            <Container fluid>
                {product.sponsor !== '' ? <Sponsor sponsorName={product.sponsor} /> : 'Not sponsored yet'}
                <h3>Paid sponsorship with</h3>
                <Card centered
                    image={process.env.REACT_APP_API_URL+`/static/${product.name}.jpg`}
                    header={product.name}
                    description={product.Description}
                    extra={extra}
                />
            </Container>
        </div>
            
    )
}



export default Product
