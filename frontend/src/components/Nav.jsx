import React, { useState } from "react";
import {
 Navbar,
 Nav,
 NavDropdown,
 Form,
 FormControl,
 Button,
 Container
} from "react-bootstrap";
import { Outlet, Link, useNavigate } from "react-router-dom";
import { useCookies } from 'react-cookie';
import axios from 'axios';


function Navigation() {
    const navigate = useNavigate();
    const [cookies, setCookie, removeCookie] = useCookies(['auth']);
    const [username, setUserame] = useState("");


    const userInfo = () => {
        axios.get('http://127.0.0.1:4000/user', {headers: {Authorization: 'Bearer ' + cookies.authToken}})
      .then(function (response) {
        setUserame(response.data.user.name);
      })
      .catch(function (error) {
        console.log(error);
        removeCookie("authToken");
      })
      .finally(function () {
      });
    }

    const logout = () => {
        removeCookie("authToken");
    }

    return (
        <>
    <Navbar expand="lg" className="bg-body-tertiary">
        <Container fluid>
            <Navbar.Brand href="#">MovieMan</Navbar.Brand>
            <Navbar.Toggle aria-controls="navbarScroll" />
            <Navbar.Collapse id="navbarScroll">
                <Nav
                className="me-auto my-2 my-lg-0"
                style={{ maxHeight: '100px' }}
                navbarScroll
                >
                    <Nav.Link onClick={() => navigate('')} >Home</Nav.Link>
                    <Nav.Link onClick={() => navigate('search')}>Search</Nav.Link>
                    <Nav.Link href="#action2">Torrents</Nav.Link>

                </Nav>
                <Nav
                className=" justify-content-end"
                style={{ maxHeight: '100px' }}
                navbarScroll
                >
                    {cookies.authToken ? userInfo(): ''}
                    {cookies.authToken ? <><Nav.Link onClick={() => navigate('profile')} >username</Nav.Link> <Nav.Link onClick={() => logout()} >Logout</Nav.Link></> : <Nav.Link onClick={() => navigate('login')} >Login</Nav.Link>}    
                </Nav>
            </Navbar.Collapse>
        </Container>
    </Navbar>
    <Outlet />
    </>);
}

export default Navigation;
