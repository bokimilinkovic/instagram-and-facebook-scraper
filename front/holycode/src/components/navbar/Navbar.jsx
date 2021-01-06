import React, { useState } from 'react'
import { Link, NavLink, useHistory, useLocation } from 'react-router-dom';
import { Button, Icon, Image, Menu } from 'semantic-ui-react';
//import logo from '../../../public/plus.png'

const Navbar = () => {
     const [isOpen, setOpen] = useState(false);
     const {pathname} = useLocation();
     const history = useHistory();
     const handleUserLogOut = () =>{
        localStorage.removeItem("token")
      history.push("/login")
     }
    return (
        <Menu secondary pointing>
            <Image src='/plus.png' width={60}  />
            <Menu.Item as={Link} to="/" style={{fontSize:24}}>All products</Menu.Item>
           {pathname==="/" && (
            <Menu.Item position="right">
                <Button as={Link} to="/product" icon  primary basic>
                    <Icon name="add"></Icon>
                    Create Product
                </Button>
            </Menu.Item>
            )}
            {pathname==="/" && (
                <Menu.Item>
                    <Button icon onClick={handleUserLogOut} color="red" basic>
                        <Icon name="log out"></Icon>
                        Logout
                    </Button>
                </Menu.Item>
            )}
        </Menu>
    )
}

export default Navbar
