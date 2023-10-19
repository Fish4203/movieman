import { Button, Container, Card, Row, Col, Form } from 'react-bootstrap';
import React, { useState } from 'react';
import axios from 'axios';
import {Result} from '../components/Results';
// import "./styles.css";


function Search() {
  const [movies, setMovies] = useState([]);
  const [shows, setShows] = useState([]);
  const [people, setPeople] = useState([]);
  const [episodes, setEpisodes] = useState([]);
  const [query, setQuery] = useState([]);
  
  const [status, setStatus] = useState("");
  const [TMDB, setTMDB] = useState(false)
  
  async function searchApis() {
    if (TMDB) {
      axios.get('http://127.0.0.1:4000/TMDB/search?q=' + query)
      .then(function (response) {})
      .catch(function (error) {
        // handle error
        console.log(error);
        return false;
      });
    }
  
    return true;
  }

    const getSearch = () => {
      searchApis().then(
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
        })
      )
    }

    return (

        <Container>
            <h1>Search </h1>
            <p>{status}</p>

            <div className="d-flex">
                <Form.Control
                    type="search"
                    placeholder="Search"
                    className="me-2"
                    onChange={(event) => setQuery(event.target.value)}
                    value={query}
                    onKeyPress={(event) => {if (event.charCode===13) {getSearch()}}}
                />
                <Button onClick={getSearch}>Search</Button>
                <Form.Check // prettier-ignore
                  type="switch"
                  id="custom-switch"
                  label="TMDB"
                  checked={TMDB}
                  onChange={() => {setTMDB(!TMDB)}}
                />
            </div>



            {/* <Show  shows={shows} lim={4}/> */}
            
            {/* <Movie movies={movies}/> */}
            <h2>Episodes</h2><Row>
            {shows.map(post => ( <Result post={post}/>))}
            </Row>
            {/* <Episode episodes={episodes}/>
            <People people={people}/> */}

            

        </Container>

    );
}

export default Search;
