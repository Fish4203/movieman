import {Form, Button, Container, Card, Row, Col } from 'react-bootstrap';
import React, { useState } from 'react';
import axios from 'axios';
import { useCookies } from 'react-cookie';
import { Outlet, Link, useNavigate } from "react-router-dom";



function Login() {
    const navigate = useNavigate();
    const [name, setName] = useState("");
    const [password, setPassword] = useState("");
    const [loginError, setLoginError] = useState("");
    const [cookies, setCookie, removeCookie] = useCookies(['auth']);

    const createUser = () => {
        axios.post('http://127.0.0.1:4000/user', {name: name, password: password}, {headers: {'Content-Type': 'application/json'}})
      .then(function (response) {
        setLoginError("Created user :" + name);
      })
      .catch(function (error) {
        setLoginError("Error:" + error.response.data.error);
        console.log(error);
      })
      .finally(function () {
      });
    }

    const login = () => {
        axios.post('http://127.0.0.1:4000/login', {name: name, password: password}, {headers: {'Content-Type': 'application/json'}})
      .then(function (response) {
        setLoginError("Loged in");
        setCookie("authToken", response.data.token);
        window.location.reload(false);
        navigate("/")
      })
      .catch(function (error) {
        setLoginError("Error:" + error.response.data.error);
        // console.log(error);
      })
      .finally(function () {
      });
    }


    return (
        <Container>
            <h1>Login </h1>
            {loginError}
            
            <Form>
                <Form.Group className="mb-3" controlId="loginForm.UserName">
                    <Form.Label>Username</Form.Label>
                    <Form.Control type="username" placeholder="name" value={name} onChange={(e) => setName(e.target.value)} />
                </Form.Group>
                <Form.Group className="mb-3" controlId="loginForm.Password">
                    <Form.Label>Password</Form.Label>
                    <Form.Control type="password" placeholder="password" value={password} onChange={(e) => setPassword(e.target.value)} />
                </Form.Group>

                <Button onClick={login}>Login</Button>
                <Button className=" justify-content-end" onClick={createUser}>Create</Button>
            </Form>

        </Container>
 );
}

export default Login;
