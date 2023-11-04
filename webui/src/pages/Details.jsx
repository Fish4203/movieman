import { Button, Container, Card, Row, Col } from 'react-bootstrap';
import axios from 'axios';
import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';




function Details() {
    const { type, id} = useParams();
    const [post, setPost] = useState([]);


    useEffect(() => {
        axios.get(import.meta.env.VITE_BACKEND + '/media/' + type + "/" + id)
        .then(function (response) {
          // handle success
          console.log(response.data["movies"][0])
          setPost(response.data["movies"][0])
    
        })
        .catch(function (error) {
          // handle error
          console.log(error);
        })
      }, [])
    

    return (
        <Container>
            <h1>Details </h1>
            <Button> Test </Button>
                {post.title}


        </Container>
 );
}

export default Details;
