import React, { useState } from 'react'
import { Button, Checkbox, Form, Grid } from 'semantic-ui-react'
import Axios from 'axios'
import { Link, useHistory } from 'react-router-dom'

const Login = () => {
    const [username,setUsername] = useState('')
    const [password, setPassword] = useState('')
    const history = useHistory()

    const handleSubmit = (e) => {
        e.preventDefault()
        Axios.post(`${process.env.REACT_APP_API_URL}/v1/users/login`,
         { username: username, password: password })
         .then(response=>{
             const {token} = response.data
             console.log(token)
             localStorage.setItem("token", token)
             history.push("/")
         })
         .catch(error=>{
             alert(error)
         })
        
    }

    return (
        <Grid className="segment centered">
            <h3>Welcome to login page</h3>
            <Form>
                <Form.Field>
                    <label>Username</label>
                    <input placeholder='First Name' autoComplete="off" value={username} onChange={(e)=>setUsername(e.target.value)} />
                </Form.Field>
                <Form.Field>
                    <label>Password</label>
                    <input type="password" placeholder='Last Name' value={password} onChange={(e)=>setPassword(e.target.value)} autoComplete="off"/>
                </Form.Field>
                <Form.Field>
                    <Checkbox label='I agree to the Terms and Conditions' />
                </Form.Field>
                <Button size="small" color="green" type='submit' onClick={handleSubmit}>Submit</Button>
                <p>Dont have account? <Link to="/register">Register</Link></p>
            </Form>
        </Grid>
    )
}

export default Login