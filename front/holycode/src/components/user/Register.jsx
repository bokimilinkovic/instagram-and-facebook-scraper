import React, { useState } from 'react'
import { Button, Checkbox, Form, Grid } from 'semantic-ui-react'
import Axios from 'axios'
import { useHistory } from 'react-router-dom'

const Register = () => {
    const [username,setUsername] = useState('')
    const [password, setPassword] = useState('')
    const history = useHistory()

    const handleSubmit = (e) => {
        e.preventDefault()
        Axios.post(`${process.env.REACT_APP_API_URL}/v1/users/register`, 
         { username: username, password: password })
         .then(response=>{
             console.log(response.data)
             history.push("/login")
         })
         .catch(error=>{
             alert(error)
         })
        
    }

    return (
        <Grid className="segment centered">
            <Form>
                <Form.Field>
                    <label>Username</label>
                    <input placeholder='First Name' autoComplete="false" value={username} onChange={(e)=>setUsername(e.target.value)} />
                </Form.Field>
                <Form.Field>
                    <label>password</label>
                    <input type="password" placeholder='Last Name' value={password} onChange={(e)=>setPassword(e.target.value)} autoComplete="off"/>
                </Form.Field>
                <Form.Field>
                    <Checkbox label='I agree to the Terms and Conditions' />
                </Form.Field>
                <Button size="small" color="green" type='submit' onClick={handleSubmit}>Submit</Button>
            </Form>
        </Grid>
    )
}

export default Register
