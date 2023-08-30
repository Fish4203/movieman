import logo from './logo.svg';
import './App.css';
import { Button, Container, Card, Row, Col } from 'react-bootstrap';
import React, { useState, useEffect } from 'react';
import axios from 'axios';


function Search() {
    const [data, setData] = useState([]);

    axios.get('http://127.0.0.1:4000/media/search?q=star trek')
      .then(function (response) {
        // handle success
        setData(response.data.shows);
        console.log(response.data);
        console.log();
      })
      .catch(function (error) {
        // handle error
        console.log(error);
      })
      .finally(function () {
      });

    return (

        <Container>
            <h1>Search </h1>

            <ul>
              {data.map(post => (
                <li>{post.title}</li>
              ))}
            </ul>

            <Row> 
                <Col> 
                    <Card>
                        <Card.Img variant="top" src="https://placehold.co/200x100" />
                        <Card.Body>
                            <Card.Title>Card Title</Card.Title>
                            <Card.Text>
                                Some quick example text to build on the card title and make up the
                                bulk of the card's content.
                            </Card.Text>
                            <Button variant="primary">Go somewhere</Button>
                        </Card.Body>
                    </Card>
                </Col>

                <Col> 
                    <Card>
                        <Card.Img variant="top" src="https://placehold.co/200x100" />
                        <Card.Body>
                            <Card.Title>Card Title</Card.Title>
                            <Card.Text>
                                Some quick example text to build on the card title and make up the
                                bulk of the card's content.
                            </Card.Text>
                            <Button variant="primary">Go somewhere</Button>
                        </Card.Body>
                    </Card>
                </Col>

                <Col> 
                    <Card>
                        <Card.Img variant="top" src="https://placehold.co/200x100" />
                        <Card.Body>
                            <Card.Title>Card Title</Card.Title>
                            <Card.Text>
                                Some quick example text to build on the card title and make up the
                                bulk of the card's content.
                            </Card.Text>
                            <Button variant="primary">Go somewhere</Button>
                        </Card.Body>
                    </Card>
                </Col>

                <Col> 
                    <Card>
                        <Card.Img variant="top" src="https://placehold.co/200x100" />
                        <Card.Body>
                            <Card.Title>Card Title</Card.Title>
                            <Card.Text>
                                Some quick example text to build on the card title and make up the
                                bulk of the card's content.
                            </Card.Text>
                            <Button variant="primary">Go somewhere</Button>
                        </Card.Body>
                    </Card>
                </Col>


            </Row>
            

        </Container>

    );
}

export default Search;
