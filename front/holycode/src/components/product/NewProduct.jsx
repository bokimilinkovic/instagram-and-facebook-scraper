import axios from 'axios'
import React, { useState } from 'react'
import { useHistory } from 'react-router-dom'
import { Button, Container, Form, Grid } from 'semantic-ui-react'

const NewProduct = () => {
    const [name,setName] = useState('')
    const [price,setPrice] = useState(0)
    const [description,setDescription] = useState('')
    const[fileName,setFilename] = useState('')
    const [file,setFile] = useState({})
    const [sponsor,setSponsor] = useState('')

    const history = useHistory();
    const handleSubmit = (event) => {
    event.preventDefault()
        // filename
    console.log('filename ' + fileName);
    //file 
    console.log('file ' + file);

    const formData = new FormData();
    let newFileName = name
    var ext =  fileName.split('.').pop();
    formData.append('file', file, newFileName+`.${ext}`)
    axios.post('http://localhost:8080/v1/products', {name, price: price,description, sponsor}, { 
                headers: { 
                    'content-type':'application/json',
                    "Authorization": `Bearer ${localStorage.getItem("token")}`
                }
            }
        ).then(data => {
            console.log('file uploaded')
            axios.post('http://localhost:8080/v1/image', formData, { headers: { 'content-type': 'multipart/form-data' }})
            .then(resp => { console.log(resp.data)})
            .catch(error => {console.log(error)})   
            history.push("/")
        }).catch(e => {
            console.log('error')
            console.log(e)
        })
    }

    return (
        <Container>
            <Grid>
                <Grid.Row centered>
                    <Grid.Column width={6}>
                    <Form>
                            <Form.Field>
                                <label>Name</label>
                                <input placeholder='Name' autoComplete="off" value={name} onChange={(e)=>setName(e.target.value)} />
                            </Form.Field>
                            <Form.Field>
                                <label>Price</label>
                                <input placeholder='Price' type="number" autoComplete="off" value={price} onChange={(e)=>setPrice(parseFloat(e.target.value))}/>
                            </Form.Field>
                            <Form.Field>
                                <label>Description</label>
                                <input placeholder="Description"  autoComplete="off" value={description} onChange={(e)=>setDescription(e.target.value)}/>
                            </Form.Field>
                            <Form.Field>
                                <label>Sponsor</label>
                                <input placeholder="Sponsor"  autoComplete="off" value={sponsor} onChange={(e)=>setSponsor(e.target.value)}/>
                            </Form.Field>
                            <Form.Field>
                                <label>Image</label>
                                <input type="file" placeholder="Description"  autoComplete="off" onChange={(e) => { 
                                    setFile(e.target.files[0])
                                    setFilename(e.target.value)
                                    }}  />
                            </Form.Field>
                            <Button color="green" onClick={handleSubmit}>Submit</Button>
                        
                    </Form>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
            <br/>
            <div style={{
                alignContent: 'center',
                backgroundImage: "url(/bck.jpg)",
                backgroundSize: "contain",
                backgroundRepeat: 'no-repeat',
                width: 1200,
                height: 800,
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                paddingtop: "66%",
            }}>

            </div>
        </Container>
        )
}

export default NewProduct
