import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { Card, Grid} from 'semantic-ui-react'
import Product from './Product'

const ProductsList = () => {
    const [products,setProducts] = useState([])
    //nolint
    useEffect(() => {
        axios.get(`${process.env.REACT_APP_API_URL}/v1/products`,{
            headers: {'content-type':'application/json'}
        }).then(resp => {
            setProducts(resp.data)
        }).catch(error => {
            console.log(error)
        })
        console.log(products)
    }, [])
    return (
        <Grid centered >
            <Card.Group itemsPerRow={5} stackable>
                    {products.map((product)=>{ return <Product key={product.ID} product={product} />})}
            </Card.Group>
        </Grid>
    )
}

export default ProductsList
