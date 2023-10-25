import { Button, Container, Card, Row, Col, Form } from 'react-bootstrap';
import React, { useState } from 'react';
import axios from 'axios';
import {Result} from '../components/Results';
// import "./styles.css";


function Search() {
  const [movies, setMovies] = useState([]);
  const [shows, setShows] = useState([]);
  const [episodes, setEpisodes] = useState([]);
  const [books, setBooks] = useState([]);
  const [games, setGames] = useState([]);
  const [people, setPeople] = useState([]);
  const [companies, setCompanies] = useState([]);
  const [groups, setGroups] = useState([]);
  const [query, setQuery] = useState([]);
  
  const [status, setStatus] = useState("");
  const [provs, setProvs] = useState([])
  
  async function searchApis(url, types) {
    
    axios.get(url + '/search?q=' + query + '&types=' + types)
    .then(function (response) {
      setStatus("Sucsess")
    })
    .catch(function (error) {
      setStatus(error)
      console.log(error);
    });
  }

  function getSearch() {

    axios.get('http://127.0.0.1:4000/media/search?q=' + query)
    .then(function (response) {
      // handle success
      // console.log(response.data)
      setMovies(response.data.movies);
      setShows(response.data.shows);
      setEpisodes(response.data.episodes);
      setBooks(response.data.books);
      setGames(response.data.Games);
      setPeople(response.data.people);
      setCompanies(response.data.companies);
      setGroups(response.data.groups);
      // setData(response.data);
    })
    .catch(function (error) {
      // handle error
      console.log(error);
    })
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



            <h2>Movies</h2><Row>
            {movies.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Shows</h2><Row>
            {shows.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Episodes</h2><Row>
            {episodes.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Books</h2><Row>
            {books.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Games</h2><Row>
            {games.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>People</h2><Row>
            {people.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Companies</h2><Row>
            {companies.map(post => ( <Result post={post}/>))}
            </Row>
            <h2>Collections</h2><Row>
            {groups.map(post => ( <Result post={post}/>))}
            </Row>
            

            

        </Container>

    );
}

export default Search;
