import { Button, Container, Card, Row, Col, Form } from 'react-bootstrap';
import React, { useState } from 'react';
import axios from 'axios';
import {Movie, Show, People, Episode} from '../components/results';
// import "./styles.css";



function Search() {
    const [movies, setMovies] = useState([]);
    const [shows, setShows] = useState([]);
    const [people, setPeople] = useState([]);
    const [episodes, setEpisodes] = useState([]);
    const [query, setQuery] = useState([]);


    const getSearch = () => {
        axios.get('http://127.0.0.1:4000/media/search?q=' + query)
      .then(function (response) {
        // handle success
        console.log(response.data)
        setMovies(response.data.movies);
        setShows(response.data.shows);
        setPeople(response.data.people);
        setEpisodes(response.data.episodes);
        // setData(response.data);
      })
      .catch(function (error) {
        // handle error
        console.log(error);
      })
      .finally(function () {
      });
    }

    return (

        <Container>
            <h1>Search </h1>

            <div className="d-flex">
                <Form.Control
                    type="search"
                    placeholder="Search"
                    className="me-2"
                    onChange={(event) => setQuery(event.target.value)}
                    value={query}
                    onKeyPress={(event) => {if (event.charCode==13) {getSearch()}}}
                />
                <Button onClick={getSearch}>Search</Button>
            </div>



            <Show  shows={shows} lim={4}/>
            <Movie movies={movies}/>
            <Episode episodes={episodes}/>
            <People people={people}/>

            

        </Container>

    );
}

export default Search;
